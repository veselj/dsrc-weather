import { Component, OnInit, OnDestroy } from '@angular/core';
import { ChartConfiguration, ChartType } from 'chart.js';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Subscription, interval } from 'rxjs';
import { NgChartsModule } from 'ng2-charts';

@Component({
  selector: 'app-wind-speed-chart',
  imports: [HttpClientModule, NgChartsModule],
  templateUrl: './wind-speed-chart.component.html',
  styleUrl: './wind-speed-chart.component.css'
})
export class WindSpeedChartComponent implements OnInit, OnDestroy {
  public lineChartData: ChartConfiguration<'line'>['data'] = {
    labels: [],
    datasets: [
      {
        data: [],
        label: 'Wind Speed (kts)',
        fill: true,
        borderColor: 'blue',
        backgroundColor: 'rgba(30,144,255,0.2)',
        tension: 0.4
      }
    ]
  };

  public lineChartOptions: ChartConfiguration<'line'>['options'] = {
    responsive: true,
    plugins: {
      legend: { display: true },
      title: { display: true, text: 'Wind Speed Over Time' }
    },
    scales: {
      x: {},
      y: { title: { display: true, text: 'kts' } }
    }
  };

  public lineChartType: ChartType = 'line';

  private apiUrl = 'https://example.com/api/wind-speed'; // TODO: Replace with real endpoint
  private pollSub?: Subscription;

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    // Initial fetch
    this.fetchWindSpeed();
    // Poll every minute
    this.pollSub = interval(60000).subscribe(() => this.fetchWindSpeed());
  }

  ngOnDestroy(): void {
    this.pollSub?.unsubscribe();
  }

  fetchWindSpeed(): void {
    this.http.get<{ timestamp: string, speed: number }[]>(this.apiUrl).subscribe(data => {
      // Assume API returns array of { timestamp, speed }
      this.lineChartData.labels = data.map(d => d.timestamp);
      this.lineChartData.datasets[0].data = data.map(d => d.speed);
    });
  }
}
