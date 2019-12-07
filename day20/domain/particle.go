package domain

import (
	"fmt"
	"math"
)

type (
	Coordinate struct {
		X int
		Y int
		Z int
	}

	Particle struct {
		ID           int
		Position     Coordinate
		Velocity     Coordinate
		Acceleration Coordinate
		Distance     int
	}
)

func (c Coordinate) ToString() string {
	return fmt.Sprintf("x:%d y:%d z:%d ", c.X, c.Y, c.Z)
}

func (c *Coordinate) Add(arg Coordinate) {
	c.X += arg.X
	c.Y += arg.Y
	c.Z += arg.Z
}

func (c Coordinate) CalculateDistance() int {
	return int(math.Abs(float64(c.X))) + int(math.Abs(float64(c.Y))) + int(math.Abs(float64(c.Z)))
}

func (p *Particle) Tick() {
	p.Velocity.Add(p.Acceleration)
	p.Position.Add(p.Velocity)
	p.Distance = p.Position.CalculateDistance()
}

func (p Particle) ToString() string {
	return fmt.Sprintf("id:%d pos:%s dist:%d", p.ID, p.Position.ToString(), p.Distance)
}

func NewParticle(p, v, a Coordinate, ID int) Particle {
	particle := Particle{
		Position:     p,
		Velocity:     v,
		Acceleration: a,
		Distance:     0,
		ID:           ID,
	}
	particle.Distance = particle.Position.CalculateDistance()
	return particle
}
