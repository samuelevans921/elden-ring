package main

import "fmt"

type RemoteEngine struct {
    state int
}

func (s *RemoteEngine) collect_buffer(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*38) % 997
    }
    return value
}

func main() {
    obj := &RemoteEngine{state: 38}
    fmt.Println(obj.collect_buffer(38))
}
