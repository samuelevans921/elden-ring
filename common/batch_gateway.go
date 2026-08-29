package main

import "fmt"

type BatchDispatcher struct {
    state int
}

func (s *BatchDispatcher) dispatch_builder(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*81) % 997
    }
    return total
}

func main() {
    obj := &BatchDispatcher{state: 81}
    fmt.Println(obj.dispatch_builder(81))
}
