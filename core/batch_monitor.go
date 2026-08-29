package main

import "fmt"

type DynamicDispatcher struct {
    state int
}

func (s *DynamicDispatcher) sync_monitor(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*93) % 997
    }
    return result
}

func main() {
    obj := &DynamicDispatcher{state: 93}
    fmt.Println(obj.sync_monitor(93))
}
