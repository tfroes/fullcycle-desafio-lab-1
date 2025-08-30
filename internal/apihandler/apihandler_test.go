package apihandler

import (
	"errors"
	"fullcycle-lab-3/internal/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClimaAppMock struct {
	mock.Mock
}

func (m *ClimaAppMock) BuscaClima(cep string) (*app.ClimaModel, error) {
	args := m.Mock.Called(cep)

	a0 := args.Get(0)

	if a0 == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*app.ClimaModel), args.Error(1)
}

func Test_ErroCepInvalido_ComCepInvalido(t *testing.T) {
	appMock := &ClimaAppMock{}

	appMock.On("BuscaClima", "12345678").Return(nil, app.ErrCepInvalido)

	apiHandler := NewClimaApiHandler(appMock)

	request := &http.Request{
		RequestURI: "/clima/12345678",
	}

	responseWriter := httptest.NewRecorder()

	apiHandler.GetClimaHandler(responseWriter, request)

	assert.Equal(t, http.StatusUnprocessableEntity, responseWriter.Code)
}

func Test_ErroCepInvalido_ComCepNaoEncontrado(t *testing.T) {
	appMock := &ClimaAppMock{}

	appMock.On("BuscaClima", "12345678").Return(nil, app.ErrCepNaoEncontrado)

	apiHandler := NewClimaApiHandler(appMock)

	request := &http.Request{
		RequestURI: "/clima/12345678",
	}

	responseWriter := httptest.NewRecorder()

	apiHandler.GetClimaHandler(responseWriter, request)

	assert.Equal(t, http.StatusNotFound, responseWriter.Code)
}

func Test_ErroCepInvalido_ComErroQualquer(t *testing.T) {
	appMock := &ClimaAppMock{}

	appMock.On("BuscaClima", "12345678").Return(nil, errors.New("Erro qualquer"))

	apiHandler := NewClimaApiHandler(appMock)

	request := &http.Request{
		RequestURI: "/clima/12345678",
	}

	responseWriter := httptest.NewRecorder()

	apiHandler.GetClimaHandler(responseWriter, request)

	assert.Equal(t, http.StatusInternalServerError, responseWriter.Code)
}
