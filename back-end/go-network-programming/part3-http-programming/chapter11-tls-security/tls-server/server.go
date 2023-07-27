package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func NewTLSServer(ctx context.Context, address string,
	maxIdle time.Duration, tlsConfig *tls.Config) *Server {
	return &Server{
		ctx:       ctx,
		ready:     make(chan struct{}),
		addr:      address,
		maxIdle:   maxIdle,
		tlsConfig: tlsConfig,
	}
}

type Server struct {
	ctx       context.Context
	ready     chan struct{}
	addr      string
	maxIdle   time.Duration
	tlsConfig *tls.Config
}

func (s *Server) Ready() {
	if s.ready != nil {
		<-s.ready
	}
}

func (s *Server) ListenAndServeTLS(certFilename, keyFilename string) error {
	if s.addr == "" {
		s.addr = "localhost:443"
	}

	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("binding to tcp %s: %w", s.addr, err)
	}

	if s.ctx != nil {
		go func() {
			// 컨텍스트 만료 시, 서버 리스너 종료를 위한 고루틴
			<-s.ctx.Done()
			_ = l.Close()
		}()
	}

	return s.ServeTLS(l, certFilename, keyFilename)
}

func (s *Server) ServeTLS(l net.Listener, certFilename, keyFilename string) error {
	if s.tlsConfig == nil {
		s.tlsConfig = &tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
			// TLS 핸드셰이크 암호화 협상 단계에서 클라이언트의 선호 스위트를 기다리지 않고 서버 측 스위트 사용
			PreferServerCipherSuites: true,
		}
	}

	// 인증서가 없는 경우, 로컬 파일 시스템에 있는 인증서 등록
	if len(s.tlsConfig.Certificates) == 0 && s.tlsConfig.GetCertificate == nil {
		cert, err := tls.LoadX509KeyPair(certFilename, keyFilename)
		if err != nil {
			return fmt.Errorf("loading key pair: %v\n", err)
		}
		s.tlsConfig.Certificates = []tls.Certificate{cert}
	}

	// TLS 리스터 생성
	tlsListener := tls.NewListener(l, s.tlsConfig)

	// 리스너 생성 후, 요청 처리가 가능하므로 ready 상태 업데이트
	if s.ready != nil {
		close(s.ready)
	}

	// 요청 처리를 위한 고루틴
	for {
		conn, err := tlsListener.Accept()
		if err != nil {
			return fmt.Errorf("accept: %v", err)
		}

		go func() {
			defer func() { _ = conn.Close() }()

			for {
				if s.maxIdle > 0 {
					err := conn.SetDeadline(time.Now().Add(s.maxIdle))
					if err != nil {
						log.Printf("[error] set deadline: %v\n", err)
						return
					}
				}

				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					log.Printf("[error] read request: %v\n", err)
					return
				}

				req := string(buf[:n])
				_, err = conn.Write([]byte(fmt.Sprintf("[Echo] %s", req)))
				if err != nil {
					log.Printf("[error] write response: %v\n", err)
					return
				}
			}
		}()
	}
}
