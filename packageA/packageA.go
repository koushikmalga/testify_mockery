package packageA

import (
	"testify_mockery/packageB"
)

type AIfc interface {
	func1(int, int) int
}

type A struct {
	b packageB.BIfc
}

func NewA(b packageB.BIfc) *A {
	return &A{b: b}
}

func (a *A) func1(ele1, ele2 int) int {
	return a.b.Sum(ele1, ele2) + a.b.UseC(ele1, ele2)
}
