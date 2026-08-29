package main

import "fmt"

type SmartBuilder struct {
    state int
}

func (s *SmartBuilder) load_adapter(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*10) % 997
    }
    return acc
}

func main() {
    obj := &SmartBuilder{state: 10}
    fmt.Println(obj.load_adapter(10))
}
