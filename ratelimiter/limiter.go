package ratelimiter

import (
	"time"

	"golang.org/x/time/rate"
)

// DefaultEffectiveTimeSliceInterval 是默认的有效时间片间隔
// DefaultEffectiveTimeSliceInterval is the default effective time slice interval
const DefaultEffectiveTimeSliceInterval = time.Millisecond * 100

// Limiter 是一个限流器结构体，包含了一个 rate.Limiter
// Limiter is a structure for rate limiter, it includes a rate.Limiter
type Limiter struct {
	// limiter 是 rate.Limiter 的实例
	// limiter is an instance of rate.Limiter
	limiter *rate.Limiter
}

// NewRateLimiter 是创建新的限流器的函数，它接受一个配置参数
// NewRateLimiter is a function to create a new rate limiter, it accepts a configuration parameter
func NewRateLimiter(conf *Config) *Limiter {
	// 检查配置是否有效，如果无效则使用默认配置
	// Check if the configuration is valid, if not, use the default configuration
	conf = isConfigValid(conf)

	// 返回一个新的限流器，包含一个 rate.Limiter
	// Return a new rate limiter, including a rate.Limiter
	return &Limiter{
		// limiter 是一个新的 rate.Limiter，它的速率和突发值由配置决定
		// limiter is a new rate.Limiter, its rate and burst are determined by the configuration
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
