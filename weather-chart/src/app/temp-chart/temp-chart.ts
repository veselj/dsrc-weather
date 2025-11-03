import {Component, ViewChild, HostListener, OnDestroy, OnInit} from '@angular/core';
import {BaseChartDirective} from 'ng2-charts';
import {ChartConfiguration, TimeScaleOptions} from 'chart.js';
import 'chartjs-adapter-date-fns';
import {WeatherData, WindChartDataService} from '../services/wind-chart-data-service';
import {OverallStats, SampleCalculation} from '../calculations/sample-calculations';

type GranularityType = 'minute' | 'hour';

@Component({
  selector: 'app-temp-chart',
  standalone: true,
  imports: [BaseChartDirective],
  templateUrl: './temp-chart.html',
  styleUrl: './temp-chart.css'
})
export class TempChart implements OnInit, OnDestroy {
  @ViewChild(BaseChartDirective) chart?: BaseChartDirective; // Declare chart property

  loading = true;
  public granularity: 'minute' | 'hour' = 'hour'; // Granularity control
  private hoursBackRetrieval  = 6; // Hours back for data retrieval
  public hoursBack: number = 1; // Hours back for display
  public tempChartData: ChartConfiguration<'line'>['data'] = {
    datasets: []
  };
  private resizeListener: () => void;
  public tempChartType: 'line' = 'line';
  private calc?: SampleCalculation;
  private subtitle: string = '';

 public tempChartOptions: ChartConfiguration<'line'>['options'] = {
    responsive: true,
   maintainAspectRatio: false,
    plugins: {
      legend: {
        display: true,
        position: 'top', // Move legend to the top
        labels: {
          font: {
            size: 14,
            family: 'Arial, sans-serif'
          },
          color: '#333'
        }
      },
      title: {
        display: true,
        text: 'Temperature Over Time (°C)',
        font: {
          size: 18,
          family: 'Arial, sans-serif'
        },
        color: '#444'
      },
      subtitle: {
        display: true,
        text: '',
        font: {
          size: 14,
          family: 'Arial, sans-serif'
        },
        color: '#666'
      },
      tooltip: {
        backgroundColor: 'rgba(0, 0, 0, 0.7)',
        titleFont: { size: 14 },
        bodyFont: { size: 12 },
        cornerRadius: 4
      }
    },
    scales: {
      y: {
        beginAtZero: true,
        title: {
          display: true,
          text: 'Temperature (°C)',
          font: {
            size: 14,
            family: 'Arial, sans-serif'
          },
          color: '#555'
        },
        grid: {
          color: 'rgba(201,201,201,0.3)'
        }
      },
      x: {
        type: 'time',
        time: {
          unit: this.granularity,
          displayFormats: {
            minute: 'HH:mm',
            hour: 'HH:mm'
          }
        },
        title: {
          display: true,
          text: 'Time',
          font: {
            size: 14,
            family: 'Arial, sans-serif'
          },
          color: '#555'
        },
        grid: {
          color: 'rgba(200, 200, 200, 0.3)'
        }
      }
    },
    elements: {
      line: {
        borderWidth: 2, // Thinner lines
        tension: 0.4 // Smooth curves
      },
      point: {
        radius: 4, // Larger points
        hoverRadius: 6
      }
    }
  };

  constructor(private windChartDataService: WindChartDataService) {

    this.windChartDataService.getData(this.hoursBackRetrieval).subscribe(data => {
       this.loading = false;
       this.calc = new SampleCalculation(data);
       this.tempChartData = this.getChartDataSet(this.hoursBack);
    });
    this.resizeListener = () => { };
  }

  ngOnInit(): void {
    this.resizeListener = this.onResize.bind(this);
    window.addEventListener('resize', this.resizeListener);
  }

  ngOnDestroy(): void {
    window.removeEventListener('resize', this.resizeListener);
  }

  @HostListener('window:resize')
  onResize(): void {
    this.chart?.update(); // Trigger chart update on resize
  }

  setGranularity(granularity :GranularityType): void {
    this.granularity = granularity
    const xScale = this.tempChartOptions?.scales?.['x'] as TimeScaleOptions;
    if (xScale?.time) {
      xScale.time.unit = this.granularity;
    }
    // Reassign the options object to trigger Angular change detection
    this.tempChartOptions = { ...this.tempChartOptions };
    this.chart?.update(); // Trigger chart update
  }

  getChartDataSet(hoursBack: number): ChartConfiguration<'line'>['data'] {

    if (!this.calc) {
      return { datasets: [] };
    }

    const samples = this.calc.getTemperatureData(hoursBack);
    if (samples.length == 0) {
      return { datasets: [] };
    }
    const overallStats = this.calc.getOverallStats(samples);
    const feelsLikeSamples = this.calc.getFeelsLikeTemperatureData(hoursBack);
    const feelsLikeStats = this.calc.getOverallStats(feelsLikeSamples);

    this.setSubtitle(overallStats, feelsLikeStats);

    return {
      datasets: [
        {
          label: 'Temperature (°C)',
          data: samples,
          fill: false,
          tension: 0.3,
          borderColor: '#1976d2',
          backgroundColor: 'rgba(25, 118, 210, 0.2)',
          pointBackgroundColor: '#1976d2'
        },
        {
          label: 'Feels like (°C)',
          data: feelsLikeSamples,
          fill: false,
          borderDash: [5, 5],
          borderColor: '#054702',
          backgroundColor: 'rgba(207,148,129,0.2)',
          pointBackgroundColor: '#2f8817'
        }
      ]
    };
  }

  setSubtitle(stats: OverallStats, feelsLike: OverallStats): void {
    if (this.tempChartOptions?.plugins?.subtitle) {
      this.tempChartOptions.plugins.subtitle.text = [`Temperature Min: ${stats.min.toFixed(2)}, Max: ${stats.max.toFixed(2)}, Avg: ${stats.average.toFixed(2)}`,
      `Feels Like Min: ${feelsLike.min.toFixed(2)}, Max: ${feelsLike.max.toFixed(2)}, Avg: ${feelsLike.average.toFixed(2)}`];
    }
  }

  setHistory(hours: number): void {
    this.hoursBack = hours;
    this.tempChartData = this.getChartDataSet(this.hoursBack);
    let granularity: GranularityType = hours <= 1 ? 'minute' : 'hour';
    this.setGranularity(granularity);
    this.chart?.update(); // Trigger chart update
  }

  public historyOptions: number[] = [0.5, 1, 3, 6]; // Options for history

}
