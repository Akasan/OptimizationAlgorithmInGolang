package agent

type Colony struct {
  ColonySize int
  Agents     []Agent
}

func (c *Colony) Initialize(){
  for id :=0 ; id < c.ColonySize; id++{
    c.Agents = append(c.Agents, Agent{
      Id: id,
    })    
  }
}
