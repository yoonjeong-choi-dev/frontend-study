// go:build darwin || linux

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
	dir, err := ioutil.TempDir("", "echo_unixgram")
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

	// 서버 소켓 파일 이름
	serverSocket := filepath.Join(dir, fmt.Sprintf("server%d.sock", os.Getpid()))

	// unixgram 네트워크 타입으로 서버 생성
	serverAddr, err := echoServer.DatagramEchoServer(ctx, "unixgram", serverSocket)
	if err != nil {
		panic(err)
	}

	// 소켓 파일 권한 확인
	err = os.Chmod(serverSocket, os.ModeSocket|0622)
	if err != nil {
		panic(err)
	}

	// 클라이언트 소켓 파일 이름
	// => UDP 통신처럼 각 호스트는 고유의 소켓 필요
	clientSocket := filepath.Join(dir, fmt.Sprintf("pinning-cert-client%d.sock", os.Getpid()))
	client, err := net.ListenPacket("unixgram", clientSocket)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// 소켓 파일 권한 확인
	err = os.Chmod(clientSocket, os.ModeSocket|0622)
	if err != nil {
		panic(err)
	}

	reqCount := 3
	for i := 0; i < reqCount; i++ {
		msg := fmt.Sprintf("ping-%d!", i+1)
		_, err = client.WriteTo([]byte(msg), serverAddr)
		if err != nil {
			panic(err)
		}
	}

	buf := make([]byte, 1024)
	for i := 0; i < reqCount; i++ {
		size, addr, err := client.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Response %d: %s\n", i, string(buf[:size]))
		fmt.Printf("Server Addr: %s\nResponse Addr: %s\n\n", serverAddr.String(), addr.String())
	}
}
