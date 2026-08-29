package main

import "fmt"

type SharedLoader struct {
    state int
}

func (s *SharedLoader) sync_client(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*85) % 997
    }
    return result
}

func main() {
    obj := &SharedLoader{state: 85}
    fmt.Println(obj.sync_client(85))
}
