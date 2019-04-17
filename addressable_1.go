package main

import "fmt"

type gopher struct {
  name string
}

func main() {
  gl := []gopher{{"tall"}, {"small"}}
  gl[0].name = "asdffds"
  fmt.Println(gl)
}
