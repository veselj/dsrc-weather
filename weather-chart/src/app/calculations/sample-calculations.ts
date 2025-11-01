import {WeatherData} from '../services/wind-chart-data-service';


export class SampleCalculation {


  constructor(private weatherData: WeatherData[]) {
  }

  public getWindSpeedData(hoursBack: number) {
      let hoursBackDateTime = Date.now() - hoursBack * 3600 * 1000;
      let da = this.weatherData
        .map(entry => ({
          x: entry.Wn * 1000, // Unix timestamp in milliseconds
          y: entry.Wd  // Wind direction in degrees
        }))
        .filter(point => point.x >= hoursBackDateTime);
      return da;
    }

}
