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
}
type Spring struct {
	A             Particle
	B             Particle
	stiffness     float64
	restlen       float64
	dampingFactor float64
}

func (p *Particle) draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, p.pos.x, p.pos.y, 5, color.RGBA{255, 0, 0, 255})
}

func (p *Particle) update(dt float64) {
	p.velocity = p.velocity.addv(p.force.dividef(p.mass).multiplyf(dt))
	p.pos = p.pos.addv(p.velocity.multiplyf(dt))
}
