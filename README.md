# go-split

A small Go utility for splitting a file into smaller ones.

## Installation
```shell script
$ go get -u github.com/knwoop/go-splitter
```

## how to use
```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/knwoop/go-splitter"
)

func main() {
    f, err := os.Open("test.txt")
    if err != nil{
        log.Fatal(err)
    }
    defer f.Close() 
    b1, b2, err := splitter.Split(f, true, 40, 60)
    if err != nil {
        log.Fatal(err)
    }
    ...
}
```
