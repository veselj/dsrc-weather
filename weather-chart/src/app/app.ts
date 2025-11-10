import { Component, signal } from '@angular/core';
import { RouterOutlet, RouterModule } from '@angular/router';
import {WindChart} from './wind-chart/wind-chart';
import {WeatherStation} from './weather-station/weather-station';
import {TempChart} from './temp-chart/temp-chart';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterModule],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('weather-chart');
  selectedChart: 'wind' | 'temp' = 'wind'; // Default to Wind Chart
}
