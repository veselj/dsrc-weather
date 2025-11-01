import { Component, signal, NgModule } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {WindChart} from './wind-chart/wind-chart';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, WindChart],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('weather-chart');
}
