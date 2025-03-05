package otherview

import (
	"fmt"
	"sync"
)

func PrintAlternate() {
	wg := sync.WaitGroup{}
	letter, number := make(chan int, 1), make(chan int, 1)

	wg.Add(1)
	// 打印数字
	go func() {
		defer wg.Done()
		for {
			v, ok := <-number
			if !ok || v == -1 {
				close(number)
				return
			}
			fmt.Printf("%v%v", v+1, v+2)
			letter <- v
		}
	}()

	wg.Add(1)
	// 打印字母
	letter_str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	go func() {
		defer wg.Done()
		for {
			index, ok := <-letter
			print_str := ""
			if ok {
				print_str = letter_str[index : index+2]
				fmt.Printf("%s", print_str)
			}
			if print_str == "" || print_str == "YZ" {
				close(letter)
				number <- -1
				return
			}
			number <- index + 2
		}
	}()

	number <- 0

	wg.Wait()
}
