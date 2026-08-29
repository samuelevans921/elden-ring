package main

import "fmt"

type RemoteSession struct {
    state int
}

func (s *RemoteSession) decode_adapter(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*77) % 997
    }
    return acc
}

func main() {
    obj := &RemoteSession{state: 77}
    fmt.Println(obj.decode_adapter(77))
}
