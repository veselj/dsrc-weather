import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import { HttpClient } from '@angular/common/http';


const RTYC_WEATHER_LINK = "https://www.weatherlink.com/embeddablePage/getData/477837b179b94d58b123a4c127c40c50";

// This worked as bulletin - but change ID automatically
// https://www.weatherlink.com/bulletin/9722cfc3-a4ef-47b9-befb-72f52592d6ed


// https://www.weatherlink.com/bulletin/64271a01-7451-40ac-b52e-e7ccf8b5b449


export interface WeatherStationData {
  windDirection: any
  forecastOverview: ForecastOverview[]
  highAtStr: any
  loAtStr: any
  timeZoneId: string
  timeFormat: string
  barometerUnits: string
  windUnits: string
  rainUnits: string
  tempUnits: string
  temperatureFeelLike: any
  temperature: string
  hiTemp: string
  hiTempDate: number
  loTemp: string
  loTempDate: number
  wind: string
  gust: string
  gustAt: number
  humidity: string
  rain: string
  seasonalRain: string
  barometer: string
  barometerTrend: string
  lastReceived: number
  systemLocation: string
  aqsLocation: any
  aqsLastReceived: any
  thwIndex: string
  thswIndex: string
  aqi: any
  aqiString: any
  aqiScheme: any
  noAccess: any
}

export interface ForecastOverview {
  date: string
  morning: Morning
  afternoon: Afternoon
  evening: Evening
  night: Night
}

export interface Morning {
  weatherCode: number
  weatherDesc: string
  weatherIconUrl: string
  temp: number
  chanceofrain: number
  rainInInches: number
}

export interface Afternoon {
  weatherCode: number
  weatherDesc: string
  weatherIconUrl: string
  temp: number
  chanceofrain: number
  rainInInches: number
}

export interface Evening {
  weatherCode: number
  weatherDesc: string
  weatherIconUrl: string
  temp: number
  chanceofrain: number
  rainInInches: number
}

export interface Night {
  weatherCode: number
  weatherDesc: string
  weatherIconUrl: string
  temp: number
  chanceofrain: number
  rainInInches: number
}


@Injectable({
  providedIn: 'root'
})
export class WeatherStationService {

  constructor(private http: HttpClient) {}

  getWeather(): Observable<WeatherStationData> {
    return this.http.get<WeatherStationData>(RTYC_WEATHER_LINK);
  }
}
