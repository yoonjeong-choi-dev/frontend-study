package main

import (
	"context"
	echoServer "echo-server"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic stacktrace:\n%s", string(debug.Stack()))
		}
	}()

	// 에코 서버의 소켓 파일을 저장하기 위한 디렉터리 생성
	dir, err := ioutil.TempDir("", "echo_unix")
	if err != nil {
		panic(err)
	}
	defer func() {
		if rErr := os.RemoveAll(dir); rErr != nil {
			panic(rErr)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 소켓 프로그래밍을 위한 소켓 파일 이름
	// 연결 기반(스트리밍) 통신이기 때문에 net.Conn 에 대응하는 하나의 소켓만 필요 like TCP
	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))

	// 서버 주소로 소켓 파일 경로 전달
	// => 소켓 파일 생성
	serverAddr, err := echoServer.StreamingEchoServer(ctx, "unix", socket)
	if err != nil {
		panic(err)
	}

	// 소켓 파일 권한 확인
	err = os.Chmod(socket, os.ModeSocket|0666)
	if err != nil {
		panic(err)
	}

	client, err := net.Dial("unix", serverAddr.String())
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	reqCount := 3
	for i := 0; i < reqCount; i++ {
		msg := fmt.Sprintf("ping-%d!", i+1)
		_, err = client.Write([]byte(msg))
		if err != nil {
			panic(err)
		}
	}

	buf := make([]byte, 1024)
	size, err := client.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server Address: %s\n", serverAddr.String())
	fmt.Printf("Response: %s\n", string(buf[:size]))
}
