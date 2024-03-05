package test

import (
	"testing"
	"time"

	rl "github.com/shengyanli1982/regula/ratelimiter"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter_When(t *testing.T) {
	qps := 2
	interval := time.Second / time.Duration(qps)

	conf := rl.NewConfig().WithRate(float64(qps)).WithBurst(1)
	rl := rl.NewRateLimiter(conf)

	for i := 0; i < 10; i++ {
		assert.Equal(t, rl.When().Round(interval).Milliseconds(), interval.Milliseconds()*int64(i))
	}
}
