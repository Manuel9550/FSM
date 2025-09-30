package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidConfig(t *testing.T) {
	type test struct {
		name         string
		states       []string
		alphabet     []rune
		initialState string
		finalStates  []string
		transitions  []Transition
	}

	tests := []test{
		{
			name:     "valid config",
			states:   []string{"q0", "q1"},
			alphabet: []rune{'a', 'b'},
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
			initialState: "q0",
			finalStates:  []string{"q1"},
		},
		{
			name:     "valid config many states, few inputs",
			states:   []string{"q0", "q1", "q2", "q3"},
			alphabet: []rune{'a'},
			transitions: []Transition{
				{
					State:       "q0",
					Input:       'a',
					ResultState: "q0",
				},
				{
					State:       "q1",
					Input:       'a',
					ResultState: "q0",
				},
				{
					State:       "q2",
					Input:       'a',
					ResultState: "q0",
				},
				{
					State:       "q3",
					Input:       'a',
					ResultState: "q0",
				},
			},
			initialState: "q0",
			finalStates:  []string{"q0"},
		},
	}

	for _, currentTest := range tests {
		_, err := NewConfig(currentTest.states, currentTest.alphabet, currentTest.initialState, currentTest.finalStates, currentTest.transitions)
		assert.Nil(t, err, currentTest.name)
	}
}

func TestInValidConfig(t *testing.T) {
	type test struct {
		name          string
		states        []string
		alphabet      []rune
		initialState  string
		finalStates   []string
		transitions   []Transition
		expectedError string
	}

	tests := []test{
		{
			name:     "invalid - missing states",
			states:   []string{},
			alphabet: []rune{'a', 'b'},
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
			initialState:  "q0",
			finalStates:   []string{"q1"},
			expectedError: "must have non-zero amount of states",
		},
		{
			name:     "invalid - missing alphabet",
			states:   []string{"q0", "q1"},
			alphabet: []rune{},
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
			initialState:  "q0",
			finalStates:   []string{"q1"},
			expectedError: "must have non-zero amount of inputs",
		},
		{
			name:     "invalid - blank initial state",
			states:   []string{"q0", "q1"},
			alphabet: []rune{'a', 'b'},
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
			initialState:  "",
			finalStates:   []string{"q1"},
			expectedError: "must have non-blank initial state",
		},
		{
			name:     "invalid - initial state is invalid",
			states:   []string{"q0", "q1"},
			alphabet: []rune{'a', 'b'},
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
			initialState:  "q2",
			finalStates:   []string{"q1"},
			expectedError: "initial state invalid",
		},
		{
			name:     "invalid - final states have invalid state",
			states:   []string{"q0", "q1"},
			alphabet: []rune{'a', 'b'},
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
			initialState:  "q1",
			finalStates:   []string{"q3"},
			expectedError: "final state is invalid",
		},
		{
			name:     "invalid - missing transition for input (q0, 'a')",
			states:   []string{"q0", "q1"},
			alphabet: []rune{'a', 'b'},
			transitions: []Transition{
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
			initialState:  "q1",
			finalStates:   []string{"q1"},
			expectedError: "missing transitions for state q0 for input a",
		},
		{
			name:          "invalid - transitions are empty",
			states:        []string{"q0", "q1"},
			alphabet:      []rune{'a', 'b'},
			transitions:   []Transition{},
			initialState:  "q1",
			finalStates:   []string{"q1"},
			expectedError: "must have transition functions",
		},
		{
			name:     "invalid - transition has character not defined in alphabet",
			states:   []string{"q0", "q1"},
			alphabet: []rune{'a', 'b'},
			transitions: []Transition{
				{
					State:       "q0",
					Input:       'c',
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
			initialState:  "q1",
			finalStates:   []string{"q1"},
			expectedError: "invalid transition for q0:c:q0 - invalid input",
		},
	}

	for _, currentTest := range tests {
		_, err := NewConfig(currentTest.states, currentTest.alphabet, currentTest.initialState, currentTest.finalStates, currentTest.transitions)
		assert.NotNil(t, err, currentTest.name)
		if err != nil {
			assert.Contains(t, err.Error(), currentTest.expectedError, currentTest.name)
		}
	}
}
