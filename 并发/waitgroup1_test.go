package main

import (
	"fmt"
	"sync"
	"testing"
)

// 主协程如何等其余协程完再操作
func Test_11(t *testing.T) {

	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	wg.Wait() // 等待，直到计数为0
}

func worker(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done()
	for {
		select {
		case <-cannel:
			return
		default:
			fmt.Println("hello")
		}
	}
}

// 如何防止goroutine泄露 WaitGroup保证其他协程能够完成清理工作
func Test_12(t *testing.T) {
	cacle := make(chan bool)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cacle)
	}
	close(cacle)
	wg.Wait()
}
