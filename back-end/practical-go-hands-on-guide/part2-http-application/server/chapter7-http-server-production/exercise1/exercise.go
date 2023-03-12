package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

const PORT = "7166"

// Mock Server Side Request
func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling Ping!")
	fmt.Fprintln(w, "pong")
}

// Exercise 7.1
func createHttpRequestWithTrace(
	ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	trace := &httptrace.ClientTrace{
		// DNS
		DNSStart: func(dnsStartInfo httptrace.DNSStartInfo) {
			fmt.Printf("DNS Start Info: %+v\n", dnsStartInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Done Info: %+v\n", dnsInfo)
		},

		// TCP Connection
		TLSHandshakeStart: func() {
			fmt.Printf("TLS HandShake Start\n")
		},
		TLSHandshakeDone: func(connState tls.ConnectionState, err error) {
			fmt.Printf("TLS HandShake Done\n")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Connection from Pool: %+v\n", connInfo)
		},
		PutIdleConn: func(err error) {
			fmt.Printf("Put Idle Conn Error: %+v\n", err)
		},
	}

	ctxTrace := httptrace.WithClientTrace(req.Context(), trace)
	request := req.WithContext(ctxTrace)
	return request, nil
}

func generateHandleUserApi(useTrace bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start Processing the Request")
		doSomethingInProcess()

		// 서버 작업 이후, 다른 서버로 네트워크 요청을 위한 준비
		url := fmt.Sprintf("http://localhost:%s/ping", PORT)
		method := http.MethodGet
		var req *http.Request
		var err error
		if useTrace {
			req, err = createHttpRequestWithTrace(
				r.Context(),
				method, url,
				nil,
			)
		} else {
			req, err = http.NewRequestWithContext(
				r.Context(),
				method, url,
				nil,
			)
		}

		if err != nil {
			http.Error(
				w, err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		client := &http.Client{}
		log.Println("Outgoing HTTP Request Start...")

		// 다른 서버로 네트워크 요청
		res, err := client.Do(req)
		if err != nil {
			log.Printf("Error in outgoing request: %v\n", err)
			http.Error(
				w, err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
		defer res.Body.Close()

		data, _ := io.ReadAll(res.Body)
		fmt.Fprint(w, string(data))
		log.Println("Finished Processing the Request")
	}
}

func doSomethingInProcess() {
	log.Println("Start - doSomethingInProcess")
	time.Sleep(2 * time.Second)
	log.Println("End - doSomethingInProcess")
}

func main() {
	userHandler := http.HandlerFunc(generateHandleUserApi(false))
	timeoutHandler := http.TimeoutHandler(userHandler, 1*time.Second, "Out of Time...")

	userHandlerWithTrace := http.HandlerFunc(generateHandleUserApi(true))
	timeoutHandlerWithTrace1 := http.TimeoutHandler(userHandlerWithTrace, 1*time.Second, "Out of Time")
	timeoutHandlerWithTrace2 := http.TimeoutHandler(userHandlerWithTrace, 3*time.Second, "Out of Time")

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/api", timeoutHandler)
	mux.Handle("/api/trace/1", timeoutHandlerWithTrace1)
	mux.Handle("/api/trace/2", timeoutHandlerWithTrace2)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux))
}
