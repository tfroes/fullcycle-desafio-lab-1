package apihandler

import (
	"encoding/json"
	"errors"
	"fullcycle-lab-3/internal/app"
	"net/http"
	"strings"
)

type ClimaApiHandler struct {
	App app.ClimaAppInterface
}

func NewClimaApiHandler(climaApp app.ClimaAppInterface) *ClimaApiHandler {
	return &ClimaApiHandler{climaApp}
}

func (api ClimaApiHandler) GetClimaHandler(w http.ResponseWriter, r *http.Request) {

	cep := strings.Split(r.RequestURI, "/")[2]

	clima, err := api.App.BuscaClima(cep)

	if err != nil {

		if errors.Is(err, app.ErrCepInvalido) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(&ErrorModel{"invalid zipcode"})
			return
		}

		if errors.Is(err, app.ErrCepNaoEncontrado) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&ErrorModel{"can not find zipcode"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorModel{err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clima)
}
