package main

import (
	"fullcycle-lab-3/internal/apihandler"
	"fullcycle-lab-3/internal/app"
	"fullcycle-lab-3/internal/viacep"
	"fullcycle-lab-3/internal/weatherapi"
	"net/http"
)

func main() {
	mux := IniciaServeMux()
	http.ListenAndServe(":8080", mux)
}

func IniciaServeMux() *http.ServeMux {

	cepService := viacep.NewViaCepService()
	climaService := weatherapi.NewWeatherService()

	application := app.NewApp(cepService, climaService)

	apiHandler := apihandler.NewClimaApiHandler(application)

	mux := http.NewServeMux()
	mux.HandleFunc("/clima/{cep}", apiHandler.GetClimaHandler)

	return mux
}
