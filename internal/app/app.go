package app

import (
	"regexp"
)

type ClimaApp struct {
	CepService   CepServiceInterface
	ClimaService ClimaServiceInterface
}

func NewApp(cepServiceInterface CepServiceInterface, climaServiceInterface ClimaServiceInterface) ClimaAppInterface {
	return &ClimaApp{cepServiceInterface, climaServiceInterface}
}

func (app *ClimaApp) BuscaClima(cep string) (*ClimaModel, error) {

	/*** CEP ***/

	regexMatch, err := regexp.MatchString("\\d{8}", cep)

	if err != nil {
		return nil, ErrCepInvalido
	}

	if !regexMatch {
		return nil, ErrCepInvalido
	}

	cepServiceModel, err := app.CepService.BuscaCep(cep)

	if err != nil {
		return nil, ErrCepInvalido
	}

	if cepServiceModel.Cep == "" {
		return nil, ErrCepNaoEncontrado
	}

	/*** Clima ***/

	climaServiceModel, err := app.ClimaService.BuscaClima(cepServiceModel.Localidade)

	if err != nil {
		return nil, ErrCepNaoEncontrado
	}

	clima := &ClimaModel{
		Temp_C: climaServiceModel.Temp_C,
		Temp_F: ConverteParaFahrenheit(climaServiceModel.Temp_C),
		Temp_K: ConverteParaKelvin(climaServiceModel.Temp_C),
	}

	return clima, nil
}

func ConverteParaFahrenheit(celsius float32) float32 {
	return celsius*1.8 + 32
}

func ConverteParaKelvin(celsius float32) float32 {
	return celsius + 273
}
