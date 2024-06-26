package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/shengyanli1982/karta"
	"github.com/shengyanli1982/regula"
	rl "github.com/shengyanli1982/regula/ratelimiter"
	wkq "github.com/shengyanli1982/workqueue/v2"
	"github.com/stretchr/testify/assert"
)

type testCallback struct{}

func (testCallback) OnExecLimited(msg any, delay time.Duration) {
	fmt.Printf("limited -> msg: %v, delay: %v\n", msg, delay.String())
}

func newTestCallback() *testCallback {
	return &testCallback{}
}

func TestFlowController_Do(t *testing.T) {
	kconf := karta.NewConfig().WithWorkerNumber(2)
	queue := karta.NewFakeDelayingQueue(wkq.NewQueue(nil))
	pl := karta.NewPipeline(queue, kconf)
	fc := regula.NewFlowController(pl, nil)

	defer fc.Stop()

	err := fc.Do(func(msg any) (any, error) {
		fmt.Printf("msg: %v\n", msg)
		return msg, nil
	}, "test")
	assert.NoError(t, err, "fc.Do should not return error")

	time.Sleep(time.Second)
}

func TestFlowController_DoAfter(t *testing.T) {
	kconf := karta.NewConfig().WithWorkerNumber(2)
	queue := karta.NewFakeDelayingQueue(wkq.NewQueue(nil))
	pl := karta.NewPipeline(queue, kconf)
	rl := rl.NewRateLimiter(rl.NewConfig().WithRate(10).WithBurst(1))
	fconf := regula.NewConfig().WithCallback(newTestCallback()).WithRateLimiter(rl)
	fc := regula.NewFlowController(pl, fconf)

	defer fc.Stop()

	for i := 0; i < 10; i++ {
		v := i
		err := fc.Do(func(msg any) (any, error) {
			fmt.Printf("msg: %v -> %v\n", msg, v)
			return msg, nil
		}, "test")
		assert.NoError(t, err, "fc.Do should not return error")
	}

	time.Sleep(time.Second * 2)
}

func TestFlowController_DoParallel(t *testing.T) {
	kconf := karta.NewConfig().WithWorkerNumber(2)
	queue := karta.NewFakeDelayingQueue(wkq.NewQueue(nil))
	pl := karta.NewPipeline(queue, kconf)
	rl := rl.NewRateLimiter(rl.NewConfig().WithRate(10).WithBurst(1))
	fconf := regula.NewConfig().WithCallback(newTestCallback()).WithRateLimiter(rl)
	fc := regula.NewFlowController(pl, fconf)

	defer fc.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := fc.Do(func(msg any) (any, error) {
				fmt.Printf("msg: %v -> %v\n", msg, v)
				return msg, nil
			}, "test")
			assert.NoError(t, err, "fc.Do should not return error")
		}()
	}
	wg.Wait()

	time.Sleep(time.Second * 2)
}
func TestFlowController_DoUniqQueue(t *testing.T) {
	kconf := karta.NewConfig().WithWorkerNumber(2)
	queue := wkq.NewDelayingQueue(nil)
	pl := karta.NewPipeline(queue, kconf)
	fc := regula.NewFlowController(pl, nil)

	defer fc.Stop()

	err := fc.Do(func(msg any) (any, error) {
		fmt.Printf("msg: %v\n", msg)
		return msg, nil
	}, "test")
	assert.NoError(t, err, "fc.Do should not return error")

	time.Sleep(time.Second)
}

func TestFlowController_DoUniqQueueAfter(t *testing.T) {
	kconf := karta.NewConfig().WithWorkerNumber(2)
	queue := wkq.NewDelayingQueue(nil)
	pl := karta.NewPipeline(queue, kconf)
	rl := rl.NewRateLimiter(rl.NewConfig().WithRate(10).WithBurst(1))
	fconf := regula.NewConfig().WithCallback(newTestCallback()).WithRateLimiter(rl)
	fc := regula.NewFlowController(pl, fconf)

	defer fc.Stop()

	for i := 0; i < 10; i++ {
		v := i
		err := fc.Do(func(msg any) (any, error) {
			fmt.Printf("msg: %v -> %v\n", msg, v)
			return msg, nil
		}, "test")
		assert.NoError(t, err, "fc.Do should not return error")
	}

	time.Sleep(time.Second * 2)
}

func TestFlowController_DoUniqQueueParallel(t *testing.T) {
	kconf := karta.NewConfig().WithWorkerNumber(2)
	queue := wkq.NewDelayingQueue(nil)
	pl := karta.NewPipeline(queue, kconf)
	rl := rl.NewRateLimiter(rl.NewConfig().WithRate(10).WithBurst(1))
	fconf := regula.NewConfig().WithCallback(newTestCallback()).WithRateLimiter(rl)
	fc := regula.NewFlowController(pl, fconf)

	defer fc.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := fc.Do(func(msg any) (any, error) {
				fmt.Printf("msg: %v -> %v\n", msg, v)
				return msg, nil
			}, "test")
			assert.NoError(t, err, "fc.Do should not return error")
		}()
	}
	wg.Wait()

	time.Sleep(time.Second * 2)
}
