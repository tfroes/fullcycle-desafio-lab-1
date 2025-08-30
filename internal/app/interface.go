package app

type ClimaAppInterface interface {
	BuscaClima(cep string) (*ClimaModel, error)
}

type CepServiceInterface interface {
	BuscaCep(cep string) (*CepServiceModel, error)
}

type CepServiceModel struct {
	Cep        string
	Localidade string
}

type ClimaServiceInterface interface {
	BuscaClima(localidade string) (*ClimaServiceModel, error)
}

type ClimaServiceModel struct {
	Temp_C float32
}

type ClimaModel struct {
	Temp_C float32
	Temp_F float32
	Temp_K float32
}
