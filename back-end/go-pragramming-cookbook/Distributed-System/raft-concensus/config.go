package raft_concensus

import "github.com/hashicorp/raft"

func Config(num int) {
	rs := getRaftSet(num)

	// 클러스터를 구성하는 서버들 간 통신 연결
	// => 통신을 통해 클러스터는 현재 상태에 대해서 리더 선임을 위한 투표 진행
	for _, from := range rs {
		for _, to := range rs {
			from.Transport.Connect(to.Transport.LocalAddr(), to.Transport)
		}
	}

	// 서버 실행
	for _, r := range rs {
		if err := raft.BootstrapCluster(r.Config, r.Store, r.Store,
			r.SnapShotStore, r.Transport, r.Configuration); err != nil {
			panic(err)
		}

		newRaft, err := raft.NewRaft(r.Config, r.FSM, r.Store, r.Store, r.SnapShotStore, r.Transport)
		if err != nil {
			panic(err)
		}

		rafts[r.Transport.LocalAddr()] = newRaft
	}
}
