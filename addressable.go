package main

import "fmt"

type gopher struct {
  name string
}

func main() {
  gs := map[string]gopher{}
  gs["tall"] = gopher{"TallGopher"}
  fmt.Println("map: ", gs)
  gs["tall"].name = "change"
  fmt.Println("map: ", gs)
}
