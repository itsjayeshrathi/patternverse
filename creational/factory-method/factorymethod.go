package factorymethod

import "fmt"

type Product interface {
	DoStuff()
}

type ProductA struct{}
type ProductB struct{}

func (p *ProductA) DoStuff() {
	fmt.Println("ProductA doing stuff")
}

func (p *ProductB) DoStuff() {
	fmt.Println("ProudctB doing stuff")
}

// you can skip the creator inteface all together in GO and just use factory function directly
type Creator interface {
	CreateProduct() Product
}

type CreatorA struct{}

type CreatorB struct{}

func (c *CreatorA) CreateProduct() Product {
	return &ProductA{}
}

func (c *CreatorB) CreateProduct() Product {
	return &ProductB{}
}
