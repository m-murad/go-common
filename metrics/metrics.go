package metrics

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

type client struct {
	cli *statsd.Client

	// service level tags
	tags []string
}

func NewClient(url string, appName string, env string) (*client, error) {
	statsd, err := statsd.New(url)
	if err != nil {
		return nil, err
	}

	tags := []string{fmt.Sprintf("app:%s", appName), fmt.Sprintf("env:%s", env)}

	cli := client{
		cli:  statsd,
		tags: tags,
	}

	return &cli, nil
}

func (c *client) Incr(name string, val float64, tags []string) {
	c.cli.Incr(name, append(tags, c.tags...), val)
}

func (c *client) Decr(name string, val float64, tags []string) {
	c.cli.Decr(name, append(tags, c.tags...), val)
}

func (c *client) Count(name string, val float64, tags []string) {
	c.cli.Count(name, 2, append(tags, c.tags...), val)
}

func (c *client) Timing(name string, dur time.Duration, tags []string) {
	c.cli.Timing(name, dur, append(tags, c.tags...), 1)
}
