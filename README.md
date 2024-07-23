# Weather CLI Tool

This is a simple Command Line Interface (CLI) tool written in Go that fetches and displays weather forecasts for a specified location. It utilizes the Weather API to gather current weather conditions and forecasts.

## Features

-   Fetch current weather conditions including temperature and weather description.
-   Fetch weather forecast for the rest of the day, displaying hourly forecasts.
-   Color-coded output for easy reading, with high chances of rain highlighted in red.
-   Accepts user input for dynamic location queries.

## Prerequisites

Before you begin, ensure you have met the following requirements:

-   You have installed Go version 1.22.2 or higher.
-   You have an account with [Weather API](http://api.weatherapi.com/) with an API key.

## Installation

To install the Weather CLI tool, follow these steps:

1. Clone the repository to your local machine:

```sh
git clone https://github.com/wrongheaven/weather-cli.git
```

2. Navigate to the cloned repository:

```sh
cd weather-cli
```

3. Create a `.env` file based on the `.env.example` provided, and insert your Weather API key.

```sh
cp .env.example .env
```

4. Build the project:

```sh
go build -o weather-cli
```

5. Move the binary to a location in path:

```sh
mv weather-cli /usr/local/bin
```

## Usage

To use the Weather CLI tool, execute the binary followed by the location for which you want to fetch the weather forecast. If no location is provided, it defaults to "Oslo".

Run the tool for a specific location:

```sh
weather-cli "New York"
```

Or, run the tool for the default location:

```sh
weather-cli
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

-   Weather data provided by [Weather API](http://api.weatherapi.com/).
-   Color output support by [fatih/color](https://github.com/fatih/color).
