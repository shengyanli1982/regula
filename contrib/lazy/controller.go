package lazy

import (
	"github.com/shengyanli1982/karta"
	"github.com/shengyanli1982/regula"
	fctrl "github.com/shengyanli1982/regula"
	rl "github.com/shengyanli1982/regula/ratelimiter"
	"github.com/shengyanli1982/workqueue"
)

// NewSimpleFlowController 是一个函数，它创建并返回一个简单版的流控制器
// NewSimpleFlowController is a function that creates and returns a simple of the flow controller
func NewSimpleFlowController(rate float64, burst int64, cb fctrl.Callback) *fctrl.FlowController {
	// 创建一个新的配置
	// Create a new configuration
	kconf := karta.NewConfig()

	// 创建一个新的假延迟队列
	// Create a new fake delay queue
	queue := workqueue.NewDelayingQueueWithCustomQueue(nil, workqueue.NewSimpleQueue(nil))

	// 使用队列和配置创建一个新的管道
	// Create a new pipeline using the queue and configuration
	pl := karta.NewPipeline(queue, kconf)

	// 创建一个新的速率限制器，并设置速率为 rate，突发为 burst
	// Create a new rate limiter and set the rate to `rate` and burst to `burst`
	rl := rl.NewRateLimiter(rl.NewConfig().WithRate(rate).WithBurst(burst))

	// 创建一个新的流控制器配置，并设置回调函数和速率限制器
	// Create a new flow controller configuration and set the callback function and rate limiter
	fconf := regula.NewConfig().WithRateLimiter(rl).WithCallback(cb)

	// 使用管道和配置创建一个新的流控制器
	// Create a new flow controller using the pipeline and configuration
	fc := regula.NewFlowController(pl, fconf)

	// 返回流控制器
	// Return the flow controller
	return fc
}
