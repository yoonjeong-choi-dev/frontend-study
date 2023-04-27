package raft_concensus

import (
	"github.com/hashicorp/raft"
	"io"
)

// FSM raft.FSM 인터페이스 구현체
type FSM struct {
	state state
}

// Apply raft.log 데이터로 상태 전환
func (f *FSM) Apply(log *raft.Log) interface{} {
	f.state.Transition(state(log.Data))
	return string(f.state)
}

// Snapshot 단순히 인터페이스 요구를 충족하기 위한 함수
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore 단순히 인터페이스 요구를 충족하기 위한 함수
func (f *FSM) Restore(snapshot io.ReadCloser) error {
	return nil
}

func NewFSM() *FSM {
	return &FSM{state: first}
}
