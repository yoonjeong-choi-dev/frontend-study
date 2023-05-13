package tftp_server

import (
	"bytes"
	"errors"
	"log"
	"net"
	"time"
)

// Server 네트워크 인터페이스와 사용자 정의 패킷 사이의 파이프 역할
type Server struct {
	Payload []byte
	Retries uint8
	Timeout time.Duration
}

func (s *Server) ListenAndServe(addr string) error {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil
	}
	defer func() { _ = conn.Close() }()

	log.Printf("Listening on %s....\n", addr)

	return s.Serve(conn)
}

func (s *Server) Serve(conn net.PacketConn) error {
	if conn == nil {
		return errors.New("nil connection")
	}

	if s.Payload == nil {
		return errors.New("payload is required")
	}

	if s.Retries == 0 {
		s.Retries = 10
	}

	if s.Timeout == 0 {
		s.Timeout = 6 * time.Second
	}

	// 현재는 파일 다운로드 요청만 처리
	var rrq ReadRequest

	for {
		buf := make([]byte, DatagramSize)

		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			return err
		}

		err = rrq.UnmarshalBinary(buf)
		if err != nil {
			log.Printf("[%s] bad request: %v\n", addr, err)
			continue
		}

		// 고루틴으로 현재 요청 처리
		go s.handle(addr.String(), rrq)
	}
}

func (s *Server) handle(clientAddr string, rrq ReadRequest) {
	log.Printf("[%s] requested file: %s\n", clientAddr, rrq.Filename)

	conn, err := net.Dial("udp", clientAddr)
	if err != nil {
		log.Printf("[%s] dial: %v\n", clientAddr, err)
		return
	}
	defer func() { _ = conn.Close() }()

	var (
		ackPkt  Ack
		errPkt  Err
		dataPkt = Data{Payload: bytes.NewReader(s.Payload)}
		buf     = make([]byte, DatagramSize)
	)

NEXTPACKET:
	// 마지막 패킷 전송까지 반복
	for n := DatagramSize; n == DatagramSize; {
		data, err := dataPkt.MarshalBinary()
		if err != nil {
			log.Printf("[%s] preparing data packet: %v", clientAddr, err)
			return
		}

	RETRY:
		for i := s.Retries; i > 0; i-- {
			n, err = conn.Write(data)
			if err != nil {
				log.Printf("[%s] write: %v\n", clientAddr, err)
				return
			}

			// 클라이언트 ACK 신호 대기
			_ = conn.SetReadDeadline(time.Now().Add(s.Timeout))

			_, err = conn.Read(buf)
			if err != nil {
				// 일시적인 에러인 경우, 재시도
				if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
					continue RETRY
				}

				log.Printf("[%s] waiting for ACK: %v\n", clientAddr, err)
				return
			}

			switch {
			case ackPkt.UnmarshalBinary(buf) == nil:
				// 현재 전송한 패킷 블럭에 대한 ACK 패킷인 경우 다음 패킷 전송
				if uint16(ackPkt) == dataPkt.Block {
					continue NEXTPACKET
				}
			case errPkt.UnmarshalBinary(buf) == nil:
				log.Printf("[%s] received error: %v\n", clientAddr, errPkt.Message)
				return
			default:
				// 현재 데이터 블록 재전송
				log.Printf("[%s] bad packet\n", clientAddr)
			}
		}
		log.Printf("[%s] exhausted retreis\n", clientAddr)
		return
	}

	log.Printf("[%s] send %d blocks", clientAddr, dataPkt.Block)
}
