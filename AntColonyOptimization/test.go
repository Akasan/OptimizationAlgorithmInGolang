package main

import (
	"AntColonyOptimization/modules/agent"
	file "AntColonyOptimization/modules/file"
	"fmt"
)

func main() {
	colony := agent.Colony{
		ColonySize: 10,
	}
	pheromone := agent.Pheromone{
		ReduceRate: 0.99,
		CitySize:   10,
	}
	pheromone.Init(0.5)
	pheromone.Reduce()
	colony.Initialize()
	fmt.Println(pheromone.Pheromone)
	file.LoadCityFile("./data/kroA100.tsp")
}
