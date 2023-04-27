package raft_concensus

type state string

const (
	first  state = "first"
	second state = "second"
	third  state = "third"
)

// allowedState 각 상태에서 변경이 가능한 상태들 정보
var allowedState map[state][]state

func init() {
	allowedState = make(map[state][]state)
	allowedState[first] = []state{second, third}
	allowedState[second] = []state{third}
	allowedState[third] = []state{first}
}

func (s *state) CanTransition(next state) bool {
	for _, n := range allowedState[*s] {
		if n == next {
			return true
		}
	}
	return false
}

func (s *state) Transition(next state) {
	if s.CanTransition(next) {
		*s = next
	}
}
