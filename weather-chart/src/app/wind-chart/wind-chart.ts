import {Component, ViewChild, HostListener, OnDestroy, OnInit} from '@angular/core';
import {BaseChartDirective} from 'ng2-charts';
import {ChartConfiguration, TimeScaleOptions} from 'chart.js';
import 'chartjs-adapter-date-fns';
import {WeatherData, WindChartDataService} from '../services/wind-chart-data-service';
import {OverallStats, SampleCalculation} from '../calculations/sample-calculations';

type GranularityType = 'minute' | 'hour';

@Component({
  selector: 'app-wind-chart',
  standalone: true,
  imports: [BaseChartDirective],
  templateUrl: './wind-chart.html',
  styleUrl: './wind-chart.css'
})
export class WindChart implements OnInit, OnDestroy {
  @ViewChild(BaseChartDirective) chart?: BaseChartDirective; // Declare chart property

  loading = true;
  public granularity: 'minute' | 'hour' = 'hour'; // Granularity control
  private hoursBackRetrival  = 6; // Hours back for data retrieval
  public hoursBack: number = 1; // Hours back for display
  public windChartData: ChartConfiguration<'line'>['data'] = {
    datasets: []
  };
  private resizeListener: () => void;
  public windChartType: 'line' = 'line';
 private calc?: SampleCalculation;
 private subtitle: string = '';

 public windChartOptions: ChartConfiguration<'line'>['options'] = {
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
        text: 'Wind Speed Over Time (knots)',
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
          text: 'Wind Speed (knots)',
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

    this.windChartDataService.getData(this.hoursBackRetrival).subscribe(data => {
       this.loading = false;
       this.calc = new SampleCalculation(data);
       this.windChartData = this.getChartDataSet(this.hoursBack);
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

  // // Method to toggle granularity
  // toggleGranularity(): void {
  //   this.granularity = this.granularity === 'minute' ? 'hour' : 'minute';
  //   const xScale = this.windChartOptions?.scales?.['x'] as TimeScaleOptions;
  //   if (xScale?.time) {
  //     xScale.time.unit = this.granularity;
  //   }
  //   // Reassign the options object to trigger Angular change detection
  //   this.windChartOptions = { ...this.windChartOptions };
  //   this.chart?.update(); // Trigger chart update
  // }

  setGranularity(granularity :GranularityType): void {
    this.granularity = granularity
    const xScale = this.windChartOptions?.scales?.['x'] as TimeScaleOptions;
    if (xScale?.time) {
      xScale.time.unit = this.granularity;
    }
    // Reassign the options object to trigger Angular change detection
    this.windChartOptions = { ...this.windChartOptions };
    this.chart?.update(); // Trigger chart update
  }

  getChartDataSet(hoursBack: number): ChartConfiguration<'line'>['data'] {

    if (!this.calc) {
      return { datasets: [] };
    }

    const samples = this.calc.getWindSpeedData(hoursBack);
    if (samples.length == 0) {
      return { datasets: [] };
    }
    const overallStats = this.calc.getOverallStats(samples);
    const movingAverages = this.calc.getMovingAverages(samples);

    this.setSubtitle(overallStats);

    return {
      datasets: [
        {
          label: 'Wind Speed (every minute)',
          data: samples,
          fill: false,
          tension: 0.3,
          borderColor: '#1976d2',
          backgroundColor: 'rgba(25, 118, 210, 0.2)',
          pointBackgroundColor: '#1976d2'
        },
        {
          label: 'Moving Average (last 10 minutes)',
          data: movingAverages,
          fill: false,
          borderDash: [5, 5],
          borderColor: '#054702',
          backgroundColor: 'rgba(207,148,129,0.2)',
          pointBackgroundColor: '#2f8817'
        }
      ]
    };
  }

  setSubtitle(stats: OverallStats): void {
    if (this.windChartOptions?.plugins?.subtitle) {
      this.windChartOptions.plugins.subtitle.text = `Min: ${stats.min.toFixed(2)}, Max: ${stats.max.toFixed(2)}, Avg: ${stats.average.toFixed(2)}`;
    }
  }

  setHistory(hours: number): void {
    this.hoursBack = hours;
    this.windChartData = this.getChartDataSet(this.hoursBack);
    let granularity: GranularityType = hours <= 1 ? 'minute' : 'hour';
    this.setGranularity(granularity);
    this.chart?.update(); // Trigger chart update
  }

  public historyOptions: number[] = [0.5, 1, 3, 6]; // Options for history

}
