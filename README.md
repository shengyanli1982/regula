English | [中文](./README_CN.md)

<div align="center">
	<img src="assets/logo.png" alt="logo" width="500px">
	</br></br></br>
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/shengyanli1982/regula)](https://goreportcard.com/report/github.com/shengyanli1982/regula)
[![Build Status](https://github.com/shengyanli1982/regula/actions/workflows/test.yaml/badge.svg)](https://github.com/shengyanli1982/regula/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/shengyanli1982/regula.svg)](https://pkg.go.dev/github.com/shengyanli1982/regula)

# What is Regula?

`Regula` is a flow control component designed to help Golang applications manage concurrency and data flow. It is designed to be simple, efficient, and easy to use, and it is designed to be used in a wide range of applications, from simple web applications to complex distributed systems.

`Regula` is based on the `workqueue` + `workpool` + `ratelimiter` pattern. This allows you to submit any function to `Regula` for concurrent and rate-limited execution. Therefore, `Regula` is suitable for scenarios that require both concurrency and rate limiting.

`Regula` is a perfect fit for limiting the rate of requests to any resource, such as a database, API, or file.

# Why need Regula?

In Golang, we have various ways to handle concurrency and data flow, such as `channel`, `sync`, `context`, etc. However, these approaches can be challenging to use and may not provide optimal efficiency. `Regula` simplifies the complexity of concurrency and data flow, allowing developers to focus on the business logic.

`Regula` offers the flexibility for core components to implement interfaces and replace internal modules based on specific needs. This design may make `Regula` appear complex, but it actually simplifies its usage significantly.

If `Regula` had been available earlier, I believe I could have finished my work more efficiently and possibly even left early every day.

## Advantages

`Regula` is designed with the following key advantages:

-   **Simple**: `Regula` is designed to be simple and easy to use. It provides a simple method that allows developers to submit functions and manage concurrency and data flow with ease.
-   **Efficient**: `Regula` is designed to be lightweight and to provide optimal performance in a wide range of applications.
-   **Flexible**: `Regula` is designed to be flexible. It doesn't bind you to specifics of the `workqueue` + `workpool` + `ratelimiter` pattern.
-   **Scalable**: `Regula` is designed to be scalable, you can custom the `pipeline` and `ratelimiter` to fit your needs.
-   **Reliable**: `Regula` is designed to be reliable. It based on proven technologies and has been tested in a wide range of applications.

# Installation

### 1. Expert mode

`Expert` mode allows you to use the `Regula` library in a more flexible way. You can create a custom flow controller by utilizing the `pipeline` and `ratelimiter` interfaces, giving you control over the details of the modules.

```bash
go get github.com/shengyanli1982/regula
```

### 2. Lazy mode

In `Lazy` mode, you can use the `NewSimpleFlowController` method to create a new flow controller with the default `pipeline` and `ratelimiter` modules, making it more convenient to use the `Regula` library.

```bash
go get github.com/shengyanli1982/regula/contrib/lazy
```

# Quick Start

`Regula` is very simple to use. Just few lines of code to get started.

## 1. Config

`Regula` has a config object, which can be used to register the `pipeline` and `ratelimiter` modules. The config object has the following fields:

-   `WithRateLimiter`: Register the `ratelimiter` module.
-   `WithCallback`: Set the callback function for `Regula` submit function.

> [!TIP]
> If you want to use a custom `pipeline` or `ratelimiter` module, you can implement the specific internal interface and pass it to the config object.
>
> `pipline` module only work in `NewFlowController` method.

**Pipeline Interface**

```go
// Pipeline 是一个管道接口，用于添加事件到管道、延迟添加事件到管道以及停止管道的操作。
// Pipeline is a pipeline interface for adding events to the pipeline, delaying events to the pipeline, and stopping the pipeline.
type Pipeline = interface {
	// SubmitWithFunc 将一个新的事件添加到管道中，并指定消息处理函数。
	// SubmitWithFunc adds a new event to the pipeline and specifies the message processing function.
	SubmitWithFunc(fn MessageHandleFunc, msg any) error

	// SubmitAfterWithFunc 将一个新的事件添加到管道中，并指定消息处理函数和延迟时间。
	// SubmitAfterWithFunc adds a new event to the pipeline and specifies the message processing function and delay time.
	SubmitAfterWithFunc(fn MessageHandleFunc, msg any, delay time.Duration) error

	// Stop 停止管道的运行。
	// Stop stops the pipeline.
	Stop()
}
```

**Ratelimiter Interface**

```go
// RateLimiter 是一个接口，定义了一个方法，该方法返回下一个事件的延迟时间
// RateLimiter is an interface that defines a method that returns the delay time of the next event
type RateLimiter = interface {
	// When 返回下一个事件的延迟时间
	// When returns the delay time of the next event
	When() time.Duration
}
```

## 2. Components

The `Regula` library has the following components:

### 2.1. Ratelimiter

`Ratelimiter` is a rate limiter module. It use google `golang.org/x/time/rate` package which mean it use bucket algorithm to control the rate of events.

#### 2.1.1. Config

-   `WithRate`: Set the rate of events per second. Default is `DefaultLimitRate`.
-   `WithBurst`: Set the burst of events. Default is `DefaultLimitBurst`.

#### 2.1.2. Methods

-   `When`: Return the delay time of the next event.

## 3. Methods

The `Regula` provides the following methods:

-   `NewFlowController`: Create a new flow controller.
-   `Stop`: Stop the flow controller.
-   `Do`: Submit a function to the flow controller.

> [!NOTE]
> If you use `lazy` mode, you can use the `NewSimpleFlowController` method to create a new flow controller. The flow controller will use the default `pipeline` and `ratelimiter` modules. The `NewSimpleFlowController` method provides the `callback` function, `rate`, and `burst` parameters.

## 4. Callback

-   `OnExecLimited`: This method is called when the event handling is limited.

## 5. Examples

Example code is located in the `examples` directory.

### 5.1. Expert mode

In `Expert` mode, you can create a custom flow controller using the `NewFlowController` method. This allows you to customize the `pipeline` and `ratelimiter` modules.

With `Expert` mode, you have full control over the implementation of the `pipeline` and `ratelimiter` modules, giving you the flexibility to customize them according to your needs.

Here is an example of how to use the `Regula` library in `Expert` mode:

1. Create a new pipeline with a custom delay queue.
2. Create a new flow controller with a custom pipeline and rate limiter.
3. Set a callback function for the flow controller.
4. Start 10 goroutines.
5. Submit a function to the flow controller.
6. Wait for all goroutines to finish.

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/shengyanli1982/karta"
	"github.com/shengyanli1982/regula"
	rl "github.com/shengyanli1982/regula/ratelimiter"
	wkq "github.com/shengyanli1982/workqueue/v2"
)

// demoCallback 是一个空结构体，用于实现回调接口
// demoCallback is an empty structure used to implement the callback interface
type demoCallback struct{}

// OnExecLimited 是一个方法，当被限制时，它会打印出被限制的延迟时间
// OnExecLimited is a method that prints the limited delay time when being limited
func (demoCallback) OnExecLimited(msg any, delay time.Duration) {
	fmt.Printf("limited -> msg: %v, delay: %v\n", msg, delay.String())
}

// newCallback 是一个函数，它创建并返回一个新的demoCallback
// newCallback is a function that creates and returns a new demoCallback
func newCallback() *demoCallback {
	return &demoCallback{}
}

func main() {
	// 创建一个新的配置，并设置工作线程数为2
	// Create a new configuration and set the number of worker threads to 2
	kconf := karta.NewConfig().WithWorkerNumber(2)

	// 创建一个新的假延迟队列
	// Create a new fake delay queue
	queue := karta.NewFakeDelayingQueue(wkq.NewQueue(nil))

	// 使用队列和配置创建一个新的管道
	// Create a new pipeline using the queue and configuration
	pl := karta.NewPipeline(queue, kconf)

	// 创建一个新的速率限制器，并设置速率为10，突发为1
	// Create a new rate limiter and set the rate to 10 and burst to 1
	rl := rl.NewRateLimiter(rl.NewConfig().WithRate(10).WithBurst(1))

	// 创建一个新的流控制器配置，并设置回调函数和速率限制器
	// Create a new flow controller configuration and set the callback function and rate limiter
	fconf := regula.NewConfig().WithCallback(newCallback()).WithRateLimiter(rl)

	// 使用管道和配置创建一个新的流控制器
	// Create a new flow controller using the pipeline and configuration
	fc := regula.NewFlowController(pl, fconf)

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

	// 等待2秒
	// Wait for 2 seconds
	time.Sleep(time.Second * 2)
}
```

**Result**

```bash
$ go run demo.go
msg: test -> 9
limited -> msg: test, delay: 400ms
limited -> msg: test, delay: 700ms
limited -> msg: test, delay: 300ms
limited -> msg: test, delay: 200ms
limited -> msg: test, delay: 500ms
limited -> msg: test, delay: 600ms
limited -> msg: test, delay: 100ms
limited -> msg: test, delay: 800ms
limited -> msg: test, delay: 900ms
msg: test -> 5
msg: test -> 3
msg: test -> 0
msg: test -> 7
msg: test -> 2
msg: test -> 4
msg: test -> 6
msg: test -> 1
msg: test -> 8
```

### 5.2. Lazy mode

In `lazy` mode, you can use the `NewSimpleFlowController` method to create a new flow controller. This method wraps the `NewFlowController` method and uses the default `pipeline` and `ratelimiter` modules.

The `NewSimpleFlowController` method allows you to specify the `rate` and `burst` parameters, and provides an optional `callback` function.

Here's an example of using the `lazy` mode:

1. Create a new flow controller with a rate of 2, a burst of 1, and an optional callback function.
2. Start 10 goroutines.
3. Submit a function to the flow controller.
4. Wait for all goroutines to finish.

```go
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
func (demoCallback) OnExecLimited(msg any, delay time.Duration) {
	fmt.Printf("limited -> msg: %v, delay: %v\n", msg, delay.String())
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
```

**Result**

```bash
$ go run demo.go
limited -> msg: test, delay: 1s
limited -> msg: test, delay: 500ms
limited -> msg: test, delay: 3.5s
limited -> msg: test, delay: 1.5s
limited -> msg: test, delay: 3s
limited -> msg: test, delay: 2s
limited -> msg: test, delay: 2.5s
limited -> msg: test, delay: 4.5s
limited -> msg: test, delay: 4s
msg: test -> 9
msg: test -> 4
msg: test -> 5
msg: test -> 0
msg: test -> 8
msg: test -> 6
msg: test -> 3
msg: test -> 2
msg: test -> 7
msg: test -> 1
```
