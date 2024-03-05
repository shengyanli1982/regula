package regula

import (
	"sync"

	rl "github.com/shengyanli1982/regula/ratelimiter"
)

// FlowController 是流控制器的结构体，包含配置、管道接口和一次性同步
// FlowController is the structure for flow controller, containing configuration, pipeline interface and once sync
type FlowController struct {
	config  *Config
	pipline PipelineInterface
	once    sync.Once
}

// NewFlowController 是创建新的流控制器的函数，它接受一个管道接口和配置
// NewFlowController is a function to create a new flow controller, it accepts a pipeline interface and configuration
func NewFlowController(pipline PipelineInterface, conf *Config) *FlowController {
	if pipline == nil {
		return nil
	}
	conf = isConfigValid(conf)
	return &FlowController{
		config:  conf,
		pipline: pipline,
		once:    sync.Once{},
	}
}

// Stop 是一个方法，它停止流控制器的管道
// Stop is a method that stops the pipeline of the flow controller
func (fc *FlowController) Stop() {
	fc.once.Do(func() {
		fc.pipline.Stop()
	})
}

// Do 是一个方法，它执行一个消息处理函数，如果有延迟，它会在延迟后提交函数，否则直接提交
// Do is a method that executes a message handle function, if there is a delay, it submits the function after the delay, otherwise it submits directly
func (fc *FlowController) Do(fn MessageHandleFunc, msg any) error {
	// 通过速率限制器获取下一个事件的延迟时间
	// Get the delay time of the next event through the rate limiter
	delay := fc.config.ratelimiter.When().Round(rl.DefaultEffectiveTimeSliceInterval)

	// 如果延迟大于0，就在延迟后提交函数，否则直接提交
	// If the delay is greater than 0, submit the function after the delay, otherwise submit directly
	if delay > 0 {
		fc.config.callback.OnLimited(delay)
		return fc.pipline.SubmitAfterWithFunc(fn, msg, delay)
	}
	return fc.pipline.SubmitWithFunc(fn, msg)
}
