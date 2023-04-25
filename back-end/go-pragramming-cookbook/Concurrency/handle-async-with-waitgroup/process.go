package waitgroup

import (
	"log"
	"sync"
	"time"
)

// Crawl 사이트 주소들에 대해 요청을 하고, 모든 요청의 응답 코드 반환
func Crawl(sites []string) ([]int, error) {
	start := time.Now()
	log.Println("[Crawl] Start to crawl")

	wg := &sync.WaitGroup{}

	var codes []int
	ce := &CrawlError{}

	for _, v := range sites {
		// 각 사이트에 대해서 비동기적으로 요청 수행
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			res, err := GetURL(url)
			if err != nil {
				ce.Add(err)
				return
			}

			codes = append(codes, res.StatusCode)
		}(v)
	}

	// 모든 비동기 처리가 완료될 때까지 대기
	wg.Wait()

	// 요청들에 대한 에러가 있는 경우 출력
	if ce.CanPresent() {
		return codes, ce
	}
	
	log.Printf("[Crawl] Completed to crawl in %s\n", time.Since(start))
	return codes, nil
}
