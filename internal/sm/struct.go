package sm

type StateMachine struct {
	states map[int]StateAttributes
}

type StateAttributes struct {
	f           func(sa *CurrentState) error
	finalStates []int
}

type CurrentState struct {
	State  int
	Entity interface{}
}
