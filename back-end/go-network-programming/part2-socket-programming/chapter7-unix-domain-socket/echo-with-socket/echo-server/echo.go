package echo_server

import (
	"context"
	"fmt"
	"net"
	"os"
)

// StreamingEchoServer TCP 통신처럼 스트림 기반의 연결
func StreamingEchoServer(ctx context.Context, network string, addr string) (net.Addr, error) {
	server, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	go func() {
		// 컨텍스트 종료 시 서버를 종료하기 위한 고루틴
		go func() {
			<-ctx.Done()
			_ = server.Close()
		}()

		for {
			conn, err := server.Accept()
			if err != nil {
				return
			}

			go func() {
				defer func() { _ = conn.Close() }()

				for {
					buf := make([]byte, 1024)
					size, err := conn.Read(buf)
					if err != nil {
						return
					}

					msg := fmt.Sprintf("[echo] %s", string(buf[:size]))
					_, err = conn.Write([]byte(msg))
					if err != nil {
						return
					}
				}
			}()
		}
	}()

	return server.Addr(), nil
}

// DatagramEchoServer UDP 통신처럼 데이터그램 기반의 통신
func DatagramEchoServer(ctx context.Context, network string, addr string) (net.Addr, error) {
	server, err := net.ListenPacket(network, addr)
	if err != nil {
		return nil, err
	}

	go func() {
		// 컨텍스트 종료 시 서버를 종료하기 위한 고루틴
		go func() {
			<-ctx.Done()
			_ = server.Close()

			// 리스너를 net.Listen, net.ListenUnix 함수로 생성하지 않는 경우
			// 매뉴얼한 소켓 파일 삭제 필요
			if network == "unixgram" {
				_ = os.Remove(addr)
			}
		}()

		buf := make([]byte, 1024)
		for {
			size, clientAddr, err := server.ReadFrom(buf)
			if err != nil {
				return
			}

			msg := fmt.Sprintf("[echo] %s", string(buf[:size]))
			_, err = server.WriteTo([]byte(msg), clientAddr)
			if err != nil {
				return
			}
		}
	}()

	return server.LocalAddr(), nil
}
