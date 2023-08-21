package packageC
// considering this package as an external dependency and generating mocks for it

import (
	"fmt"
)

// By using mockery generating C Interface mock

//go:generate mockery --name CIfc
// the above line generates mock interfaces for the provided names we can replace '--name CIfc' with '--all' to generate mocks for all interfaces

type CIfc interface {
	Sub(int, int) int
}
type C struct{}

func (c *C) Sub(ele1, ele2 int) int {
	fmt.Println("hoal")
	return ele1 - ele2
}
