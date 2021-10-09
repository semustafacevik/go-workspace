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

	if v, ok := counter[path]; ok && v.count >= p.limit {
		if v.ttl.After(time.Now()) {
			c.Response().SetStatusCode(fiber.StatusTooManyRequests)
			log.Printf("request limit reached for \"%s\"", path)
			return fiber.ErrTooManyRequests
		} else {
			log.Printf("resetting counter values of \"%s\"", path)
			defineCounter(path, p.ttl)
		}
	} else if !ok {
		defineCounter(path, p.ttl)
	}

	if err := c.SendString("go-workspace | proxy -> " + path); err != nil {
		return err
	}

	incrementCounter(path)

	return nil
}

func defineCounter(path string, ttl time.Duration) {
	mutex.Lock()
	counter[path] = &Limit{count: 0, ttl: time.Now().Add(ttl)}
	mutex.Unlock()
}

func incrementCounter(path string) {
	mutex.Lock()
	counter[path].count++
	mutex.Unlock()
}
