import { Component } from '@angular/core';
import {WeatherStationService} from '../services/weather-station-service';

@Component({
  selector: 'app-weather-station',
  standalone: true,
  imports: [],
  templateUrl: './weather-station.html',
  styleUrl: './weather-station.css'
})
export class WeatherStation {

  constructor(private weatherStationService: WeatherStationService) {
    this.weatherStationService.getWeather().subscribe((data) => {
      console.log("Weatherstation", JSON.stringify(data));
    })
  }

}
