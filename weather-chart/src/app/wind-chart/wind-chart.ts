import { Component } from '@angular/core';
import { BaseChartDirective } from 'ng2-charts';
import { ChartConfiguration, ChartType } from 'chart.js';
import 'chartjs-adapter-date-fns';

@Component({
  selector: 'app-wind-chart',
  standalone: true,
  imports: [BaseChartDirective],
  templateUrl: './wind-chart.html',
  styleUrl: './wind-chart.css'
})
export class WindChart {
  // Example: Wind speed readings every 5 minutes for 1 hour
  public windChartData: ChartConfiguration<'line'>['data'] = {
    datasets: [
      {
        label: 'Wind Speed (knots)',
        data: [
          { x: Date.now(), y: 5 },
          { x: Date.now() + 600000, y: 6 }, // 10 minutes later
          { x: Date.now() + 1200000, y: 7 }, // 20 minutes later
          { x: Date.now() + 1800000, y: 8 }, // 30 minutes later
          { x: Date.now() + 2400000, y: 6 }, // 40 minutes later
          { x: Date.now() + 3000000, y: 5 }, // 50 minutes later
          { x: Date.now() + 3600000, y: 7 }  // 60 minutes later
        ],
        fill: false,
        tension: 0.3,
        borderColor: '#1976d2',
        backgroundColor: 'rgba(25, 118, 210, 0.2)',
        pointBackgroundColor: '#1976d2'
      }
    ]
  };

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
            unit: 'minute',
            displayFormats: {
                minute: 'HH:mm'
            }
        },
        title: { display: true, text: 'Time' }
      }
    }
  };
}
