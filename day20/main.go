package main

import (
	"fmt"
	"sync"

	"github.com/maciekzieba/aoc2017/day20/domain"
	"github.com/maciekzieba/aoc2017/day20/input"
)

func main() {
	particles := input.LoadData()
	var wg sync.WaitGroup
	wg.Add(len(particles))

	for i := range particles {
		go func(pId int) {
			for j := 0; j < 10000; j++ {
				particles[pId].Tick()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	cp := domain.Particle{}
	for _, p := range particles {
		if cp.Distance == 0 || p.Distance < cp.Distance {
			cp = p
		}
	}

	fmt.Println("Result:", cp.ID)
}
