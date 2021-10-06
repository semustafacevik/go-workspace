package proxy

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Proxy interface {
	Accept(key string) bool
	Proxy(c *fiber.Ctx) error
}

var proxies = []Proxy{
	NewLimitProxy("user", 5, 3*time.Minute),
	NewLimitProxy("channel", 10, 10*time.Minute),
}

func ProxyHandler(c *fiber.Ctx) error {
	for _, v := range proxies {
		if v.Accept(c.Params("key")) {
			return v.Proxy(c)
		}
	}
	c.Response().SetStatusCode(fiber.StatusNotFound)
	return fiber.ErrNotFound
}
