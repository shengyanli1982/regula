package ratelimiter

import (
	"math"

	"golang.org/x/time/rate"
)

// DefaultInfniteRate 是无限速率的默认值
// DefaultInfniteRate is the default value for infinite rate
const DefaultInfniteRate = rate.Limit(math.MaxFloat64)

// DefaultLimitRate 是限制速率的默认值
// DefaultLimitRate is the default value for limit rate
const DefaultLimitRate = 10.0

// DefaultLimitBurst 是限制突发的默认值
// DefaultLimitBurst is the default value for limit burst
const DefaultLimitBurst = 5

// Config 是配置的结构体
// Config is the structure for configuration
type Config struct {
	rate  float64
	burst int64
}

// NewConfig 是创建新配置的函数
// NewConfig is a function to create a new configuration
func NewConfig() *Config {
	return &Config{
		rate:  DefaultLimitRate,
		burst: DefaultLimitBurst,
	}
}

// DefaultConfig 是获取默认配置的函数
// DefaultConfig is a function to get the default configuration
func DefaultConfig() *Config {
	return NewConfig()
}

// WithRate 它设置配置的速率
// WithRate is a method that sets the rate of the configuration
func (c *Config) WithRate(rate float64) *Config {
	c.rate = rate
	return c
}

// WithBurst 它设置配置的突发
// WithBurst is a method that sets the burst of the configuration
func (c *Config) WithBurst(burst int64) *Config {
	c.burst = burst
	return c
}

// isConfigValid 它检查配置是否有效，如果无效，它将设置为默认值
// isConfigValid is a function that checks if the configuration is valid, if not, it sets it to the default values
func isConfigValid(conf *Config) *Config {
	if conf != nil {
		if conf.rate <= 0 {
			conf.rate = DefaultLimitRate
		}
		if conf.burst <= 0 {
			conf.burst = DefaultLimitBurst
		}
	} else {
		conf = DefaultConfig()
	}

	return conf
}
