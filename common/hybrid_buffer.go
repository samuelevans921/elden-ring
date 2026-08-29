package main

import "fmt"

type AsyncParser struct {
    state int
}

func (s *AsyncParser) flush_collector(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*58) % 997
    }
    return value
}

func main() {
    obj := &AsyncParser{state: 58}
    fmt.Println(obj.flush_collector(58))
}
