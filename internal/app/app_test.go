package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CepServiceMock struct {
	mock.Mock
}

type ClimaServiceMock struct {
	mock.Mock
}

func (m *CepServiceMock) BuscaCep(cep string) (*CepServiceModel, error) {
	args := m.Mock.Called(cep)

	a0 := args.Get(0)

	if a0 == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*CepServiceModel), args.Error(1)
}

func (m *ClimaServiceMock) BuscaClima(localidade string) (*ClimaServiceModel, error) {
	args := m.Mock.Called(localidade)

	a0 := args.Get(0)

	if a0 == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*ClimaServiceModel), args.Error(1)
}

func Test_ErroCepInvalido_ComCepInvalido(t *testing.T) {
	cepMock := &CepServiceMock{}
	climaMock := &ClimaServiceMock{}

	app := NewApp(cepMock, climaMock)

	_, err := app.BuscaClima("10")

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrCepInvalido)
}

func Test_ErroCepInvalido_ComCepValido_ErroBuscaCep(t *testing.T) {
	cepMock := &CepServiceMock{}
	climaMock := &ClimaServiceMock{}

	cepMock.On("BuscaCep", "12345678").Return(nil, errors.New("Erro qualquer"))

	app := NewApp(cepMock, climaMock)

	_, err := app.BuscaClima("12345678")

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrCepInvalido)
}

func Test_ErrCepNaoEncontrado_ComCepValido_BuscaCepNaoEncontrado(t *testing.T) {
	cepMock := &CepServiceMock{}
	climaMock := &ClimaServiceMock{}

	cepServiceModel := &CepServiceModel{
		Cep: "",
	}

	cepMock.On("BuscaCep", "12345678").Return(cepServiceModel, nil)

	app := NewApp(cepMock, climaMock)

	_, err := app.BuscaClima("12345678")

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrCepNaoEncontrado)
}

func Test_ErrCepNaoEncontrado_ComCepValido_BuscaCep_ErroBuscaClima(t *testing.T) {
	cepMock := &CepServiceMock{}
	climaMock := &ClimaServiceMock{}

	cepServiceModel := &CepServiceModel{
		Cep:        "12345678",
		Localidade: "Local",
	}

	cepMock.On("BuscaCep", "12345678").Return(cepServiceModel, nil)
	climaMock.On("BuscaClima", "Local").Return(nil, errors.New("Erro qualquer"))

	app := NewApp(cepMock, climaMock)

	_, err := app.BuscaClima("12345678")

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrCepNaoEncontrado)
}

func Test_BuscaClimaCerto(t *testing.T) {
	cepMock := &CepServiceMock{}
	climaMock := &ClimaServiceMock{}

	cepServiceModel := &CepServiceModel{
		Cep:        "12345678",
		Localidade: "Local",
	}

	climaServiceModel := &ClimaServiceModel{
		Temp_C: 10.0,
	}

	cepMock.On("BuscaCep", "12345678").Return(cepServiceModel, nil)
	climaMock.On("BuscaClima", "Local").Return(climaServiceModel, nil)

	app := NewApp(cepMock, climaMock)

	climaModel, err := app.BuscaClima("12345678")

	var c float32 = 10
	var f float32 = 50
	var k float32 = 283

	assert.Nil(t, err)
	assert.NotNil(t, climaModel)
	assert.Equal(t, c, climaModel.Temp_C)
	assert.Equal(t, f, climaModel.Temp_F)
	assert.Equal(t, k, climaModel.Temp_K)
}
