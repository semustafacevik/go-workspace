package proxy

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex
var counter = map[string]*Limit{}

type Limit struct {
	count int
	ttl   time.Time
}

type LimitProxy struct {
	key   string
	limit int
	ttl   time.Duration
}

func NewLimitProxy(key string, limit int, ttl time.Duration) *LimitProxy {
	return &LimitProxy{key: key, limit: limit, ttl: ttl}
}

func (p LimitProxy) Accept(key string) bool {
	return p.key == key
}

func (p LimitProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()

	if v, ok := counter[path]; ok && v.count >= p.limit && v.ttl.After(time.Now()) {
		c.Response().SetStatusCode(fiber.StatusTooManyRequests)
		log.Printf("request limit reached for \"%s\"", path)
		return fiber.ErrTooManyRequests
	} else if (ok && v.ttl.Before(time.Now())) || !ok {
		mutex.Lock()
		counter[path] = &Limit{count: 0, ttl: time.Now().Add(p.ttl)}
		mutex.Unlock()
	}

	c.SendString("go-workspace | proxy -> " + path)

	mutex.Lock()
	counter[path].count++
	mutex.Unlock()

	return nil
}
