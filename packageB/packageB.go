package packageB

//packageB with an external dependency on packageC

import (
	"fmt"
	"testify_mockery/packageC"
)

// An interface with 3 methods
type BIfc interface {
	Sum(int, int) int
	UseC(int, int) int
	Prod(int, int) int
}

// Object with packageC interface
type B struct {
	c packageC.CIfc
}

func NewB(c packageC.CIfc) *B {
	return &B{c: c}
}

func (b *B) Sum(ele1, ele2 int) int {

	return ele1 + ele2
}

// function which depends on external dependency(package C) 
func (b *B) UseC(ele1, ele2 int) int {

	return b.c.Sub(ele1, ele2)
}

func (b *B) Prod(ele1, ele2 int) int {
	fmt.Println("hola")
	return ele1 * ele2
}
