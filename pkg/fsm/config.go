package fsm

import (
	"fmt"
	"strings"
)

type Config struct {
	States       map[string]struct{}
	initialState *string
	finalStates  []string
	Alphabet     map[rune]struct{}
	Transitions  TransitionsMap
}

func NewConfig(states []string, alphabet []rune, initialState string, finalStates []string, transitions []Transition) (*Config, error) {
	// Sanity checks: avoid empy input
	if len(states) == 0 {
		return nil, fmt.Errorf("must have non-zero amount of states")
	}
	if len(alphabet) == 0 {
		return nil, fmt.Errorf("must have non-zero amount of inputs")
	}
	if initialState == "" {
		return nil, fmt.Errorf("must have non-blank initial state")
	}
	if len(transitions) == 0 {
		return nil, fmt.Errorf("must have transition functions")
	}
	if len(finalStates) == 0 {
		return nil, fmt.Errorf("must have some final states")
	}

	newConfig := Config{
		initialState: &initialState,
		finalStates:  finalStates,
	}

	// Create the States, ALphabet, and TransitionMap, and then validate the config
	newConfig.States = make(map[string]struct{}, len(states))
	for _, currentState := range states {
		if strings.TrimSpace(currentState) == "" {
			return nil, ErrEmptyState
		}
		newConfig.States[currentState] = struct{}{}
	}

	newConfig.Alphabet = make(map[rune]struct{}, len(alphabet))
	for _, currentCharacter := range alphabet {
		newConfig.Alphabet[currentCharacter] = struct{}{}
	}

	newConfig.Transitions = NewTransitionsMap(newConfig.States, newConfig.Alphabet)
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
	if _, ok := c.States[*c.initialState]; !ok {
		return fmt.Errorf("initial state invalid")
	}

	for _, finalState := range c.finalStates {
		if _, ok := c.States[finalState]; !ok {
			return fmt.Errorf("%s final state is invalid", finalState)
		}
	}

	return c.Transitions.Validate()
}
