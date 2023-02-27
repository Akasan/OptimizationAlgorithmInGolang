package agent

import (
  "errors"
  "fmt"
  "os"
)


type PheromoneInteface interface {
  Init(float32)
  Reduce()
  AddPheromone(int, int, float32)
}


type Pheromone struct {
  ReduceRate float32
  CitySize int
  Pheromone [][]float32 
}


func (p *Pheromone) Init(initialPheromone float32){
  for i := 0; i<p.CitySize-1; i++{
    p.Pheromone = append(p.Pheromone, make([]float32, p.CitySize-i-1))
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
      p.Pheromone[idx1][idx2] *= p.ReduceRate
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


func (p *Pheromone) AddPheromone(c1 int, c2 int, additionalPheromone float32){
  cityPairErr := isInvalidCityPair(c1, c2)
  if cityPairErr != nil{
    fmt.Println(cityPairErr)
    os.Exit(1)
  }
  c1, c2 = swapCityIndex(c1, c2)
  p.Pheromone[c1][c2] += additionalPheromone
}
