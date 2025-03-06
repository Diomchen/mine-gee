package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MySlice[T int | float64] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func TestGenerics(t *testing.T) {
	t.Log("Generics test")
	s := MySlice[int]{1, 2, 3, 4, 5}
	s.Sum()
	assert.Equal(t, 15, s.Sum())

	f := MySlice[float64]{1.1, 2.2, 3.3, 4.4, 5.5}
	f.Sum()

	assert.Equal(t, 15, s.Sum())
}
