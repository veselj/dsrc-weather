import {Component, ViewChild} from '@angular/core';
import {BaseChartDirective} from 'ng2-charts';
import {ChartConfiguration, TimeScaleOptions} from 'chart.js';
import 'chartjs-adapter-date-fns';
import {WeatherData, WindChartDataService} from '../services/wind-chart-data-service';

@Component({
  selector: 'app-wind-chart',
  standalone: true,
  imports: [BaseChartDirective],
  templateUrl: './wind-chart.html',
  styleUrl: './wind-chart.css'
})
export class WindChart {
  @ViewChild(BaseChartDirective) chart?: BaseChartDirective; // Declare chart property

  public granularity: 'minute' | 'hour' = 'minute'; // Granularity control
  public windChartData: ChartConfiguration<'line'>['data'] = {
    datasets: []
  };
  public windChartType: 'line' = 'line';
  private retrievedData: WeatherData[] = [];

  public windChartOptions: ChartConfiguration<'line'>['options'] = {
    responsive: true,
    plugins: {
      legend: { display: true },
      title: {
        display: true,
        text: 'Wind Speed (knots) Over Time'
      }
    },
    scales: {
      y: {
        beginAtZero: true,
        title: { display: true, text: 'Speed (knots)' }
      },
      x: {
        type: 'time',
        // adapters: {
        //   date: {
        //     // Use Unix time (seconds) for parsing
        //     parse: (value: number) => new Date(value * 1000)
        //   }
        // },
        time: {
          unit: this.granularity, // Dynamically set granularity
          displayFormats: {
            minute: 'HH:mm',
            hour: 'HH:mm'
          }
        },
        title: { display: true, text: 'Time' }
      }
    }
  };

  constructor(private windChartDataService: WindChartDataService) {

    this.windChartDataService.getData().subscribe(data => {
       console.log('retrieved data:', JSON.stringify(data));
       this.retrievedData = data;

       this.windChartData = {
          datasets: [
            {
              label: 'Wind Speed (knots)',
              data: this.getWindSpeedData(this.retrievedData),
              fill: false,
              tension: 0.3,
              borderColor: '#1976d2',
              backgroundColor: 'rgba(25, 118, 210, 0.2)',
              pointBackgroundColor: '#1976d2'
            }
          ]
        };
    });
  }

  // Method to toggle granularity
  toggleGranularity(): void {
    this.granularity = this.granularity === 'minute' ? 'hour' : 'minute';
    const xScale = this.windChartOptions?.scales?.['x'] as TimeScaleOptions;
    if (xScale?.time) {
      xScale.time.unit = this.granularity;
    }
    // Reassign the options object to trigger Angular change detection
    this.windChartOptions = { ...this.windChartOptions };
    this.chart?.update(); // Trigger chart update
  }

  getWindSpeedData(data: WeatherData[]) {

    let da = data.map(entry => ({
      x: entry.Wn * 1000, // Unix timestamp in milliseconds
      y: entry.Wd  // Wind direction in degrees
    }));
    console.log("Converted" + da.length + " data points for wind chart.");
    console.log(JSON.stringify(da));
    return da;
  }
}
