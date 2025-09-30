package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransitionsMap(t *testing.T) {
	states := map[string]struct{}{
		"q0": {},
		"q1": {},
	}
	alphabet := map[rune]struct{}{
		'a': {},
		'b': {},
	}

	tm := NewTransitionsMap(states, alphabet)

	assert.NotNil(t, tm.states, "states should be initialized")
	assert.NotNil(t, tm.alphabet, "alphabet should be initialized")
	assert.NotNil(t, tm.transitions, "Transitions should be initialized")
	assert.Equal(t, 2, len(tm.states), "should have 2 states")
	assert.Equal(t, 2, len(tm.alphabet), "should have 2 alphabet symbols")
	assert.Equal(t, 0, len(tm.transitions), "should start with empty transitions")
}

func TestNewTransitions(t *testing.T) {
	states := map[string]struct{}{
		"q0": {},
		"q1": {},
		"q2": {},
	}
	alphabet := map[rune]struct{}{
		'a': {},
		'b': {},
	}

	tm := NewTransitionsMap(states, alphabet)

	type test struct {
		name          string
		input         Transition
		expectedError error
	}

	tests := []test{
		{
			name: "valid input",
			input: Transition{
				State:       "q0",
				Input:       'a',
				ResultState: "q1",
			},
			expectedError: nil,
		},
		{
			name: "valid input - overwritng previous input",
			input: Transition{
				State:       "q0",
				Input:       'a',
				ResultState: "q0",
			},
			expectedError: nil,
		},
		{
			name: "invalid state",
			input: Transition{
				State:       "q5",
				Input:       'a',
				ResultState: "q1",
			},
			expectedError: ErrInvalidState,
		},
		{
			name: "invalid input",
			input: Transition{
				State:       "q0",
				Input:       'c',
				ResultState: "q1",
			},
			expectedError: ErrInvalidInput,
		},
		{
			name: "invalid result",
			input: Transition{
				State:       "q0",
				Input:       'a',
				ResultState: "q5",
			},
			expectedError: ErrInvalidResultState,
		},
	}

	for _, currentTest := range tests {
		err := tm.NewTransition(currentTest.input)
		if currentTest.expectedError == nil {
			assert.Nil(t, err, currentTest.name)
		}
		if currentTest.expectedError != nil {
			assert.ErrorIs(t, err, currentTest.expectedError, currentTest.name)
		}
	}
}

func TestValidateTransitionsMap(t *testing.T) {

	type test struct {
		name                string
		states              map[string]struct{}
		alphabet            map[rune]struct{}
		transitions         []Transition
		expectedErrorString string
	}

	tests := []test{
		{
			name: "valid setup",
			states: map[string]struct{}{
				"q0": {},
				"q1": {},
			},
			alphabet: map[rune]struct{}{
				'a': {},
				'b': {},
			},
			transitions: []Transition{
				{
					State:       "q0",
					Input:       'a',
					ResultState: "q1",
				},
				{
					State:       "q0",
					Input:       'b',
					ResultState: "q0",
				},
				{
					State:       "q1",
					Input:       'a',
					ResultState: "q0",
				},
				{
					State:       "q1",
					Input:       'b',
					ResultState: "q1",
				},
			},
		},
		{
			name: "valid setup - loop transition",
			states: map[string]struct{}{
				"q0": {},
			},
			alphabet: map[rune]struct{}{
				'a': {},
				'b': {},
			},
			transitions: []Transition{
				{
					State:       "q0",
					Input:       'a',
					ResultState: "q0",
				},
				{
					State:       "q0",
					Input:       'b',
					ResultState: "q0",
				},
			},
		},
		{
			name: "invalid setup - missing state transitions",
			states: map[string]struct{}{
				"q0": {},
				"q1": {},
			},
			alphabet: map[rune]struct{}{
				'a': {},
				'b': {},
			},
			transitions: []Transition{
				{
					State:       "q0",
					Input:       'a',
					ResultState: "q1",
				},
				{
					State:       "q0",
					Input:       'b',
					ResultState: "q0",
				},
			},
			expectedErrorString: "missing transitions for state q1",
		},
		{
			name: "invalid setup - missing input transitions",
			states: map[string]struct{}{
				"q0": {},
			},
			alphabet: map[rune]struct{}{
				'a': {},
				'b': {},
			},
			transitions: []Transition{
				{
					State:       "q0",
					Input:       'a',
					ResultState: "q0",
				},
			},
			expectedErrorString: "missing transitions for state q0 for input b",
		},
	}
	for _, currentTest := range tests {
		tm := NewTransitionsMap(currentTest.states, currentTest.alphabet)

		for _, currentTransistion := range currentTest.transitions {
			err := tm.NewTransition(currentTransistion)
			assert.Nil(t, err, currentTest.name)
		}

		validationError := tm.Validate()
		if currentTest.expectedErrorString == "" {
			assert.Nil(t, validationError, currentTest.name)
		}
		if currentTest.expectedErrorString != "" {
			assert.Contains(t, validationError.Error(), currentTest.expectedErrorString, currentTest.name)
		}
	}
}

func TestSpecialCharacterTransitions(t *testing.T) {
	states := map[string]struct{}{
		"start": {},
		"space": {},
		"tab":   {},
		"end":   {},
	}
	alphabet := map[rune]struct{}{
		' ':  {},
		'\t': {},
		'\n': {},
		'@':  {},
	}

	tm := NewTransitionsMap(states, alphabet)

	transitions := []Transition{
		{State: "start", Input: ' ', ResultState: "space"},
		{State: "start", Input: '\t', ResultState: "tab"},
		{State: "start", Input: '\n', ResultState: "end"},
		{State: "start", Input: '@', ResultState: "end"},
		{State: "space", Input: ' ', ResultState: "space"},
		{State: "space", Input: '\t', ResultState: "tab"},
		{State: "space", Input: '\n', ResultState: "end"},
		{State: "space", Input: '@', ResultState: "end"},
		{State: "tab", Input: ' ', ResultState: "space"},
		{State: "tab", Input: '\t', ResultState: "tab"},
		{State: "tab", Input: '\n', ResultState: "end"},
		{State: "tab", Input: '@', ResultState: "end"},
		{State: "end", Input: ' ', ResultState: "space"},
		{State: "end", Input: '\t', ResultState: "tab"},
		{State: "end", Input: '\n', ResultState: "end"},
		{State: "end", Input: '@', ResultState: "end"},
	}

	for _, tr := range transitions {
		err := tm.NewTransition(tr)
		assert.Nil(t, err, "should handle special character transitions")
	}

	err := tm.Validate()
	assert.Nil(t, err, "special character FSM should be valid")
}
