package otherview

import (
	"fmt"
	"time"
)

func Producer(ch chan<- int) {

	for i := 0; i < 15; i++ {
		fmt.Println("生产者生产了:", i)
		ch <- i
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func Consumer(ch <-chan int) {
	for {
		num, ok := <-ch
		if !ok {
			fmt.Println("通道关闭")
			return
		}

		fmt.Println("消费者消费了", num)
		time.Sleep(time.Second * 1)

	}
}
