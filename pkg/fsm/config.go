package fsm

import (
	"fmt"
	"strings"
)

type Config struct {
	initialState string
	finalStates  map[string]struct{}
	Transitions  TransitionsMap
}

func NewConfig(states []string, alphabet []rune, initialState string, finalStates []string, transitions []Transition) (*Config, error) {
	// Sanity checks: avoid empy input
	if len(states) == 0 {
		return nil, ErrEmptyStates
	}
	if len(alphabet) == 0 {
		return nil, ErrEmptyAlphabet
	}
	if initialState == "" {
		return nil, ErrEmptyInitialState
	}
	if len(transitions) == 0 {
		return nil, ErrEmptyTransitions
	}
	if len(finalStates) == 0 {
		return nil, ErrEmptyFinalStates
	}

	newConfig := Config{
		initialState: initialState,
	}

	// Create the States, Alphabet, and TransitionMap, and then validate the config
	newStates := make(map[string]struct{}, len(states))
	for _, currentState := range states {
		if strings.TrimSpace(currentState) == "" {
			return nil, ErrEmptyState
		}
		newStates[currentState] = struct{}{}
	}

	newConfig.finalStates = make(map[string]struct{}, len(finalStates))
	for _, currentState := range finalStates {
		if strings.TrimSpace(currentState) == "" {
			return nil, ErrEmptyState
		}
		newConfig.finalStates[currentState] = struct{}{}
	}

	newAlphabet := make(map[rune]struct{}, len(alphabet))
	for _, currentCharacter := range alphabet {
		newAlphabet[currentCharacter] = struct{}{}
	}

	newConfig.Transitions = NewTransitionsMap(newStates, newAlphabet)
	for _, transition := range transitions {
		transitionError := newConfig.Transitions.NewTransition(transition)
		if transitionError != nil {
			return nil, fmt.Errorf("invalid transition for %s:%c:%s - %s", transition.State, transition.Input, transition.ResultState, transitionError)
		}
	}

	err := newConfig.Validate()
	if err != nil {
		return nil, err
	}

	return &newConfig, nil
}

func (c *Config) Validate() error {
	if _, ok := c.Transitions.states[c.initialState]; !ok {
		return fmt.Errorf("initial state invalid")
	}

	for finalState := range c.finalStates {
		if _, ok := c.Transitions.states[finalState]; !ok {
			return fmt.Errorf("%s final state is invalid", finalState)
		}
	}

	return c.Transitions.Validate()
}
