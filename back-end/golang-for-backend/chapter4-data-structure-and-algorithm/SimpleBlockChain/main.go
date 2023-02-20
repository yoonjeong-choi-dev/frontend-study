package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"simpleblockchain/blockchain"
	"time"
)

// pow : Proof of Work(자격 증명)
// 블록체인(prefix)에 추가 가능한 요소(현재는 문자열)를 찾는 작업
func pow(prefix string, bitLength int) {
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
			fmt.Println("Block Chain: ", string(randomBytes))
			break
		}
	}

	end := time.Now()
	fmt.Println("time(s):", end.Sub(start).Seconds())
	fmt.Println("processed", humanize.Comma(int64(totalHashingProcessed)))
	fmt.Printf("processed/sec: %s\n", humanize.Comma(int64(float64(totalHashingProcessed)/end.Sub(start).Seconds())))
}

func main() {
	pow("Tanmay Bakshi + Baheer Kamal", 24)
}
