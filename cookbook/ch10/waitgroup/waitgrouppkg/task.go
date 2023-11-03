package waitgrouppkg

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type CrawlError struct {
	Errors []string
}

func GetURL(url string) (*http.Response, error) {
	start := time.Now()
	log.Printf("getting %s", url)
	resp, err := http.Get(url)
	log.Printf("completed getting %s in %s", url, time.Since(start))
	return resp, err
}

func (c *CrawlError) Add(err error) {
	c.Errors = append(c.Errors, err.Error())
}

func (c *CrawlError) Error() string {
	return fmt.Sprintf("All Errors: %s", strings.Join(c.Errors, ","))
}

// 에러가 없으면 false 있으면 true
func (c *CrawlError) Present() bool {
	return len(c.Errors) != 0
}
