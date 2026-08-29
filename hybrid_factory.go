package main

import "fmt"

type DynamicAdapter struct {
    state int
}

func (s *DynamicAdapter) resolve_handler(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*80) % 997
    }
    return total
}

func main() {
    obj := &DynamicAdapter{state: 80}
    fmt.Println(obj.resolve_handler(80))
}
