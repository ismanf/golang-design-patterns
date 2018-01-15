package visitor

import (
	"fmt"
)

// Abstractions ...
type Visitor interface {
	VisitProductWithHighTask(ProductWithHighTax) float32
	VisitProductWithLowTask(ProductWithLowTax) float32
}

type Visitable interface {
	Accept(Visitor) float32
}

// Visitor
type TaxVisitor struct {}
func (v TaxVisitor) VisitProductWithHighTask(p ProductWithHighTax) float32 {
	fmt.Println("Visiting Product With High Tax ...")
	return p.Price * 1.3
}
func (v TaxVisitor) VisitProductWithLowTask(p ProductWithLowTax) float32 {
	fmt.Println("Visiting Product With Low Tax ...")
	return p.Price * 1.1
}

// Visitable objects...
type ProductWithHighTax struct {
	Price float32
}
func (p ProductWithHighTax) Accept(v Visitor) float32 {
	return v.VisitProductWithHighTask(p)
}

type ProductWithLowTax struct{
	Price float32
}
func (p ProductWithLowTax) Accept(v Visitor) float32 {
	return v.VisitProductWithLowTask(p)
}