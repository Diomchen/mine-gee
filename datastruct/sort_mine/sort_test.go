package sort_mine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	sorter := &Sorter{
		strategy: &QuickSort{},
		arr:      arr,
	}

	sorter.Execute()

	for _, v := range arr {
		fmt.Printf("%v ", v)
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	sorter := &Sorter{
		strategy: &InsertStrategy{},
		arr:      arr,
	}

	sorter.Execute()

	for _, v := range arr {
		fmt.Printf("%v ", v)
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	sorter := &Sorter{
		strategy: &HeapStrategy{},
		arr:      arr,
	}

	sorter.Execute()

	for _, v := range arr {
		fmt.Printf("%v ", v)
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, arr)
}

func TestShellSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	sorter := &Sorter{
		strategy: &ShellStrategy{},
		arr:      arr,
	}

	sorter.Execute()

	for _, v := range arr {
		fmt.Printf("%v ", v)
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, arr)
}
