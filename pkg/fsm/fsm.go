package fsm

type FiniteStateMachine struct {
	Config Config
}

func New(config Config) (*FiniteStateMachine, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	newFSM := FiniteStateMachine{
		Config: config,
	}

	return &newFSM, nil
}

func (f *FiniteStateMachine) Process(input string) (*string, bool) {
	if len(input) == 0 {
		return nil, false
	}
	currentState := f.Config.initialState

	for _, currentRune := range input {
		_, ok := f.Config.Transitions.alphabet[currentRune]
		if !ok {
			return nil, false
		}

		// Accepted input, get the current transition for this state
		inputMap, ok := f.Config.Transitions.transitions[currentState]
		if !ok {
			return nil, false
		}
		// is there a valid new state for this input?
		newState, ok := inputMap[currentRune]
		if !ok {
			return nil, false
		}

		currentState = newState
	}

	// final check: did we end up in a correct state?
	// in this implementation, ending up in a final state not specified in the config will return an invalid result
	if _, ok := f.Config.finalStates[currentState]; !ok {
		return nil, false
	}

	return &currentState, true
}
