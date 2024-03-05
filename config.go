package regula

import (
	rl "github.com/shengyanli1982/regula/ratelimiter"
)

// Config 是配置的结构体，包含一个速率限制器接口
// Config is the structure for configuration, containing a rate limiter interface
type Config struct {
	ratelimiter RateLimiterInterface
	callback    Callback
}

// NewConfig 是创建新配置的函数，它返回一个包含默认无操作限制器的配置
// NewConfig is a function to create a new configuration, it returns a configuration with a default no-operation limiter
func NewConfig() *Config {
	return &Config{
		ratelimiter: rl.NewNopLimiter(),
		callback:    NewEmptyCallback(),
	}
}

// DefaultConfig 是获取默认配置的函数，它返回一个新的配置
// DefaultConfig is a function to get the default configuration, it returns a new configuration
func DefaultConfig() *Config {
	return NewConfig()
}

// WithRateLimiter 它设置配置的速率限制器
// WithRateLimiter is a method that sets the rate limiter of the configuration
func (c *Config) WithRateLimiter(rl RateLimiterInterface) *Config {
	c.ratelimiter = rl
	return c
}

// WithCallback 它设置配置的回调函数
// WithCallback is a method that sets the callback function of the configuration
func (c *Config) WithCallback(cb Callback) *Config {
	c.callback = cb
	return c
}

// isConfigValid 是一个函数，它检查配置是否有效，如果无效，它将设置为默认值
// isConfigValid is a function that checks if the configuration is valid, if not, it sets it to the default values
func isConfigValid(conf *Config) *Config {
	if conf != nil {
		if conf.ratelimiter == nil {
			conf.ratelimiter = rl.NewNopLimiter()
		}
		if conf.callback == nil {
			conf.callback = NewEmptyCallback()
		}
	} else {
		conf = DefaultConfig()
	}

	return conf
}
