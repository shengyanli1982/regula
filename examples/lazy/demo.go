package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/shengyanli1982/regula/contrib/lazy"
)

// demoCallback 是一个空结构体，用于实现回调接口
// demoCallback is an empty structure used to implement the callback interface
type demoCallback struct{}

// OnExecLimited 是一个方法，当被限制时，它会打印出被限制的延迟时间
// OnExecLimited is a method that prints the limited delay time when being limited
func (demoCallback) OnExecLimited(delay time.Duration) {
	fmt.Printf("limited: %v\n", delay.String())
}

// newCallback 是一个函数，它创建并返回一个新的demoCallback
// newCallback is a function that creates and returns a new demoCallback
func newCallback() *demoCallback {
	return &demoCallback{}
}

func main() {
	// 创建一个新的流控制器，设置速率为2，突发为1和回调函数
	// Create a new flow controller, set the rate to 2, burst to 1 and callback function
	fc := lazy.NewSimpleFlowController(float64(2), 1, newCallback())

	// 在函数结束时停止管道、队列和流控制器 (FlowController 会关闭 Pipeline，Pipeline 会关闭 Queue)
	// Stop the pipeline, queue and flow controller when the function ends (FlowController will close Pipeline, Pipeline will close Queue)
	defer fc.Stop()

	// 创建一个等待组
	// Create a wait group
	wg := sync.WaitGroup{}

	// 启动10个协程
	// Start 10 goroutines
	for i := 0; i < 10; i++ {
		v := i
		// 增加等待组的计数
		// Increase the count of the wait group
		wg.Add(1)
		go func() {
			// 在协程结束时减少等待组的计数
			// Decrease the count of the wait group when the goroutine ends
			defer wg.Done()
			// 执行一个消息处理函数，如果有错误，打印错误信息
			// Execute a message handle function, if there is an error, print the error message
			if err := fc.Do(func(msg any) (any, error) {
				fmt.Printf("msg: %v -> %v\n", msg, v)
				return msg, nil
			}, "test"); err != nil {
				fmt.Printf("fc.Do should not return error: %v\n", err)
			}
		}()
	}

	// 等待所有协程结束
	// Wait for all goroutines to end
	wg.Wait()

	// 等待5秒
	// Wait for 5 seconds
	time.Sleep(time.Second * 5)
}
