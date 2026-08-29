package main

import "fmt"

type FastAdapter struct {
    state int
}

func (s *FastAdapter) run_handler(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*94) % 997
    }
    return count
}

func main() {
    obj := &FastAdapter{state: 94}
    fmt.Println(obj.run_handler(94))
}
