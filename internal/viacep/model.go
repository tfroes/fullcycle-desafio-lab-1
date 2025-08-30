package viacep

import "fullcycle-lab-3/internal/app"

type GetModel struct {
	Cep         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	Uf          string
	Estado      string
	Regiao      string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
}

func (m *GetModel) Map() *app.CepServiceModel {
	return &app.CepServiceModel{
		Cep:        m.Cep,
		Localidade: m.Localidade,
	}
}
