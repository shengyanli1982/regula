[English](./README.md) | 中文

<div align="center">
	<img src="assets/logo.png" alt="logo" width="500px">
</div>

## 什么是 Regula？

`Regula` 是一个流控组件，旨在帮助 Golang 应用程序管理并发和数据流。它被设计为简单、高效和易于使用，并可用于各种应用程序，从简单的 Web 应用程序到复杂的分布式系统。

`Regula` 基于 `workqueue` + `workpool` + `ratelimiter` 模式。因此，您可以将任何函数提交给 `Regula`，它将以并发和限速的方式执行。

## 为什么需要 Regula？

在 Golang 中，我们有多种处理并发和数据流的方式，例如 `channel`、`sync`、`context` 等。然而，这些方法使用起来可能具有挑战性，并且可能无法提供最佳的效率。`Regula` 简化了并发和数据流的复杂性，使开发人员能够专注于业务逻辑。

如果早些时候有了 `Regula`，我相信我可以更高效地完成工作，甚至可能每天提前下班。

### 优势

`Regula` 具有以下关键优势：

-   **简单**：`Regula` 设计简单易用。它提供了一种简单的方法，让开发人员轻松提交函数并管理并发和数据流。
-   **高效**：`Regula` 设计轻量且在各种应用程序中提供最佳性能。
-   **灵活**：`Regula` 设计灵活。它不限制您使用 `workqueue` + `workpool` + `ratelimiter` 模式。
-   **可扩展**：`Regula` 设计可扩展，您可以自定义 `pipeline` 和 `ratelimiter` 来满足您的需求。
-   **可靠**：`Regula` 设计可靠。它基于经过验证的技术，并在各种应用程序中进行了测试。

## 安装

### 1. 标准模式

```bash
go get github.com/shengyanli1982/regula
```

### 2. 懒惰模式

```bash
go get github.com/shengyanli1982/regula/contrib/lazy
```

# 快速入门

`Regula` 的使用非常简单，只需几行代码即可开始使用。

## 1. 配置

`Regula` 有一个配置对象，用于注册 `pipeline` 和 `ratelimiter` 模块。配置对象具有以下字段：

-   `WithRateLimiter`：注册 `ratelimiter` 模块。
-   `WithCallback`：为 `Regula` 提交函数设置回调函数。

> [!TIP]
> 如果您想使用自定义的 `pipeline` 或 `ratelimiter` 模块，可以实现特定的内部接口并将其传递给配置对象。
>
> `pipeline` 模块仅在 `NewFlowController` 方法中起作用。

**Pipeline 接口**

```go
// PipelineInterface 是一个管道接口，用于添加事件到管道、延迟添加事件到管道以及停止管道的操作。
// PipelineInterface is a pipeline interface for adding events to the pipeline, delaying events to the pipeline, and stopping the pipeline.
type PipelineInterface = interface {
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

**速率限制器接口**

```go
// RateLimiterInterface 是一个接口，定义了一个方法，该方法返回下一个事件的延迟时间
// RateLimiterInterface is an interface that defines a method that returns the delay time of the next event
type RateLimiterInterface = interface {
	// When 返回下一个事件的延迟时间
	// When returns the delay time of the next event
	When() time.Duration
}
```

## 2. 组件

`Regula` 库具有以下组件：

### 2.1. 速率限制器

`Ratelimiter` 是一个速率限制器模块。它使用 Google 的 `golang.org/x/time/rate` 包，采用桶算法来控制事件的速率。

#### 2.1.1. 配置

-   `WithRate`：设置每秒的事件速率。默认值为 `DefaultLimitRate`。
-   `WithBurst`：设置事件的突发数量。默认值为 `DefaultLimitBurst`。

#### 2.1.2. 方法

-   `When`：返回下一个事件的延迟时间。

## 3. 方法

`Regula` 提供以下方法：

-   `NewFlowController`：创建一个新的流控制器。
-   `Stop`：停止流控制器。
-   `Do`：将函数提交给流控制器。

> [!NOTE]
> 如果您使用 `懒惰模式`，可以使用 `NewSimpleFlowController` 方法创建一个新的流控制器。流控制器将使用默认的 `pipeline` 和 `ratelimiter` 模块。`NewSimpleFlowController` 方法提供了 `回调函数`、`速率` 和 `突发数量` 参数。

## 4. 示例

示例代码位于 `examples` 目录中。

### 4.1. 标准模式

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/shengyanli1982/karta"
	"github.com/shengyanli1982/regula"
	rl "github.com/shengyanli1982/regula/ratelimiter"
	"github.com/shengyanli1982/workqueue"
)

// demoCallback 是一个空结构体，用于实现回调接口
// demoCallback is an empty structure used to implement the callback interface
type demoCallback struct{}

// OnLimited 是一个方法，当被限制时，它会打印出被限制的延迟时间
// OnLimited is a method that prints the limited delay time when being limited
func (demoCallback) OnLimited(delay time.Duration) {
	fmt.Printf("limited: %v\n", delay.String())
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
	queue := workqueue.NewDelayingQueueWithCustomQueue(nil, workqueue.NewSimpleQueue(nil))

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

**执行结果**

```bash
$ go run demo.go
limited: 400ms
limited: 800ms
limited: 100ms
limited: 300ms
limited: 900ms
limited: 500ms
limited: 700ms
msg: test -> 9
limited: 200ms
limited: 600ms
msg: test -> 4
msg: test -> 0
msg: test -> 1
msg: test -> 5
msg: test -> 6
msg: test -> 2
msg: test -> 8
msg: test -> 3
msg: test -> 7
```

### 4.2. 懒惰模式

在懒惰模式下，您可以使用 `NewSimpleFlowController` 方法创建一个新的流控制器。该方法封装了 `NewFlowController` 方法，并使用默认的 `pipeline` 和 `ratelimiter` 模块。

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

// OnLimited 是一个方法，当被限制时，它会打印出被限制的延迟时间
// OnLimited is a method that prints the limited delay time when being limited
func (demoCallback) OnLimited(delay time.Duration) {
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
```

**执行结果**

```bash
$ go run demo.go
limited: 4.5s
limited: 1.5s
msg: test -> 4
limited: 500ms
limited: 2s
limited: 2.5s
limited: 3s
limited: 1s
limited: 3.5s
limited: 4s
msg: test -> 9
msg: test -> 1
msg: test -> 0
msg: test -> 5
msg: test -> 2
msg: test -> 8
msg: test -> 3
msg: test -> 6
msg: test -> 7
```
