package proxy

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"log"
	"strings"
	"time"
)

var cache = map[string]Cache{}

type Cache struct {
	body []byte
	ttl  time.Time
}

type CacheProxy struct {
	key string
	ttl time.Duration
}

func NewCacheProxy(key string, ttl time.Duration) *CacheProxy {
	return &CacheProxy{key: key, ttl: ttl}
}

func (p CacheProxy) Accept(key string) bool {
	return p.key == key
}

func (p CacheProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()
	key := c.Params("key")

	if v, ok := cache[path]; ok && v.ttl.After(time.Now()) {
		c.Response().SetBody(v.body)
		c.Response().Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		c.Response().Header.Add("cache-control", fmt.Sprintf("max-age:%d", p.ttl/time.Second))
		return nil
	}

	// https://mocki.io/v1/2bd1f2fe-5580-4d9f-84c7-de59f4a5bbfe
	url := "https://mocki.io" + strings.TrimPrefix(path, "/"+key)
	log.Printf("http request redirecting to %s", url)

	if err := proxy.Do(c, url); err != nil {
		return err
	}

	if c.Response().StatusCode() != fiber.StatusOK {
		return fiber.NewError(c.Response().StatusCode(), "Check your request")
	}

	mutex.Lock()
	cache[path] = Cache{c.Response().Body(), time.Now().Add(p.ttl)}
	mutex.Unlock()

	c.Response().Header.Del(fiber.HeaderServer)

	return nil
}
