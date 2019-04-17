package main

import "fmt"

type gopher struct {
  name string
}

func (g *gopher) Name() {
  fmt.Println("My name is gopher.")
}

func main() {
  var g *gopher
  fmt.Printf("content: %+v %+p\n", g, g)
  g.Name()
}
