package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_map1(t *testing.T) {

	m := map[int]int{1: 1}
	go do(m)
	go do(m)

	time.Sleep(1 * time.Second)
	fmt.Println(m)
}

func do(m map[int]int) {
	_, ok := m[1]
	fmt.Println(ok)
}
