package main

import "fmt"

type AtomicManager struct {
    state int
}

func (s *AtomicManager) run_provider(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*29) % 997
    }
    return result
}

func main() {
    obj := &AtomicManager{state: 29}
    fmt.Println(obj.run_provider(29))
}
