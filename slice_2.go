package main

import (  
    "fmt"
    "bytes"
)

func main() {  
    path := []byte("AAAA/BBBBBBBBB")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex:sepIndex]
    dir2 := path[sepIndex+1:]
    fmt.Println("dir1 =>",string(dir1))
    fmt.Println("dir2 =>",string(dir2))

    dir1 = append(dir1,"suffix"...)
    path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})

    fmt.Println("dir1 =>",string(dir1))
    fmt.Println("dir2 =>",string(dir2))

    fmt.Println("new path =>",string(path))
}
