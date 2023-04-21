package wrapping

import (
	"log"
	"net/http"
	"time"
)

func Logger(logger *log.Logger) Decorator {
	return func(c http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			start := time.Now()
			logger.Printf("started request to %s at %s\n",
				r.URL,
				start.Format("2006-01-02 15:04:05"))

			res, err := c.RoundTrip(r)
			logger.Printf("completed request to %s in %s\n",
				r.URL,
				time.Since(start))

			return res, err
		})
	}
}

func BasicAuth(username, password string) Decorator {
	return func(c http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			r.SetBasicAuth(username, password)
			res, err := c.RoundTrip(r)

			return res, err
		})
	}
}
