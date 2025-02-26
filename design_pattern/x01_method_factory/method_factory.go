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

type Creator interface {
	createFactory() Product
}

type FactoryA struct{}

func (f *FactoryA) createFactory() Product {
	return &ProductA{}
}

type FactoryB struct{}

func (f *FactoryB) createFactory() Product {
	return &ProductB{}
}
