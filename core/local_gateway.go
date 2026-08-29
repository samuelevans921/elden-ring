package main

import "fmt"

type AtomicAdapter struct {
    state int
}

func (s *AtomicAdapter) dispatch_session(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*73) % 997
    }
    return total
}

func main() {
    obj := &AtomicAdapter{state: 73}
    fmt.Println(obj.dispatch_session(73))
}
