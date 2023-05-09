package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"runtime/debug"
)

type Monitoring struct {
	*log.Logger
}

// Write implements io.Writer interface
func (m *Monitoring) Write(payload []byte) (int, error) {
	return len(payload), m.Output(2, string(payload))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("stacktrace:\n%s", string(debug.Stack()))
		}
	}()

	monitor := &Monitoring{Logger: log.New(os.Stdout, "[monitor] ", 0)}

	server, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}

	serverDone := make(chan interface{})
	go func() {
		defer close(serverDone)

		conn, err := server.Accept()
		if err != nil {
			return
		}
		defer func() { _ = conn.Close() }()

		// 클라이언트 요청을 읽어 모니터링 객체에 씀
		reader := io.TeeReader(conn, monitor)
		buffer := make([]byte, 1024)
		size, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}

		// 클라이언트 연결 및 모니터링 객체에 모두 동일한 바이트를 쓸 수 있게 해줌
		writer := io.MultiWriter(conn, monitor)
		resMsg := fmt.Sprintf("[Echo] %s", string(buffer[:size]))
		_, err = writer.Write([]byte(resMsg))
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}
	}()

	client, err := net.Dial("tcp", server.Addr().String())
	if err != nil {
		panic(err)
	}

	_, err = client.Write([]byte("Monitor Test!"))
	if err != nil {
		monitor.Println(err)
		panic(err)
	}

	_ = client.Close()
	<-serverDone
}
