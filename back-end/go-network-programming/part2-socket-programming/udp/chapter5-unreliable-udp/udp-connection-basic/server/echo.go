package server

import (
	"context"
	"fmt"
	"net"
)

func EchoUDPServer(ctx context.Context, addr string) (net.Addr, error) {
	// packet 지향적인 net.PacketConn 인터페이스 사용
	server, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	// listening
	go func() {
		// 컨텍스트가 취소되면 서버를 닫는 고루틴
		go func() {
			// ctx.Done 에 의해 블로킹
			<-ctx.Done()
			_ = server.Close()
		}()

		buf := make([]byte, 1024)
		for {
			// 세션 기능이 없으므로, udp 연결 객체로 부터 송신자(클라이언트) 주소를 받아와야 함
			size, clientAddr, err := server.ReadFrom(buf)
			if err != nil {
				return
			}

			// 에코 서버
			resMsg := fmt.Sprintf("[Echo] %s", string(buf[:size]))

			// 세션 기능이 없으므로, 명시적으로 송신자(클라이언트)의 주소를 전달
			_, err = server.WriteTo([]byte(resMsg), clientAddr)
			if err != nil {
				return
			}
		}
	}()

	return server.LocalAddr(), nil
}
