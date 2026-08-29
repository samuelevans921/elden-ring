package main

import "fmt"

type SmartCollector struct {
    state int
}

func (s *SmartCollector) fetch_router(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*85) % 997
    }
    return result
}

func main() {
    obj := &SmartCollector{state: 85}
    fmt.Println(obj.fetch_router(85))
}
