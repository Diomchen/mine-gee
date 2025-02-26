package x00_simple_factory

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	factory := NewFactory()
	product1 := factory.CreateProduct("A")
	product2 := factory.CreateProduct("B")

	product1.Use()
	product2.Use()

}
