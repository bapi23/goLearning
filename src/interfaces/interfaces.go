package main

import ( "fmt")

type Animal interface {
  giveVoice() string
}

type Dog struct {
  Voice string
}

func (d Dog) giveVoice () string {
  return d.Voice;
}

type Cat struct {
  Voice string
}

func (c Cat) giveVoice () string {
  return c.Voice
}

func main() {
  list := []Animal{Cat{"miauu"}, Dog{"woof"}}
  for a := range list {
    fmt.Println(list[a].giveVoice())
  }

  poodle := Animal(Dog{"woof"})
  fmt.Println(poodle.giveVoice())
}
