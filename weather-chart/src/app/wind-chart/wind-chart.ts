import { Component, ViewChild } from '@angular/core';
import { BaseChartDirective } from 'ng2-charts';
import { ChartConfiguration, ChartType, TimeScaleOptions } from 'chart.js';
import 'chartjs-adapter-date-fns';
import {WindChartDataService} from '../services/wind-speed-service';

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
  public windChartData: ChartConfiguration<'line'>['data'];
  public windChartType: 'line' = 'line';

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
    this.windChartData = {
      datasets: [
        {
          label: 'Wind Speed (knots)',
          data: this.windChartDataService.getWindChartData(),
          fill: false,
          tension: 0.3,
          borderColor: '#1976d2',
          backgroundColor: 'rgba(25, 118, 210, 0.2)',
          pointBackgroundColor: '#1976d2'
        }
      ]
    };
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
}
