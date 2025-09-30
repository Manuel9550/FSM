package fsm

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidState       = errors.New("invalid state")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidResultState = errors.New("invalid result state")
	ErrEmptyState         = errors.New("state cannot be empty")
	ErrEmptyResultState   = errors.New("result state cannot be empty")
	ErrNilTransition      = errors.New("transition cannot be nil")
)

type Transition struct {
	State       string
	Input       rune
	ResultState string
}

type TransitionsMap struct {
	states      map[string]struct{}
	alphabet    map[rune]struct{}
	transitions map[string]map[rune]string
}

func NewTransitionsMap(states map[string]struct{}, alphabet map[rune]struct{}) TransitionsMap {
	if states == nil {
		states = make(map[string]struct{})
	}
	if alphabet == nil {
		alphabet = make(map[rune]struct{})
	}

	return TransitionsMap{
		states:      states,
		alphabet:    alphabet,
		transitions: make(map[string]map[rune]string),
	}
}

func (t *TransitionsMap) NewTransition(transition Transition) error {
	if _, ok := t.states[transition.State]; !ok {
		return ErrInvalidState
	}

	if _, ok := t.alphabet[transition.Input]; !ok {
		return ErrInvalidInput
	}

	if _, ok := t.states[transition.ResultState]; !ok {
		return ErrInvalidResultState
	}

	if t.transitions[transition.State] == nil {
		t.transitions[transition.State] = make(map[rune]string)
	}

	t.transitions[transition.State][transition.Input] = transition.ResultState
	return nil
}

// Note: In this implementation, the transition map will be invalid if
// there is a state that doesn't have an input set for a possible alphabet character.
// Ex: If S1 is a State, and 'A' and 'B' are both valid inputs, but there is no (S1, 'B') mapping, it's marked as Invalid
func (t *TransitionsMap) Validate() error {
	for state := range t.states {
		inputMap, ok := t.transitions[state]
		if !ok {
			return fmt.Errorf("missing transitions for state %s", state)
		}
		for input := range t.alphabet {
			_, ok := inputMap[input]
			if !ok {
				return fmt.Errorf("missing transitions for state %s for input %c", state, input)
			}
		}
	}

	return nil
}
