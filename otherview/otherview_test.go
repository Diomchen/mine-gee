package otherview

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestX001(t *testing.T) {
	t.Log("x001_security_single_test")

	s1 := GetInstance()
	s2 := GetInstance()

	assert.Equal(t, s1, s2)
}

func TestX002(t *testing.T) {
	t.Log("x002_security_single_test")
	wg := new(sync.WaitGroup)
	wg.Add(2)

	ch := make(chan int, 1)
	go func() {
		defer wg.Done()
		Producer(ch)
	}()

	go func() {
		defer wg.Done()
		Consumer(ch)
	}()

	wg.Wait()

}

func TestX003(t *testing.T) {
	PrintAlternate()
}

func TestX004(t *testing.T) {
	assert.Equal(t, true, JudgeDiffStr("asd"))
	assert.Equal(t, false, JudgeDiffStr("asdd"))

}

func TestX005(t *testing.T) {
	t.Log("x005_multi_consumer_producer_test")
	wg := new(sync.WaitGroup)

	queue := NewMessageQueue[string](10)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(producerId int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				msg := fmt.Sprintf("producer %d, message %d", producerId, j)
				queue.EnqueueBlocking(msg)
				fmt.Println(msg)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(producerId int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				msg := queue.DequeueBlocking()
				// if err != nil {
				// 	fmt.Println(err)

				// 	return
				// }
				fmt.Printf("consumer %d, message %s\n", producerId, msg)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	queue.Close()

	wg.Wait()
}
