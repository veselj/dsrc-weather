import { Component } from '@angular/core';
import {WeatherResponse, WeatherStationService} from '../services/weather-station-service';
import { CommonModule} from '@angular/common';

@Component({
  selector: 'app-tide-table',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './tide-table.html',
  styleUrl: './tide-table.css'
})
export class TideTable {
  weatherData?: WeatherResponse;

  constructor(private weatherStationService: WeatherStationService) {
    this.weatherStationService.getData().subscribe((data: WeatherResponse) => {
      this.weatherData = data;
    });
  }

  getNextTideIndex(): number | null {
    const currentTime = Date.now() / 1000;
    if (this.weatherData?.tides) {
      const index = this.weatherData.tides.findIndex(tide => tide.Time  > currentTime);
      return index;
    }
    return null;
  }

  getTimeRemaining(tideTime: number): string {
    const currentTime = Date.now() / 1000;
    const timeDiff = tideTime - currentTime;

    const hours = Math.floor(Math.abs(timeDiff) / 3600);
    const minutes = Math.floor((Math.abs(timeDiff) % 3600) / 60);

    let prefix = '';
    if (hours > 0) {
       prefix = `${hours}h `;
    }
    if (timeDiff <= 0) {
      return `-${prefix}${minutes}m`;
    }

    return `${prefix}${minutes}m`;
  }

  currentTime(): Date {
    return new Date();
  }

  sampleTime() : Date {
    return this.weatherData?.weather?.When ? new Date(this.weatherData.weather.When) : new Date();
  }

}
