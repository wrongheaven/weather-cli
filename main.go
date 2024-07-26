package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

//go:embed .env
var envFile []byte

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain int8 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func init() {
	tmpFile, err := os.CreateTemp("", ".env")
	if err != nil {
		panic("Error creating temporary file for .env")
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(envFile); err != nil {
		panic("Error writing to temporary .env file")
	}
	if err := tmpFile.Close(); err != nil {
		panic("Error closing temporary .env file")
	}

	if err := godotenv.Load(tmpFile.Name()); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	key := os.Getenv("WEATHERAPI_KEY")

	q := "Oslo"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	url := "http://api.weatherapi.com/v1/forecast.json"
	url += "?key=" + key
	url += "&days=1&q=" + strings.Replace(q, " ", "+", -1)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Convert to json/object/struct/whatever
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %.1f°C - %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		msg := fmt.Sprintf(
			"%s: %.1f°C - %d%% - %s",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Println(msg)
		} else {
			color.Red(msg)
		}
	}
}
