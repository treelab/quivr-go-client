package quivr_go_client

import "time"

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}
