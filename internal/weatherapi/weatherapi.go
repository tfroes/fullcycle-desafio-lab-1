package weatherapi

import (
	"encoding/json"
	"fullcycle-lab-3/internal/app"
	"io"
	"net/http"
	"net/url"
)

var key = "8fee6bae245a40c2a16200419251908"

type WeatherService struct{}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (v *WeatherService) BuscaClima(localidade string) (*app.ClimaServiceModel, error) {

	req, err := http.NewRequest(
		"GET",
		"http://api.weatherapi.com/v1/current.json?key="+key+"&q="+url.PathEscape(localidade)+"&aqi=no",
		nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var m GetCurretModel
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return m.Map(), nil
}
