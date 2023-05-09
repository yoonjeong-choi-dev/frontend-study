package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"proxy"
	"runtime/debug"
	"sync"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	var wg sync.WaitGroup

	// Server
	server, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}

	wg.Add(1)

	// server listening
	go func() {
		defer wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				return
			}

			// server handler
			go func(c net.Conn) {
				defer func() {
					_ = c.Close()
				}()

				for {
					buf := make([]byte, 1024)
					size, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return
					}

					switch msg := string(buf[:size]); msg {
					case "ping":
						_, err = c.Write([]byte("pong"))
					default:
						_, err = c.Write([]byte(fmt.Sprintf("[Echo] %s", msg)))
					}

					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return
					}
				}
			}(conn)
		}
	}()

	// Proxy
	proxyServer, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}

	wg.Add(1)

	// proxy server listening
	go func() {
		defer wg.Done()

		for {
			conn, err := proxyServer.Accept()
			if err != nil {
				return
			}

			// proxy handler
			go func(from net.Conn) {
				defer func() { from.Close() }()

				// pinning-cert-client -> from(conn) -> proxy -> to(conn) ->server
				to, err := net.Dial("tcp", server.Addr().String())
				if err != nil {
					log.Println(err)
					return
				}
				defer func() { _ = to.Close() }()

				err = proxy.ProxyGeneral(from, to)
				if err != nil && err != io.EOF {
					log.Println(err)
				}
			}(conn)
		}
	}()

	// Client
	client, err := net.Dial("tcp", proxyServer.Addr().String())
	if err != nil {
		panic(err)
	}

	msgs := []string{"ping", "pong", "test"}
	for _, m := range msgs {
		_, err = client.Write([]byte(m))
		if err != nil {
			panic(err)
		}

		buf := make([]byte, 1024)
		size, err := client.Read(buf)
		if err != nil {
			panic(err)
		}

		res := string(buf[:size])
		fmt.Printf("%s -> proxy -> %s\n", m, res)
	}

	// close all connection and servers
	_ = client.Close()
	_ = proxyServer.Close()
	_ = server.Close()

	wg.Wait()
}
