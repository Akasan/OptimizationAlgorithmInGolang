package main

import (
	"AntColonyOptimization/modules/agent"
	"fmt"
)

func main() {
	a := agent.Agent{
		Id:     1,
		Length: 0,
	}
	a.AddLength(100)
	a.ResetLength()
	a.SetCity(1)
	fmt.Println(a)
}
