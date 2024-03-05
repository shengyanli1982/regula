package regula

import "time"

// MessageHandleFunc 是一个消息处理函数类型，接收任意类型的消息并返回任意类型的结果和错误。
// MessageHandleFunc is a message processing function type that receives messages of any type and returns results and errors of any type.
type MessageHandleFunc = func(msg any) (any, error)

// PipelineInterface 是一个管道接口，用于添加事件到管道、延迟添加事件到管道以及停止管道的操作。
// PipelineInterface is a pipeline interface for adding events to the pipeline, delaying events to the pipeline, and stopping the pipeline.
type PipelineInterface = interface {
	// SubmitWithFunc 将一个新的事件添加到管道中，并指定消息处理函数。
	// SubmitWithFunc adds a new event to the pipeline and specifies the message processing function.
	SubmitWithFunc(fn MessageHandleFunc, msg any) error

	// SubmitAfterWithFunc 将一个新的事件添加到管道中，并指定消息处理函数和延迟时间。
	// SubmitAfterWithFunc adds a new event to the pipeline and specifies the message processing function and delay time.
	SubmitAfterWithFunc(fn MessageHandleFunc, msg any, delay time.Duration) error

	// GetWorkerNumber 获取管道的工作线程数量。
	// GetWorkerNumber gets the number of working threads in the pipeline.
	GetWorkerNumber() int64

	// Stop 停止管道的运行。
	// Stop stops the pipeline.
	Stop()
}

// RateLimiterInterface 是一个接口，定义了一个方法，该方法返回下一个事件的延迟时间
// RateLimiterInterface is an interface that defines a method that returns the delay time of the next event
type RateLimiterInterface = interface {
	// When 返回下一个事件的延迟时间
	// When returns the delay time of the next event
	When() time.Duration
}

// Callback 是一个接口，定义了一个方法，该方法是达到速率限制时的回调函数
// Callback is an interface that defines a method that is the callback function when the rate limit is reached
type Callback = interface {
	// OnLimited 当达到速率限制时的回调函数
	// OnLimited is the callback function when the rate limit is reached
	OnLimited(delay time.Duration)
}
