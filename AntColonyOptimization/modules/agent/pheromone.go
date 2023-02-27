package agent

import (
  "errors"
  "fmt"
  "os"
)


type ReduceRate float32
type PheromoneType float32
type CitySize int

type PheromoneInteface interface {
  Init(PheromoneType)
  Reduce()
  AddPheromone(int, int, PheromoneType)
}


type Pheromone struct {
  ReduceRate ReduceRate
  CitySize CitySize
  Pheromone [][]PheromoneType 
}


func (p *Pheromone) Init(initialPheromone PheromoneType){
  for i := 0; i<int(p.CitySize)-1; i++{
    p.Pheromone = append(p.Pheromone, make([]PheromoneType, int(p.CitySize)-i-1))
  }

  for i := 0; i<int(p.CitySize)-1; i++{
    for j := 0; j<int(p.CitySize)-i-1; j++{
      p.Pheromone[i][j] = initialPheromone
    }
  }
}


func (p *Pheromone) Reduce(){
  for i := 0; i<int(p.CitySize)-1; i++{
    for j := i+1; j<int(p.CitySize)-i; j++{
      p.Pheromone[i][j] = p.Pheromone[i][j] * PheromoneType(p.ReduceRate)
    }
  }
}

func swapCityIndex(c1 int, c2 int) (int, int){
  if c1 < c2{
    return c1, c2
  }else{
    return c2, c1
  }
}


func isInvalidCityPair(c1 int, c2 int) error {
  if c1 == c2{
    return errors.New("Can't specify the same city")
  }
  return nil
}


func (p *Pheromone) AddPheromone(c1 int, c2 int, additionalPheromone PheromoneType){
  cityPairErr := isInvalidCityPair(c1, c2)
  if cityPairErr != nil{
    fmt.Println(cityPairErr)
    os.Exit(1)
  }
  c1, c2 = swapCityIndex(c1, c2)
  p.Pheromone[c1][c2] += additionalPheromone
}
