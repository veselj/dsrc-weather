import { Component } from '@angular/core';
import {WeatherResponse, WeatherStationService} from '../services/weather-station-service';
import { DecimalPipe, CommonModule } from '@angular/common';

@Component({
  selector: 'app-weather-station',
  standalone: true,
  imports: [DecimalPipe, CommonModule],
  templateUrl: './weather-station.html',
  styleUrl: './weather-station.css'
})
export class WeatherStation {
  weatherData?: WeatherResponse;

  constructor(private weatherStationService: WeatherStationService) {
    this.weatherStationService.getData().subscribe((data: WeatherResponse) => {
      this.weatherData = data;
    });
  }

  currentTime(): Date {
    return new Date();
  }

}
