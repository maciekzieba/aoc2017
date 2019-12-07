package main

import (
	"fmt"

	"github.com/maciekzieba/aoc2017/day20/domain"
	"github.com/maciekzieba/aoc2017/day20/input"
)

func main() {
	particles := input.LoadData()

	for j := 0; j < 10000; j++ {
		for i := range particles {
			particles[i].Tick()
		}
	}

	cp := domain.Particle{}
	for _, p := range particles {
		if cp.Distance == 0 || p.Distance < cp.Distance {
			cp = p
		}
	}

	fmt.Printf("Result:%d", cp.ID)
}
