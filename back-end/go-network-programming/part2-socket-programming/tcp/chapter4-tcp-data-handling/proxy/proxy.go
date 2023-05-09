package proxy

import (
	"io"
	"log"
	"net"
)

// ProxyConnection only useful when src and dst hosts are listening with their own port
// => Hard to Test
func ProxyConnection(src, dst string) error {
	connSrc, err := net.Dial("tcp", src)
	if err != nil {
		return err
	}
	defer func() {
		_ = connSrc.Close()
	}()

	connDst, err := net.Dial("tcp", dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = connDst.Close()
	}()

	// reverse proxy: dst response -> src
	// => 포워딩 이후에 이루어지는 작업이기 때문에 고루틴
	go func() {
		// 네트워크 연결 에러 발생 시, io.Copy 자체가 자동 종료
		// => 고루틴이 에러 발생 시, 알아서 종료됨
		_, _ = io.Copy(connSrc, connDst)
	}()

	// forwarding proxy: src request -> dst
	_, err = io.Copy(connDst, connSrc)
	return err
}

// ProxyGeneral general version of proxy
// : can use with os.File, net.Conn etc
// => Easy to test
func ProxyGeneral(src io.Reader, dst io.Writer) error {
	// 리버스 프록시를 위해서는 두 인터페이스를 모두 구현하고 있어야 함
	writerSrc, isSrcWriter := src.(io.Writer)
	readerDst, isDstReader := dst.(io.Reader)

	if isSrcWriter && isDstReader {
		go func() {
			log.Println("Reverse Proxy!")
			_, _ = io.Copy(writerSrc, readerDst)
		}()
	}

	log.Println("Forwarding Proxy!")
	_, err := io.Copy(dst, src)
	return err
}
