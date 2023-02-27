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

  for idx1, val := range p.Pheromone{
    for idx2, _ := range val{
      p.Pheromone[idx1][idx2] = initialPheromone
    }
  }
}

func (p *Pheromone) Reduce(){
  for idx1, val := range p.Pheromone{
    for idx2, _ := range val{
      p.Pheromone[idx1][idx2] *= PheromoneType(p.ReduceRate)
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
