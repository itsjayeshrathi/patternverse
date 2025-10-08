package abstractfactory

type ProductA interface {
	UseA() string
}

type ProductB interface {
	UseB() string
}

type ProductA1 struct{}

func (p ProductA1) UseA() string {
	return "Using Product A1"
}

type ProductA2 struct{}

func (p ProductA2) UseA() string {
	return "Using Product A2"
}

type ProductB1 struct{}

func (p ProductB1) UseB() string {
	return "Using Product B1"
}

type ProductB2 struct{}

func (p ProductB2) UseB() string {
	return "Using Product B2"
}

type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type Factory1 struct{}

func (f Factory1) CreateProductA() ProductA {
	return ProductA2{}
}

func (f Factory1) CreateProductB() ProductB {
	return ProductB2{}
}
