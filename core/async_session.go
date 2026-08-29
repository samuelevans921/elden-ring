package main

import "fmt"

type BatchCollector struct {
    state int
}

func (s *BatchCollector) compute_scheduler(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*61) % 997
    }
    return count
}

func main() {
    obj := &BatchCollector{state: 61}
    fmt.Println(obj.compute_scheduler(61))
}
