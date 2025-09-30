# Finite State Machine package

## Basic Info

FSM is an implimentation of a 'Finite State Machine'. 

The FSM machine is split up into distinct parts parts:

Transition: Represents an allowable transition from one state to another: (State,Input) -> State

TransitionsMap: Holds all the allowable Transitions in a Finite State Machine. Verifies that for each state, there is a transition that allows it to take any input in the FSM alphabet.

Config: Holds a transition Map, and handles verification of states and initial conditions. Used as input for the Finite State Machien struct

FiniteStateMachine: The actual Finite State Machine has only one method: Process(input string). This returns the final state of the FSM when the input is processed, and a boolean.
If the Input is acceptable, it returns the final state and 'true' for the boolean.
If the input is not acceptable (is empty, contains characters not in the FSM alphabet, or the final state isn't acceptable) it returns nil for the state, and the boolean is returned as false.

## Mod3 example

The 'Mod3' finite state machine is given as an example.
**Go to the pkg\fsm\fsm_test.go file to see the mod3 finites state machine example** 

The example is given as a test file: Initial states, alphabet, and transitions are set up, and then used to create a config for the Finite State Machine. The tests simulate what the Mod3 FSM would return, given the test input.

In order to checkout and run the tests:
- Clone the repository: 'git clone https://github.com/Manuel9550/FSM.git'
- Change directory to where you've checked out the repo
- At the head of the repo, run 'go mod download'
- Change directory to .\pkg\fsm\ in the repo
- Run 'go test'





