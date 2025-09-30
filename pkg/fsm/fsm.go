package fsm

type FiniteStateMachine struct {
	states       map[string]struct{}
	initialState string
	alphabet     map[rune]struct{}
	transitions  TransitionsMap
	finalStates  map[string]struct{}
}

func New(config Config) (*FiniteStateMachine, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	newFSM := FiniteStateMachine{
		states:       config.States,
		initialState: *config.initialState,
		alphabet:     config.Alphabet,
		transitions:  config.Transitions,
	}

	newFSM.finalStates = make(map[string]struct{}, len(config.finalStates))
	for _, currentFinalState := range config.finalStates {
		newFSM.finalStates[currentFinalState] = struct{}{}
	}

	return &newFSM, nil
}

func (f *FiniteStateMachine) Process(input string) (*string, bool) {
	if len(input) == 0 {
		return nil, false
	}
	currentState := f.initialState

	for _, currentRune := range input {
		_, ok := f.alphabet[currentRune]
		if !ok {
			return nil, false
		}

		// Accepted input, get the current transition for this state
		inputMap, ok := f.transitions.transitions[currentState]
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
	if _, ok := f.finalStates[currentState]; !ok {
		return nil, false
	}

	return &currentState, true
}
