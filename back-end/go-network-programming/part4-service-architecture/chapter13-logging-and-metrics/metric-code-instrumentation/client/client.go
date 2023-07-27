package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"metric-code-instrumentation/metrics"
	"net/http"
	"sync"
)

const (
	metricsAddr = "127.0.0.1:8081"
	webAddr     = "127.0.0.1:7166"
)

func main() {
	numClient := 500
	numReq := 100
	wg := new(sync.WaitGroup)

	for i := 0; i < numClient; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &http.Client{
				// numClient 개수만큼 다른 TCP 연결을 사용하기 위해 복제를 해줌
				Transport: http.DefaultTransport.(*http.Transport).Clone(),
			}

			for j := 0; j < numReq; j++ {
				path := "good"
				if j%3 == 0 {
					path = "bad"
				}
				res, _ := client.Get(fmt.Sprintf("http://%s/%s", webAddr, path))
				if res != nil {
					_, _ = io.Copy(ioutil.Discard, res.Body)
					_ = res.Body.Close()
				}
			}
		}()
	}

	// Wait all request
	wg.Wait()

	fmt.Print("All Request Done\n\n")

	// Get all metrics
	res, err := http.Get(fmt.Sprintf("http://%s/metrics", metricsAddr))
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = res.Body.Close()

	metricPrefix := []byte(fmt.Sprintf("%s_%s", *metrics.Namespace, *metrics.Subsystem))
	fmt.Println("Current Metrics:")
	for _, line := range bytes.Split(data, []byte("\n")) {
		if bytes.HasPrefix(line, metricPrefix) {
			fmt.Printf("%s\n", line)
		}
	}
}
