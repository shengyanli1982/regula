package regula

import (
	rl "github.com/shengyanli1982/regula/ratelimiter"
)

// Config 是配置的结构体，包含一个速率限制器接口
// Config is the structure for configuration, containing a rate limiter interface
type Config struct {
	ratelimiter RateLimiter
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
func (c *Config) WithRateLimiter(rl RateLimiter) *Config {
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
	// 如果配置不为空
	// If the configuration is not null
	if conf != nil {
		// 如果配置中的速率限制器为空，则设置为无操作限制器
		// If the rate limiter in the configuration is null, set it to a no-operation limiter
		if conf.ratelimiter == nil {
			conf.ratelimiter = rl.NewNopLimiter()
		}
		
		// 如果配置中的回调函数为空，则设置为空回调函数
		// If the callback function in the configuration is null, set it to an empty callback function
		if conf.callback == nil {
			conf.callback = NewEmptyCallback()
		}
	} else {
		// 如果配置为空，则设置为默认配置
		// If the configuration is null, set it to the default configuration
		conf = DefaultConfig()
	}

	// 返回配置
	// Return the configuration
	return conf
}
