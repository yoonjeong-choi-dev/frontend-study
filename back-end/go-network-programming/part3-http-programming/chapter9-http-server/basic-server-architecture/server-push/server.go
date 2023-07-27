package main

import (
	"basic-server-architecture/handlers"
	"basic-server-architecture/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

var (
	addr       = "127.0.0.1:7166"
	cert       = "./certs/server.crt"
	privateKey = "./certs/server.key"
	staticDir  = "../files"
)

func main() {
	mux := http.NewServeMux()

	// Handler 1. 정적 리소스 핸들러
	mux.Handle("/static/", http.StripPrefix("/static/",
		middleware.RestrictPrefix(".", http.FileServer(http.Dir(staticDir))),
	))

	// Handler 2. HTML 핸들러
	mux.Handle("/",
		handlers.Methods{
			http.MethodGet: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					if pusher, ok := w.(http.Pusher); ok {

						options := &http.PushOptions{
							Header: http.Header{
								"Accept-Encoding": r.Header["Accept-Encoding"],
							},
						}

						// index.html 에 존재하는 정적 리소스 서버 푸시
						// 이 때, 정적 리소스 경로는 서버가 아닌 클라이언트 요청 기준
						// => /static/ 경로로 처리
						targets := []string{
							"/static/style.css",
							"/static/hiking.svg",
						}

						for _, target := range targets {
							if err := pusher.Push(target, options); err != nil {
								log.Printf("%s push failed: %#v\n", target, err)
							}
						}
					}

					http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
				},
			),
		},
	)

	// Handler 3.
	mux.Handle("/2", handlers.Methods{
		http.MethodGet: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, filepath.Join(staticDir, "index2.html"))
			}),
	})

	handler := middleware.Logger(mux)

	// server setting
	server := &http.Server{
		Addr:              addr,
		Handler:           handler,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
	}

	// Graceful Termination
	done := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		for {
			if <-c == os.Interrupt {
				log.Println("Server is shutdown gracefully...")
				if err := server.Shutdown(context.Background()); err != nil {
					log.Printf("Shutdown : %v\n", err)
				}
				close(done)
				return
			}
		}
	}()

	log.Printf("Serving files in %q over %s\n", staticDir, server.Addr)

	// https://localhost:7166/ 로 접속 필요
	fmt.Println(server.ListenAndServeTLS(cert, privateKey))
	//fmt.Println(server.ListenAndServe())
	<-done
}
