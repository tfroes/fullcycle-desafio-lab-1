package viacep

import (
	"encoding/json"
	"fullcycle-lab-3/internal/app"
	"io"
	"net/http"
)

type ViaCepService struct{}

func NewViaCepService() *ViaCepService {
	return &ViaCepService{}
}

func (v *ViaCepService) BuscaCep(cep string) (*app.CepServiceModel, error) {

	req, err := http.NewRequest(
		"GET",
		"http://viacep.com.br/ws/"+cep+"/json/",
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

	var m GetModel
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return m.Map(), nil
}
