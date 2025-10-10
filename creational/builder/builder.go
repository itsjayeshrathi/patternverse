package builder

import "fmt"

type Product struct {
	partA string
	partB string
	partC string
}

func (p Product) Show() {
	fmt.Printf("Product built with: %s, %s, %s\n", p.partA, p.partB, p.partC)
}

type Builder interface {
	Reset()
	SetPartA()
	SetPartB()
	SetPartC()
	SetPartZ()
	GetResult() Product
}
