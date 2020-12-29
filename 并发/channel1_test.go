package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello, world"
	done <- true
}

// <-done会阻塞
func Test_2(t *testing.T) {
	go aGoroutine()
	<-done
	println(msg)
}

func aGoroutine1() {
	msg = "hello, world"
	<-done
}

// 对于非缓冲型的 channel，第一个 receive 一定 happened before 第一个 send finished。也就是说，
//在 done <- true 完成之前，<-done 就已经发生了，也就意味着 msg 已经被赋上值了，最终也会打印出 hello, world
func Test_3(t *testing.T) {
	go aGoroutine1()
	done <- true
	println(msg)
}

/*// 控制并发数
func Test_4(t *testing.T) {
	limit := make(chan int, 3)
	for _, w := range work {
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
	select {}
}*/

// 生产者消费着
func Test_5(t *testing.T) {
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// 死锁问题
func f1(in chan int) {
	fmt.Println(<-in)
}

func Test_1(t *testing.T) {
	out := make(chan int)
	out <- 2
	go f1(out) // 和上面对调顺序后可以通过
}

// 控制 Goroutine 的并发数量的方式
func Test_4(t *testing.T) {
	poolCount := 3
	wg := sync.WaitGroup{}
	jobsChan := make(chan int, poolCount)
	for i := 0; i < poolCount; i++ {
		go func() {
			for job := range jobsChan {
				fmt.Println(job)
				time.Sleep(time.Second)
				wg.Done()
			}
		}()
	}

	jobCount := 10
	for i := 0; i < jobCount; i++ {
		jobsChan <- i
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("done!")
}

// channel是否可比较	`
func Test_6(t *testing.T) {
	c1 := make(chan int)
	c2 := c1

	fmt.Println(c1 == c2)
}
