package raft_concensus

import (
	"fmt"
	"github.com/hashicorp/raft"
)

// rafts 각 서버를 구성하는 raft 정보 저장
var rafts map[raft.ServerAddress]*raft.Raft

func init() {
	rafts = make(map[raft.ServerAddress]*raft.Raft)
}

// raftSet raft 패키지에서 필요한 정보들을 저장하는 구조체
// 클러스터를 구성하는 각 서버(리더 및 추종자)의 설정
type raftSet struct {
	Config        *raft.Config
	Store         *raft.InmemStore
	SnapShotStore raft.SnapshotStore
	FSM           *FSM
	Transport     raft.LoopbackTransport
	Configuration raft.Configuration
}

// getRaftSet raft 클러스터에 해당하는 raftSet 슬라이스 반환
// => 인메모리 raft 클러스터로 사용
func getRaftSet(num int) []*raftSet {
	rs := make([]*raftSet, num)
	servers := make([]raft.Server, num)

	for i := 0; i < num; i++ {
		addr := raft.ServerAddress(fmt.Sprint(i))
		_, transport := raft.NewInmemTransport(addr)

		servers[i] = raft.Server{
			Suffrage: raft.Voter,
			ID:       raft.ServerID(addr),
			Address:  addr,
		}

		config := raft.DefaultConfig()
		config.LocalID = raft.ServerID(addr)
		rs[i] = &raftSet{
			Config:        config,
			Store:         raft.NewInmemStore(),
			SnapShotStore: raft.NewInmemSnapshotStore(),
			FSM:           NewFSM(),
			Transport:     transport,
		}
	}

	// 클러스터를 구성하는 서버들의 정보를 모든 서버 구성에 저장
	for _, r := range rs {
		r.Configuration = raft.Configuration{Servers: servers}
	}

	return rs
}
