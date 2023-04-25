package waitgroup

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetURL(url string) (*http.Response, error) {
	start := time.Now()
	log.Printf("[GetURL] Start with %s\n", url)
	res, err := http.Get(url)
	log.Printf("[GetURL] Completed to get %s in %s\n", url, time.Since(start))

	return res, err
}

type CrawlError struct {
	Errors []string
}

func (c *CrawlError) Add(err error) {
	c.Errors = append(c.Errors, err.Error())
}

// Error error 인터페이스 구현
func (c *CrawlError) Error() string {
	return fmt.Sprintf("All errors: %s", strings.Join(c.Errors, ","))
}

func (c *CrawlError) CanPresent() bool {
	return len(c.Errors) != 0
}
