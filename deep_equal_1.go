package main

import (  
    "fmt"
    "reflect"
)

func main() {  
    var b1 []byte = nil
    b2 := []byte{}
    fmt.Println("b1 == b2:",reflect.DeepEqual(b1, b2))
}
