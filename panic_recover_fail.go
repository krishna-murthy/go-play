package main

import "fmt"

func doRecover() {  
    fmt.Println("recovered =>",recover()) //prints: recovered => <nil>
}

func main() {  
    defer func() {
        doRecover() //panic is not recovered
    }()

    panic("not good")
}
