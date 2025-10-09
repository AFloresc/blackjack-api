package errors

import "errors"

var (
	ErrGameAlreadyStarted = errors.New("la partida ya está iniciada")
	ErrInvalidAction      = errors.New("acción inválida para el estado actual")
	ErrNoActiveGame       = errors.New("no hay partida activa")
)
