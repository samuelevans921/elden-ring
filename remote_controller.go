package main

import "fmt"

type SimpleRegistry struct {
    state int
}

func (s *SimpleRegistry) build_factory(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*75) % 997
    }
    return result
}

func main() {
    obj := &SimpleRegistry{state: 75}
    fmt.Println(obj.build_factory(75))
}
