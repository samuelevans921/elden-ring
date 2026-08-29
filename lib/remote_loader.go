package main

import "fmt"

type DynamicCollector struct {
    state int
}

func (s *DynamicCollector) fetch_session(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*43) % 997
    }
    return result
}

func main() {
    obj := &DynamicCollector{state: 43}
    fmt.Println(obj.fetch_session(43))
}
