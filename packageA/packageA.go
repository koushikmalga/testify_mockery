package packageA

import (
	"testify_mockery/packageB"
)

// An Interface with a method func1
type AIfc interface {
	func1(int, int) int
}

// object with packageB Interface
type A struct {
	b packageB.BIfc
}

func NewA(b packageB.BIfc) *A {
	return &A{b: b}
}

// func1 calling packageB
func (a *A) func1(ele1, ele2 int) int {
	return a.b.Sum(ele1, ele2) + a.b.UseC(ele1, ele2)
}
