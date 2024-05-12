package ratelimiter

import (
	"math"

	"golang.org/x/time/rate"
)

// DefaultInfniteRate 是默认的无限速率，它的值是 math.MaxFloat64
// DefaultInfniteRate is the default infinite rate, its value is math.MaxFloat64
const DefaultInfniteRate = rate.Limit(math.MaxFloat64)

// DefaultLimitRate 是默认的限制速率，它的值是 10.0
// DefaultLimitRate is the default limit rate, its value is 10.0
const DefaultLimitRate = 10.0

// DefaultLimitBurst 是默认的限制突发，它的值是 5
// DefaultLimitBurst is the default limit burst, its value is 5
const DefaultLimitBurst = 5

// Config 是配置的结构体，包含了速率和突发值
// Config is the structure for configuration, it includes rate and burst
type Config struct {
	// rate 是限制的速率
	// rate is the limit rate
	rate float64

	// burst 是限制的突发值
	// burst is the limit burst
	burst int64
}

// NewConfig 是创建新配置的函数，它返回一个包含默认限制速率和突发值的配置
// NewConfig is a function to create a new configuration, it returns a configuration with default limit rate and burst
func NewConfig() *Config {
	return &Config{
		// rate 是默认的限制速率
		// rate is the default limit rate
		rate: DefaultLimitRate,

		// burst 是默认的限制突发值
		// burst is the default limit burst
		burst: DefaultLimitBurst,
	}
}

// DefaultConfig 是获取默认配置的函数
// DefaultConfig is a function to get the default configuration
func DefaultConfig() *Config {
	return NewConfig()
}

// WithRate 是一个方法，它设置配置的速率
// WithRate is a method that sets the rate of the configuration
func (c *Config) WithRate(rate float64) *Config {
	// 设置配置的速率
	// Set the rate of the configuration
	c.rate = rate

	// 返回配置
	// Return the configuration
	return c
}

// WithBurst 是一个方法，它设置配置的突发
// WithBurst is a method that sets the burst of the configuration
func (c *Config) WithBurst(burst int64) *Config {
	// 设置配置的突发
	// Set the burst of the configuration
	c.burst = burst

	// 返回配置
	// Return the configuration
	return c
}

// isConfigValid 是一个函数，它检查配置是否有效，如果无效，它将设置为默认值
// isConfigValid is a function that checks if the configuration is valid, if not, it sets it to the default values
func isConfigValid(conf *Config) *Config {
	// 如果配置不为空
	// If the configuration is not null
	if conf != nil {
		// 如果配置的速率小于等于0
		// If the rate of the configuration is less than or equal to 0
		if conf.rate <= 0 {
			// 将配置的速率设置为默认限制速率
			// Set the rate of the configuration to the default limit rate
			conf.rate = DefaultLimitRate
		}

		// 如果配置的突发小于等于0
		// If the burst of the configuration is less than or equal to 0
		if conf.burst <= 0 {
			// 将配置的突发设置为默认限制突发
			// Set the burst of the configuration to the default limit burst
			conf.burst = DefaultLimitBurst
		}
	} else {
		// 如果配置为空，将配置设置为默认配置
		// If the configuration is null, set the configuration to the default configuration
		conf = DefaultConfig()
	}

	// 返回配置
	// Return the configuration
	return conf
}
