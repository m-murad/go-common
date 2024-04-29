package http

import (
	"net/http"

	"github.com/m-murad/go-common/metrics"
)

type Client struct {
	cli         *http.Client
	metrics     metrics.Service
	serviceName string

	// client level tags
	tags []string
}

func NewClient(httpCli *http.Client, metrics metrics.Service, serviceName string) *Client {
	cli := Client{
		cli:         httpCli,
		metrics:     metrics,
		serviceName: serviceName,
		tags:        []string{"service:" + serviceName},
	}

	return &cli
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	rsp, err := c.cli.Do(req)
	c.metrics.Incr("http_request", 1, append([]string{"status:" + rsp.Status, "method:" + req.Method, "path:" + req.URL.Path}, c.tags...))
	c.metrics.Timing("http_request", 1, append([]string{"status:" + rsp.Status, "method:" + req.Method, "path:" + req.URL.Path}, c.tags...))
	return rsp, err
}
