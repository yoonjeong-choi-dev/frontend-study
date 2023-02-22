package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/dustin/go-humanize"
	"runtime"
	"simpleblockchain/blockchain"
	"time"
)

func powMultiThread(prefix string, bitLength int) {
	start := time.Now()
	var hashingProcessed []int

	// numCPU: CPU 코어 수보다 고루틴을 더 생성하여 불필요한 컨텍스트 스위칭를 방지하기 위함
	numCPU := runtime.NumCPU()

	// closeChan: 특정 고루틴에서 정답을 찾은 경우, 모든 고루틴을 종료하기 위한 채널
	closeChan := make(chan int, 1)

	// solutionChan: 특정 고루틴에서 정답 찾은 경우, 해당 정답을 main 고루틴에게 전달하기 위한 채널
	solutionChan := make(chan []byte, 1)

	// 자격 증명을 완료한 문자열을 찾는 고루틴들 실행
	for i := 0; i < numCPU; i++ {
		hashingProcessed = append(hashingProcessed, 0)

		// powSingleThread 부분과 로직이 거의 비슷
		go func(threadIdx int) {
			seed := uint64(time.Now().Local().UnixNano())
			randomBytes := make([]byte, 20)
			randomBytes = append([]byte(prefix), randomBytes...)

			for {
				select {
				case <-closeChan:
					// 현재 고루틴 종료 전에, 아직 실행 중인 고루틴 하나에게 종료 신호를 전파
					closeChan <- 1
					return
				case <-time.After(time.Nanosecond):
					count := 0
					for count < 5000 {
						count++
						seed = blockchain.RandomString(randomBytes, len(prefix), seed)
						if blockchain.Consensus(randomBytes, bitLength) {
							// 정답 찾으면 메인 고루틴 및 다른 서브 고루틴에게 통보
							hashingProcessed[threadIdx] += count
							solutionChan <- randomBytes
							closeChan <- 1
							return
						}
					}
					hashingProcessed[threadIdx] += count
				}
			}
		}(i)
	}

	solution := <-solutionChan

	totalHashingProcessed := 0
	for _, v := range hashingProcessed {
		totalHashingProcessed += v
	}

	end := time.Now()
	fmt.Println("Block Chain: ", string(solution))
	fmt.Println("Hashed:", sha256.Sum256(solution))
	fmt.Println("time(s):", end.Sub(start).Seconds())
	fmt.Println("processed", humanize.Comma(int64(totalHashingProcessed)))
	fmt.Printf("processed/sec: %s\n", humanize.Comma(int64(float64(totalHashingProcessed)/end.Sub(start).Seconds())))
}

// powSingleThread : Proof of Work(자격 증명)
// 블록체인(prefix)에 추가 가능한 요소(현재는 문자열)를 찾는 작업
func powSingleThread(prefix string, bitLength int) {
	start := time.Now()

	totalHashingProcessed := 0
	seed := uint64(time.Now().Local().UnixNano())
	randomBytes := make([]byte, 20)
	randomBytes = append([]byte(prefix), randomBytes...)

	// goal :randomBytes 에 대한 자격 증명
	for {
		totalHashingProcessed++
		seed = blockchain.RandomString(randomBytes, len(prefix), seed)
		if blockchain.Consensus(randomBytes, bitLength) {
			fmt.Println("Block Chain:", string(randomBytes))
			fmt.Println("Hashed:", sha256.Sum256(randomBytes))
			break
		}
	}

	end := time.Now()
	fmt.Println("time(s):", end.Sub(start).Seconds())
	fmt.Println("processed", humanize.Comma(int64(totalHashingProcessed)))
	fmt.Printf("processed/sec: %s\n", humanize.Comma(int64(float64(totalHashingProcessed)/end.Sub(start).Seconds())))
}
