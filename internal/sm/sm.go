package sm

import (
	"errors"
	"fmt"
)

func New() *StateMachine {
	return &StateMachine{
		states: make(map[int]StateAttributes),
	}
}

func (s *StateMachine) CreateRoute(startState int, funcToCall func(sa *CurrentState) error, finalStates ...int) error {
	if s.states[startState].f != nil {
		return errors.New("state transition already defined")
	}

	if funcToCall == nil || finalStates == nil || len(finalStates) == 0 {
		return errors.New("route attributes are missing")
	}

	s.states[startState] = StateAttributes{
		f:           funcToCall,
		finalStates: finalStates,
	}

	return nil
}

func (s *StateMachine) Run(sa *CurrentState) error {
	if sa == nil {
		return errors.New("no start state defined")
	}

	if len(s.states) == 0 || s.states[sa.State].f == nil {
		return errors.New("no transitions defined for this start state")
	}

	currState := sa.State

	if err := s.states[currState].f(sa); err != nil {
		return fmt.Errorf("sm returned an error; err=%+v", err)
	}

	finalState := sa.State

	eligibleFinalState := false
	for _, endState := range s.states[currState].finalStates {
		if finalState == endState {
			eligibleFinalState = true
			break
		}
	}

	if !eligibleFinalState {
		return errors.New("function moved to a non-initialized final state from the provider state date")
	}

	if err := s.Save(sa); err != nil {
		return fmt.Errorf("error returned while trying to save the updated state; err=%+v", err)
	}

	return nil
}

func (s *StateMachine) Save(sa *CurrentState) error {

	// Save to DB

	return nil
}
