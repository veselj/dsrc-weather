import { Routes } from '@angular/router';
import { WindChart } from './wind-chart/wind-chart';
import { TempChart } from './temp-chart/temp-chart';
import { WeatherStation } from './weather-station/weather-station';

export const routes: Routes = [
  { path: 'wind', component: WindChart },
  { path: 'temp', component: TempChart },
  { path: 'weather', component: WeatherStation },
  { path: '', redirectTo: '/wind', pathMatch: 'full' }
];
