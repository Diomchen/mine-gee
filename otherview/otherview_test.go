package otherview

import (
	"sync"
	"testing"

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
