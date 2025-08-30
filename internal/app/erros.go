package app

import "errors"

var ErrCepInvalido = errors.New("cep inválido")
var ErrCepNaoEncontrado = errors.New("cep não encontrado")
