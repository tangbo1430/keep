package 如何限制协程数量

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 参考GOMAXPROCS方法
// import _ "github.com/go-pay/errgroup"

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	p.wg.Add(delta)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func TestMaxGo(t *testing.T) {
	pool := NewPool(10)
	fmt.Println("the NumGoroutine begin is:", runtime.NumGoroutine())
	for i := 0; i < 30; i++ {
		pool.Add(1)
		go func() {
			fmt.Println("hahahaha")
			fmt.Println("the NumGoroutine continue is:", runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	fmt.Println("the NumGoroutine done is:", runtime.NumGoroutine())
}
