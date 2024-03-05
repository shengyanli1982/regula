package ratelimiter

import (
	"time"

	"golang.org/x/time/rate"
)

// DefaultEffectiveTimeSliceInterval 是默认的有效时间片间隔
// DefaultEffectiveTimeSliceInterval is the default effective time slice interval
const DefaultEffectiveTimeSliceInterval = time.Millisecond * 100

// Limiter 是一个限流器结构体
// Limiter is a structure for rate limiter
type Limiter struct {
	limiter *rate.Limiter
}

// NewRateLimiter 是创建新的限流器的函数，它接受一个配置参数
// NewRateLimiter is a function to create a new rate limiter, it accepts a configuration parameter
func NewRateLimiter(conf *Config) *Limiter {
	conf = isConfigValid(conf)
	return &Limiter{
		limiter: rate.NewLimiter(rate.Limit(conf.rate), int(conf.burst)),
	}
}

// When 是一个方法，它返回下一个事件发生的延迟时间
// When is a method that returns the delay for the next event to occur
func (l *Limiter) When() time.Duration {
	return l.limiter.Reserve().Delay()
}

// NopLimiter 是一个不执行任何操作的限流器结构体
// NopLimiter is a structure for a limiter that does not perform any operations
type NopLimiter struct{}

// When 是一个方法，它总是返回0，表示没有延返
// When is a method that always returns 0, indicating no delay
func (l *NopLimiter) When() time.Duration { return 0 }

// NewNopLimiter 是创建新的不执行任何操作的限流器的函数
// NewNopLimiter is a function to create a new limiter that does not perform any operations
func NewNopLimiter() *NopLimiter {
	return &NopLimiter{}
}
