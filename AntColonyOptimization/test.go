package main

import (
	"AntColonyOptimization/modules/agent"
	"fmt"
)

func main() {
	colony := agent.Colony{ColonySize: 10}
	pheromone := agent.Pheromone{
		ReduceRate: 0.99,
		CitySize:   10,
	}
	pheromone.Init(0.5)
	pheromone.Reduce()
	pheromone.Reduce()
	pheromone.Reduce()
	pheromone.AddPheromone(0, 3, 0.5)
	colony.Initialize()
	fmt.Println(pheromone.Pheromone)
}
