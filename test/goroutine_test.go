package test

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	wg      sync.WaitGroup
	counter int64
	mutex   sync.Mutex
)

func TestTask1(t *testing.T) {

	runtime.GOMAXPROCS(1)
	wg.Add(2)

	t.Log("Creating two goroutines to select prime numbers")
	go selectPrime("A: ")
	go selectPrime("B: ")

	t.Log("Waiting for goroutines to finish")
	wg.Wait()
}

func selectPrime(prefix string) {
	defer wg.Done()
next:
	for i := 2; i <= 5000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue next
			}
		}
		fmt.Println(prefix, i)
	}
	fmt.Println(prefix, "done")
}

func TestTask2(t *testing.T) {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	t.Log("Creating two goroutines to increment counter")
	go incCounter(1)
	go incCounter(2)

	t.Log("Waiting for goroutines to finish")
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}

func incCounter(id int) {
	fmt.Println("Starting goroutine", id)
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}

}

func TestTask3(t *testing.T) {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	t.Log("Creating two goroutines to increment counter using mutex")
	go incCounterMutex(1)
	go incCounterMutex(2)
	t.Log("Waiting for goroutines to finish")
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}

func incCounterMutex(id int) {
	defer wg.Done()
	fmt.Printf("`incCounterMutex` goroutine %d started\n", id)

	for i := 0; i < 2; i++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()

			value++
			counter = value
		}

		mutex.Unlock()

	}
	fmt.Printf("`incCounterMutex` goroutine %d finished\n", id)
}

func TestTask4(t *testing.T) {
	runtime.GOMAXPROCS(1)
	court := make(chan int)
	wg.Add(2)

	go player("A", court)
	go player("B", court)

	// 发球
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	fmt.Printf("Player: %s started\n", name)
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player :%s won the game\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player: %s missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player: %s hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
