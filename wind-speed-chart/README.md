# Wind Speed Chart Application

This Angular application displays a real-time chart of wind speed (in knots) using Chart.js and ng2-charts. The chart fetches wind speed data from a REST API every minute and updates automatically.

## How to Run

1. Install dependencies:
	```bash
	npm install
	```
2. Start the development server:
	```bash
	ng serve
	```
3. Open your browser and navigate to [http://localhost:4200](http://localhost:4200).

## How It Works

- The main chart is displayed on the homepage.
- The app fetches wind speed data from a REST API endpoint every minute.
- The chart updates in real time as new data arrives.
- The REST endpoint URL is currently set as a placeholder in `wind-speed-chart.component.ts` (replace `https://example.com/api/wind-speed` with your actual API).

## Customizing the Data Source

Edit the `apiUrl` property in `src/app/wind-speed-chart/wind-speed-chart.component.ts` to point to your REST API. The API should return an array of objects with the following structure:

```
[
  { "timestamp": "2025-09-15T10:00:00Z", "speed": 12 },
  { "timestamp": "2025-09-15T10:01:00Z", "speed": 14 },
  ...
]
```

## Dependencies

- Angular 19+
- Chart.js v4
- ng2-charts v4
# WindSpeedChart

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 19.2.16.

## Development server

To start a local development server, run:

```bash
ng serve
```

Once the server is running, open your browser and navigate to `http://localhost:4200/`. The application will automatically reload whenever you modify any of the source files.

## Code scaffolding

Angular CLI includes powerful code scaffolding tools. To generate a new component, run:

```bash
ng generate component component-name
```

For a complete list of available schematics (such as `components`, `directives`, or `pipes`), run:

```bash
ng generate --help
```

## Building

To build the project run:

```bash
ng build
```

This will compile your project and store the build artifacts in the `dist/` directory. By default, the production build optimizes your application for performance and speed.

## Running unit tests

To execute unit tests with the [Karma](https://karma-runner.github.io) test runner, use the following command:

```bash
ng test
```

## Running end-to-end tests

For end-to-end (e2e) testing, run:

```bash
ng e2e
```

Angular CLI does not come with an end-to-end testing framework by default. You can choose one that suits your needs.

## Additional Resources

For more information on using the Angular CLI, including detailed command references, visit the [Angular CLI Overview and Command Reference](https://angular.dev/tools/cli) page.
