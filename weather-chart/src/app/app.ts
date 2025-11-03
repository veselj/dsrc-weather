import { Component, signal, NgModule } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {WindChart} from './wind-chart/wind-chart';
import {WeatherStation} from './weather-station/weather-station';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, WindChart, WeatherStation],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('weather-chart');
}
