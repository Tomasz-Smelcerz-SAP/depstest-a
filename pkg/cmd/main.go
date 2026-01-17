package main

import (
	"fmt"
	"github.com/Tomasz-Smelcerz-SAP/depstest-a/pkg/component"
	"github.com/Tomasz-Smelcerz-SAP/depstest-b/pkg/container"
)

func main() {
	component := component.Component{Name: "MyComponent"}
	container := &container.Container{}
	container.Elements = append(container.Elements, component)

	err := container.Visit(visitComponent)
	if err != nil {
		fmt.Println("Error during visit:", err)
	}
	fmt.Println("Done.")
}

func visitComponent(c *component.Component) error {
	fmt.Println("Visiting component:", c.Name)
	return nil
}

