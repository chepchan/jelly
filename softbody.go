package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Particle struct {
	pos      Vector
	force    Vector
	velocity Vector
	mass     float64
	springs  []*Spring
}
type Spring struct {
	A         *Particle
	B         *Particle
	stiffness float64
	restlen   float64
	// dampingFactor float64
}

func (p *Particle) draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, p.pos.x, p.pos.y, 3, color.RGBA{204, 156, 255, 255})
}

func (p *Particle) update(dt float64) {
	p.velocity = p.velocity.addv(p.force.dividef(p.mass).multiplyf(dt))
	p.pos = p.pos.addv(p.velocity.multiplyf(dt))
}

func (p *Particle) connect(screen *ebiten.Image, pos_1, pos_2 *Vector, stiffness, restlen float64) {
	spring := &Spring{p, nil, stiffness, restlen}
	p.springs = append(p.springs, spring)
	//creating new particle for the second point and connecting it to the spring
	p2 := &Particle{pos: *pos_2, mass: 0}
	// spring.p2 = p2                            idk
	p2.springs = append(p2.springs, spring)
	ebitenutil.DrawLine(screen, pos_1.x, pos_1.y, pos_2.x, pos_2.y, color.RGBA{204, 156, 255, 255})
}

func (p *Particle) collide(screen *ebiten.Image) bool {
	if p.pos.x < float64(screen.Bounds().Max.X-10) && p.pos.y < float64(screen.Bounds().Max.Y-10) && p.pos.x > 10.0 && p.pos.y > 10.0 {
		return true
	}
	return false
}

func (p *Particle) computeForces(particles []*Particle, dt float64) {
	for _, p := range particles {
		p.force = Vector{}
		for _, s := range p.springs {
			delta := s.p1.pos.subtract(s.p2.pos)
			dist := delta.length()
			magnitude := s.stiffness * (dist - s.restlen)
			force := delta.normalize().multiplyf(-magnitude)

			s.p1.force = s.p1.force.addv(force)
			s.p2.force = s.p2.force.addv(force.negate())
		}
	}
}
