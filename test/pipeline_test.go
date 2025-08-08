package test

import (
	"fmt"
	"testing"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func printer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func add(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1
		}
		close(out)
	}()
	return out
}

func TestPipeline(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	gen(nums...)
	sq(gen(nums...))
	printer(add(sq(gen(nums...))))
}
