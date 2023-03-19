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
	ebitenutil.DrawCircle(screen, p.pos.x, p.pos.y, 3, color.RGBA{204, 156, 255, 255})
}

func (p *Particle) update(dt float64) {
	p.velocity = p.velocity.addv(p.force.dividef(p.mass).multiplyf(dt))
	p.pos = p.pos.addv(p.velocity.multiplyf(dt))
}

func (p *Particle) connect(screen *ebiten.Image, pos_1 *Vector, pos_2 *Vector) {
	ebitenutil.DrawLine(screen, pos_1.x, pos_1.y, pos_2.x, pos_2.y, color.RGBA{204, 156, 255, 255})
}

func (p *Particle) collide(screen *ebiten.Image) bool {
	if p.pos.x < float64(screen.Bounds().Max.X-10) && p.pos.y < float64(screen.Bounds().Max.Y-10) && p.pos.x > 10.0 && p.pos.y > 10.0 {
		return true
	}
	return false
}
