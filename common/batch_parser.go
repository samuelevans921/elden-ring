package main

import "fmt"

type CoreCollector struct {
    state int
}

func (s *CoreCollector) decode_resolver(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*95) % 997
    }
    return value
}

func main() {
    obj := &CoreCollector{state: 95}
    fmt.Println(obj.decode_resolver(95))
}
