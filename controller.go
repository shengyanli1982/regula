package regula

import (
	"sync"

	rl "github.com/shengyanli1982/regula/ratelimiter"
)

// FlowController 是流控制器的结构体，它包含配置、管道接口和一次性同步
// FlowController is the structure of the flow controller, it contains configuration, pipeline interface and once sync
type FlowController struct {
	// config 是流控制器的配置
	// config is the configuration of the flow controller
	config *Config

	// pipline 是流控制器的管道接口
	// pipline is the pipeline interface of the flow controller
	pipline PipelineInterface

	// once 是用于确保某个操作只执行一次的同步原语
	// once is a synchronization primitive used to ensure that an operation is performed only once
	once sync.Once
}

// NewFlowController 是创建新的流控制器的函数，它接受一个管道接口和配置
// NewFlowController is a function to create a new flow controller, it accepts a pipeline interface and configuration
func NewFlowController(pipline PipelineInterface, conf *Config) *FlowController {
	// 如果管道接口为空，则返回 nil
	// If the pipeline interface is nil, return nil
	if pipline == nil {
		return nil
	}

	// 检查配置是否有效，如果无效则使用默认配置
	// Check if the configuration is valid, if not, use the default configuration
	conf = isConfigValid(conf)

	// 返回一个新的流控制器，包含配置、管道接口和一次性同步
	// Return a new flow controller, including configuration, pipeline interface and once sync
	return &FlowController{
		// config 是流控制器的配置
		// config is the configuration of the flow controller
		config: conf,

		// pipline 是流控制器的管道接口
		// pipline is the pipeline interface of the flow controller
		pipline: pipline,

		// once 是用于确保某个操作只执行一次的同步原语
		// once is a synchronization primitive used to ensure that an operation is performed only once
		once: sync.Once{},
	}
}

// Stop 是一个方法，它停止流控制器的管道
// Stop is a method that stops the pipeline of the flow controller
func (fc *FlowController) Stop() {
	// 使用 sync.Once 确保管道只被停止一次
	// Use sync.Once to ensure the pipeline is stopped only once
	fc.once.Do(func() {
		// 停止管道
		// Stop the pipeline
		fc.pipline.Stop()
	})
}

// Do 是一个方法，它执行一个消息处理函数，如果有延迟，它会在延迟后提交函数，否则直接提交
// Do is a method that executes a message handle function, if there is a delay, it submits the function after the delay, otherwise it submits directly
func (fc *FlowController) Do(fn MessageHandleFunc, msg any) error {
	// 通过速率限制器获取下一个事件的延迟时间
	// Get the delay time of the next event through the rate limiter
	delay := fc.config.ratelimiter.When().Round(rl.DefaultEffectiveTimeSliceInterval)

	// 如果有延迟
	// If there is a delay
	if delay > 0 {
		// 调用回调函数，通知有延迟
		// Call the callback function to notify that there is a delay
		fc.config.callback.OnExecLimited(delay)

		// 在延迟后提交函数
		// Submit the function after the delay
		return fc.pipline.SubmitAfterWithFunc(fn, msg, delay)
	}

	// 如果没有延返，直接提交函数
	// If there is no delay, submit the function directly
	return fc.pipline.SubmitWithFunc(fn, msg)
}
