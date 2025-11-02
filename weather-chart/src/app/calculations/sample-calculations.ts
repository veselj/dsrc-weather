import {WeatherData} from '../services/wind-chart-data-service';


export type GraphDataPoint = { x: number; y: number };

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

    public getAverageWindSpeed(hoursBack: number) {
      const sampleSet = this.filterDataByHoursBack(hoursBack);
      return sampleSet.reduce((sum, entry) => sum + entry.Wd, 0) / sampleSet.length;
    }


  public getMovingAverages(fromSecs: number): GraphDataPoint[] {
    const movingAverages: GraphDataPoint[] = [];

    const sampleSet = this.filterDataFrom(fromSecs);
    const windowSize = 10; // Number of data points in the moving average window

    for (let i = 0; i < sampleSet.length; i++) {
      let from = Math.max(0, i - windowSize);
      let localAverage = this.localAverageSpeed(sampleSet, from, i);
      movingAverages.push({ x: sampleSet[i].Wn * 1000, y: localAverage });
    }
    return movingAverages;
  }

  private filterDataFrom(fromSecs: number): WeatherData[] {
    return this.weatherData.filter(entry => entry.Wn >= fromSecs);
  }

  private filterDataByHoursBack(hoursBack: number): WeatherData[] {
    const hoursBackDateTimeSecs = Math.floor(Date.now()/1000) - (hoursBack * 3600);
    return this.weatherData.filter(entry => entry.Wn >= hoursBackDateTimeSecs);
  }

  private localAverageSpeed(sampleSet: WeatherData[], from:number, to:number):number{
    let sum = 0;
    for(let i=from; i<=to; i++){
      sum += sampleSet[i].Wd;
    }
    let la = sum / (to - from + 1);
    return la;
  }
}
