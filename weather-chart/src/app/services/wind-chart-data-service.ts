import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


export type WeatherData = {
  Wd: number; // Wind speed in knots
  Dn: number; // Directional number
  Te: number; // Temperature
  Fl: number; // Feels-like temperature
  Wn: number; // When sample taken (Unix timestamp)
  Bt: string; // Hourly Group in string format
};

@Injectable({
  providedIn: 'root'
})
export class WindChartDataService {
  private apiUrl = 'https://4w4vljd7q24rgeo7c42afyzpze0xqmhx.lambda-url.eu-west-1.on.aws';

  constructor(private http: HttpClient) {}


  getData(hourSpan?: number): Observable<WeatherData[]> {
   let url: string;
    if (hourSpan) {
      url = `${this.apiUrl}?hours=${hourSpan}`;
    } else {
      url = this.apiUrl;
    }
    return this.http.get<WeatherData[]>(url);
  }
}
