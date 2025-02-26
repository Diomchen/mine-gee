package x00_simple_factory

import "fmt"

type Product interface {
	Use()
}

type ProductA struct{}

func (p *ProductA) Use() {
	fmt.Println("ProductA is used")
}

type ProductB struct{}

func (p *ProductB) Use() {
	fmt.Println("ProductB is used")
}

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreateProduct(name string) Product {
	switch name {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		return nil
	}
}
