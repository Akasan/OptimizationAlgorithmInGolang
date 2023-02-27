package agent


type ReduceRate float32
type PheromoneType float32
type CitySize int

type Pheromone struct {
  ReduceRate ReduceRate
  CitySize CitySize
  Pheromone [][]PheromoneType 
}


func (p *Pheromone) Init(initialPheromone PheromoneType){
  for i := 0; i<int(p.CitySize); i++{
    p.Pheromone = append(p.Pheromone, make([]PheromoneType, p.CitySize))
  }

  for i := 0; i<int(p.CitySize); i++{
    for j := i+1; j<int(p.CitySize); j++{
      p.Pheromone[i][j] = initialPheromone
      p.Pheromone[j][i] = initialPheromone
    }
  }
}


func (p *Pheromone) Reduce(){
  for i := 0; i<int(p.CitySize); i++{
    for j := i+1; j<int(p.CitySize); j++{
      p.Pheromone[i][j] = p.Pheromone[i][j] * PheromoneType(p.ReduceRate)
      p.Pheromone[j][i] = p.Pheromone[j][i] * PheromoneType(p.ReduceRate)
    }
  }
}

func (p *Pheromone) AddPheromone(c1 int, c2 int, additionalPheromone PheromoneType){
  p.Pheromone[c1][c2] += additionalPheromone
  p.Pheromone[c2][c1] += additionalPheromone
}
