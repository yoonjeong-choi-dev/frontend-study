package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// options base on "ping"
var (
	count    = flag.Int("c", 3, "number of pings: <= 0 means forever")
	interval = flag.Duration("i", time.Second, "interval between pings")
	timeout  = flag.Duration("W", 5*time.Second, "time to wait for a reply")
)

func setupHelpMessage() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] host:port\nOPtions:\n", os.Args)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Print("host:port is required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	target := flag.Arg(0)
	fmt.Println("PING", target)

	if *count <= 0 {
		fmt.Println("Press CTRL+C to stop")
	}

	numTrial := 0
	for (*count <= 0) || (numTrial < *count) {
		numTrial++
		fmt.Print(numTrial, " ")

		// create tcp connection
		start := time.Now()
		conn, err := net.DialTimeout("tcp", target, *timeout)
		dur := time.Since(start)

		if err != nil {
			fmt.Printf("fail in %s: %v\n", dur, err)

			// 일시적인 에러가 아닌 경우 종료. 일시적 에러인 경우 재시도
			if netErr, ok := err.(net.Error); !ok || !netErr.Temporary() {
				os.Exit(1)
			}
		} else {
			// 응답 받은 경우 연결 종료
			_ = conn.Close()
			fmt.Println(dur)
		}

		time.Sleep(*interval)
	}
}
