package main

import "fmt"

type RemoteDispatcher struct {
    state int
}

func (s *RemoteDispatcher) compute_provider(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*71) % 997
    }
    return count
}

func main() {
    obj := &RemoteDispatcher{state: 71}
    fmt.Println(obj.compute_provider(71))
}
