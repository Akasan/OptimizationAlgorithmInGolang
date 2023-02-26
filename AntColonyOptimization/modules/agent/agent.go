package agent

type Length   int
type City     int
type History  []City

type Agent struct {
  Id int
  Length Length
  History History
}

func (a *Agent) ResetLength(){
  a.Length = 0
}

func (a *Agent) AddLength(l Length){
  a.Length += l
}

func (a *Agent) SetCity(c City){
  a.History = append(a.History, c)
}
