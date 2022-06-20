package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type weather struct {
}

type Weather interface {
	Predict(city string, date time.Time, wind bool) string
}

func NewWeather() Weather {
	return weather{}
}

type Location struct {
	Latitude  float32
	Longitude float32
}
type PositionStackResponse struct {
	Data []Location
}

type OpenMeteoResponse struct {
	Daily WeatherPrediction
}

type WeatherPrediction struct {
	Windspeed_10m_max []float32
	Weathercode       []int
	Time              []string
}

func (w weather) Predict(city string, date time.Time, wind bool) string {
	if date.Before(time.Now().AddDate(0, 0, 6)) {

		// Find the latitude and longitude to get the prediction
		responsePS, _ := http.Get("https://positionstack.com/geo_api.php?query=" + city)
		defer responsePS.Body.Close()
		contentPS, _ := ioutil.ReadAll(responsePS.Body)
		var decoded PositionStackResponse
		json.Unmarshal(contentPS, &decoded)
		latitude := decoded.Data[0].Latitude
		longitude := decoded.Data[0].Longitude

		// Find the predictions for the location
		var url = fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&daily=weathercode,windspeed_10m_max&current_weather=true&timezone=Europe/Berlin", latitude, longitude)
		response, _ := http.Get(url)
		defer response.Body.Close()
		content, _ := ioutil.ReadAll(response.Body)
		var openMeteoResponse OpenMeteoResponse
		json.Unmarshal(content, &openMeteoResponse)
		for i := 0; i < 6; i++ {
			if openMeteoResponse.Daily.Time[i] == date.Format("2006-01-02") {
				if wind {
					return fmt.Sprintf("%.2f", openMeteoResponse.Daily.Windspeed_10m_max[i])
				} else {
					return w.codeToText(openMeteoResponse.Daily.Weathercode[i])
				}
			}
		}
	}
	return ""
}

func (w weather) codeToText(code int) string {
	texts := map[int]string{
		0:  "Clear sky",
		1:  "Mainly clear, partly cloudy, and overcast",
		2:  "Mainly clear, partly cloudy, and overcast",
		3:  "Mainly clear, partly cloudy, and overcast",
		45: "Fog and depositing rime fog",
		48: "Fog and depositing rime fog",
		51: "Drizzle: Light, moderate, and dense intensity",
		53: "Drizzle: Light, moderate, and dense intensity",
		55: "Drizzle: Light, moderate, and dense intensity",
		56: "Freezing Drizzle: Light and dense intensity",
		57: "Freezing Drizzle: Light and dense intensity",
		61: "Rain: Slight, moderate and heavy intensity",
		63: "Rain: Slight, moderate and heavy intensity",
		65: "Rain: Slight, moderate and heavy intensity",
		66: "Freezing Rain: Light and heavy intensity",
		67: "Freezing Rain: Light and heavy intensity",
		71: "Snow fall: Slight, moderate, and heavy intensity",
		73: "Snow fall: Slight, moderate, and heavy intensity",
		75: "Snow fall: Slight, moderate, and heavy intensity",
		77: "Snow grains",
		80: "Rain showers: Slight, moderate, and violent",
		81: "Rain showers: Slight, moderate, and violent",
		82: "Rain showers: Slight, moderate, and violent",
		85: "Snow showers slight and heavy",
		86: "Snow showers slight and heavy",
		95: "Thunderstorm: Slight or moderate",
		96: "Thunderstorm with slight and heavy hail",
		99: "Thunderstorm with slight and heavy hail",
	}
	return texts[code]

}
