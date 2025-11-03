import {WeatherData} from '../services/wind-chart-data-service';


export type GraphDataPoint = { x: number; y: number };
export type OverallStats = { min: number; max: number; average: number };

export class SampleCalculation {

  constructor(private weatherData: WeatherData[]) {
  }

  public getWindSpeedData(hoursBack: number): GraphDataPoint[] {
    const sampleSet = this.filterDataByHoursBack(hoursBack);
    return sampleSet
        .map(entry => ({
          x: entry.Wn * 1000, // Unix timestamp in milliseconds
          y: entry.Wd  // Wind speed in knots
        }));
    }

    public getOverallStats(sampleSet :GraphDataPoint[]): OverallStats {
      let sum = 0;
      let min = Number.POSITIVE_INFINITY;
      let max = Number.NEGATIVE_INFINITY;
      for (let i = 0; i < sampleSet.length; i++) {
        sum += sampleSet[i].y;
        if (sampleSet[i].y < min) {
          min = sampleSet[i].y;
        }
        if (sampleSet[i].y > max) {
          max = sampleSet[i].y;
        }
      }
      return { min: min, max: max, average: sum / sampleSet.length };
    }


  public getMovingAverages(sampleSet :GraphDataPoint[]): GraphDataPoint[] {
    const movingAverages: GraphDataPoint[] = [];

    const windowSize = 10; // Number of data points in the moving average window

    for (let i = 0; i < sampleSet.length; i++) {
      let from = Math.max(0, i - windowSize);
      let localAverage = this.localAverageSpeed(sampleSet, from, i);
      movingAverages.push({ x: sampleSet[i].x, y: localAverage });
    }
    return movingAverages;
  }

  private filterDataByHoursBack(hoursBack: number): WeatherData[] {
    const hoursBackDateTimeSecs = Math.floor(Date.now()/1000) - (hoursBack * 3600);
    return this.weatherData.filter(entry => entry.Wn >= hoursBackDateTimeSecs);
  }

  private localAverageSpeed(sampleSet: GraphDataPoint[], from:number, to:number):number{
    let sum = 0;
    for(let i=from; i<=to; i++){
      sum += sampleSet[i].y;
    }
    let la = sum / (to - from + 1);
    return la;
  }
}
