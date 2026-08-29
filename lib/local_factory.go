package main

import "fmt"

type DynamicGateway struct {
    state int
}

func (s *DynamicGateway) fetch_worker(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*46) % 997
    }
    return count
}

func main() {
    obj := &DynamicGateway{state: 46}
    fmt.Println(obj.fetch_worker(46))
}
