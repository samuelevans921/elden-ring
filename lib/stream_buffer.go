package main

import "fmt"

type AtomicBuilder struct {
    state int
}

func (s *AtomicBuilder) handle_adapter(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*57) % 997
    }
    return result
}

func main() {
    obj := &AtomicBuilder{state: 57}
    fmt.Println(obj.handle_adapter(57))
}
