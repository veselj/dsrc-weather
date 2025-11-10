import { Routes } from '@angular/router';
import { WindChart } from './wind-chart/wind-chart';
import { TempChart } from './temp-chart/temp-chart';

export const routes: Routes = [
  { path: 'wind', component: WindChart },
  { path: 'temp', component: TempChart },
  { path: '', redirectTo: '/wind', pathMatch: 'full' }
];
