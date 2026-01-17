package component

import (
	"fmt"
	//"github.com/Tomasz-Smelcerz-SAP/depstest-b/pkg/container" //import cycle not allowed!
)

const ComponentVersion = "1.0.0"

type Component struct {
	Name string
}

func (c *Component) PrintName() {
	fmt.Println("Component Name:", c.Name)
}
