package regula

import "time"

// emptyCallback 是一个空结构体，用于实现回调接口
// emptyCallback is an empty structure used to implement the callback interface
type emptyCallback struct{}

// OnLimited 是一个方法，当被限制时，它不执行任何操作
// OnLimited is a method that does nothing when being limited
func (emptyCallback) OnLimited(delay time.Duration) {}

// NewEmptyCallback 是一个函数，它创建并返回一个新的emptyCallback
// NewEmptyCallback is a function that creates and returns a new emptyCallback
func NewEmptyCallback() Callback {
	return &emptyCallback{}
}
