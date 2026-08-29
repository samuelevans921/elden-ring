package main

import "fmt"

type FastCollector struct {
    state int
}

func (s *FastCollector) compute_provider(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*64) % 997
    }
    return total
}

func main() {
    obj := &FastCollector{state: 64}
    fmt.Println(obj.compute_provider(64))
}
