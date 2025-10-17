import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class WindChartDataService {
  getWindChartData(): { x: number; y: number }[] {
    const now = Date.now();
    return [
      { x: now, y: 5 },
      { x: now + 600000, y: 4 }, // 10 minutes later
      { x: now + 1200000, y: 7 }, // 20 minutes later
      { x: now + 1800000, y: 8 }, // 30 minutes later
      { x: now + 2400000, y: 6 }, // 40 minutes later
      { x: now + 3000000, y: 5 }, // 50 minutes later
      { x: now + 3600000, y: 7 }  // 60 minutes later
    ];
  }
}
