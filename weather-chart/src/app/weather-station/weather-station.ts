import { Component } from '@angular/core';
import {WeatherResponse, WeatherStationService} from '../services/weather-station-service';
import { DecimalPipe, CommonModule, Location } from '@angular/common';

@Component({
  selector: 'app-weather-station',
  standalone: true,
  imports: [DecimalPipe, CommonModule],
  templateUrl: './weather-station.html',
  styleUrl: './weather-station.css'
})
export class WeatherStation {
  weatherData?: WeatherResponse;

  constructor(private weatherStationService: WeatherStationService, private location: Location) {
    this.weatherStationService.getData().subscribe((data: WeatherResponse) => {
      this.weatherData = data;
    });
  }

  getNextTideIndex(): number | null {
    const currentTime = Date.now();
    if (this.weatherData?.tides) {
      return this.weatherData.tides.findIndex(tide => tide.Time * 1000 > currentTime);
    }
    return null;
  }

  getTimeRemaining(tideTime: number): string {
    const currentTime = Date.now();
    const timeDiff = tideTime * 1000 - currentTime;

    const hours = Math.floor(Math.abs(timeDiff) / (1000 * 60 * 60));
    const minutes = Math.floor((Math.abs(timeDiff) % (1000 * 60 * 60)) / (1000 * 60));

    if (timeDiff <= 0) {
      return `-${hours}h ${minutes}m`;
    }

    return `${hours}h ${minutes}m`;
  }

  currentTime(): Date {
    return new Date();
  }

  goBack(): void {
    this.location.back();
  }


}
