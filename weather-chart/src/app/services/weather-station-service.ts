import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, shareReplay } from 'rxjs';


export interface WeatherResponse {
  weather: Weather
  tides: Tide[]
}

export interface Weather {
  Bucket: string
  When: number // Unix timestamp in milliseconds
  WindSpeed: number
  Temperature: number
  FeelsLike: number
  WindDirection: number
  WindDirectionName: string
  Barometer: number
  BarometerUnits: string
  BarometerTrend: string
  Rain: number
  RainUnits: string
  ChanceOfRain: number
  Humidity: number
  Forecast: string
}

export interface Tide {
  Type: number
  Time: number
  Height: number
}

@Injectable({
  providedIn: 'root'
})
export class WeatherStationService {
  private apiUrl = 'https://4w4vljd7q24rgeo7c42afyzpze0xqmhx.lambda-url.eu-west-1.on.aws?current=yes';
  private cachedData?: Observable<WeatherResponse>;

  constructor(private http: HttpClient) {}

  getData(hourSpan?: number): Observable<WeatherResponse> {
    if (!this.cachedData) {
      this.cachedData = this.http.get<WeatherResponse>(this.apiUrl).pipe(shareReplay(1));
    }
    return this.cachedData
  }
}
