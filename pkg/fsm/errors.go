package fsm

import (
	"errors"
)

var (
	ErrInvalidState        = errors.New("invalid state")
	ErrInvalidInitialState = errors.New("invalid initial state")
	ErrInvalidInput        = errors.New("invalid input")
	ErrInvalidResultState  = errors.New("invalid result state")
	ErrEmptyState          = errors.New("state cannot be empty")
	ErrEmptyFinalState     = errors.New("final state cannot be blank string")
	ErrEmptyResultState    = errors.New("result state cannot be empty")
	ErrNilTransition       = errors.New("transition cannot be nil")

	ErrEmptyAlphabet     = errors.New("must have non-zero amount of inputs")
	ErrEmptyStates       = errors.New("must have non-zero amount of states")
	ErrEmptyTransitions  = errors.New("must have transition functions")
	ErrEmptyFinalStates  = errors.New("must have some final states")
	ErrEmptyInitialState = errors.New("must have non-blank initial state")
)
