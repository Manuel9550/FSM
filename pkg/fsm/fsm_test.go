package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is a representation of the 'mod-three' Finite State Machine.
// The tests here will verify that the behaviour works as intended.

func TestMod3(t *testing.T) {

	StateS0 := "S0"
	StateS1 := "S1"
	StateS2 := "S2"

	states := []string{StateS0, StateS1, StateS2}
	alphabet := []rune{'0', '1'}
	initialState := StateS0
	finalStates := []string{StateS0, StateS1, StateS2}
	transitions := []Transition{
		{
			State:       StateS0,
			Input:       '0',
			ResultState: StateS0,
		},
		{
			State:       StateS0,
			Input:       '1',
			ResultState: StateS1,
		},
		{
			State:       StateS1,
			Input:       '0',
			ResultState: StateS2,
		},
		{
			State:       StateS1,
			Input:       '1',
			ResultState: StateS0,
		},
		{
			State:       StateS2,
			Input:       '0',
			ResultState: StateS1,
		},
		{
			State:       StateS2,
			Input:       '1',
			ResultState: StateS2,
		},
	}

	conf, err := NewConfig(states, alphabet, initialState, finalStates, transitions)
	assert.Nil(t, err)
	if err != nil {
		t.Fatal("Configuration should not have resulted in an error")
	}

	fsm, err := New(*conf)
	assert.Nil(t, err)
	if err != nil {
		t.Fatal("Creating finite state machine should not have resulted in an error")
	}

	type test struct {
		name             string
		input            string
		expectedState    *string
		expectedValidity bool
	}

	tests := []test{
		// Valid tests
		{
			name:             "Standard test 1",
			input:            "110",
			expectedState:    &StateS0,
			expectedValidity: true,
		},
		{
			name:             "Standard test 2",
			input:            "1010",
			expectedState:    &StateS1,
			expectedValidity: true,
		},
		{
			name:             "Standard test 3",
			input:            "1010001010101001101",
			expectedState:    &StateS1,
			expectedValidity: true,
		},
		{
			name:             "Standard test 4",
			input:            "10101010111001",
			expectedState:    &StateS2,
			expectedValidity: true,
		},
		{
			name:             "Leading 0's",
			input:            "0010101010111001",
			expectedState:    &StateS2,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - All 0",
			input:            "00000",
			expectedState:    &StateS0,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - All 1",
			input:            "111111",
			expectedState:    &StateS0,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - Just 0",
			input:            "0",
			expectedState:    &StateS0,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - Just 1",
			input:            "0",
			expectedState:    &StateS0,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - Very long value",
			input:            "101010101010101010101010101010101010101010",
			expectedState:    &StateS0,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - 1",
			input:            "1",
			expectedState:    &StateS1,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - 2",
			input:            "10",
			expectedState:    &StateS2,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - 3",
			input:            "11",
			expectedState:    &StateS0,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - 4",
			input:            "100",
			expectedState:    &StateS1,
			expectedValidity: true,
		},
		{
			name:             "Valid Input - 8",
			input:            "1000",
			expectedState:    &StateS2,
			expectedValidity: true,
		},
		// Invalid tests
		{
			name:             "Invalid Input - Empty",
			input:            "",
			expectedState:    nil,
			expectedValidity: false,
		},
		{
			name:             "Invalid Input - Extra Integer",
			input:            "012001010",
			expectedState:    nil,
			expectedValidity: false,
		},
		{
			name:             "Invalid Input - Newline",
			input:            "01001010\n",
			expectedState:    nil,
			expectedValidity: false,
		},
		{
			name:             "Invalid Input - Wrong Characters",
			input:            "01001A010",
			expectedState:    nil,
			expectedValidity: false,
		},
		{
			name:             "Invalid Input - one space",
			input:            "101010 0101",
			expectedState:    nil,
			expectedValidity: false,
		},
	}

	for _, currentTest := range tests {
		resultState, validity := fsm.Process(currentTest.input)
		assert.Equal(t, validity, currentTest.expectedValidity, currentTest.name)
		assert.Equal(t, resultState == nil, currentTest.expectedState == nil, currentTest.name)
		if currentTest.expectedState != nil && resultState != nil {
			assert.Equal(t, *resultState, *currentTest.expectedState, currentTest.name)
		}
	}
}
