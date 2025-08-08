package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(jobs <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range jobs {
		fmt.Printf("🤔worker %v received job %v \n", id, val)
		time.Sleep(1 * time.Second)
		fmt.Printf("✅worker %v done job %v\n", id, val)
	}
}

func factory() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(jobs, i, &wg)
	}

	jobs <- 11
	jobs <- 12
	jobs <- 13
	jobs <- 14
	jobs <- 15
	jobs <- 16
	jobs <- 17
	close(jobs)

	wg.Wait()
}

func TestWorker(t *testing.T) {
	factory()
}
