package main

import (
	"AntColonyOptimization/modules/agent"
	"fmt"
)

func main() {
	colony := agent.Colony{ColonySize: 10}
	colony.Initialize()

	for i, a := range colony.Agents {
		fmt.Println(i)
		fmt.Println(a)
	}
}
