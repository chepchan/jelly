package main

import (
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	prevTime          int64
	particles         []*Particle
	numberOfParticles int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	now := time.Now().UnixMilli()
	// Delta time in milliseconds
	dtMs := now - g.prevTime
	// Delta time in seconds
	dt := (float64(dtMs) / 1000.0)

	for i := 0; i < g.numberOfParticles-1; i++ {
		g.particles[i].draw(screen)
		g.particles[i].update(dt)
		g.particles[i].connect(screen, &g.particles[i].pos, &g.particles[i+1].pos)
	}
	g.particles[g.numberOfParticles-1].draw(screen)
	g.particles[0].connect(screen, &g.particles[g.numberOfParticles-1].pos, &g.particles[0].pos)
	g.particles[g.numberOfParticles-1].update(dt)

	g.prevTime = now
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1400, 950
}

func main() {
	numberOfParticles := 12
	var particles []*Particle = make([]*Particle, numberOfParticles)

	tau := 6.28
	step := tau / float64(numberOfParticles)
	radius := 100.0
	center := Vector{300.0, 150.0}

	for i := 0; i < numberOfParticles; i++ {
		angle := step * float64(i)
		x := center.x + radius*math.Cos(angle)
		y := center.y + radius*math.Sin(angle)
		particles[i] = &Particle{pos: Vector{x, y}, force: Vector{0.0, 9.8}, velocity: Vector{0.0, 0.0}, mass: 0.1}
	}

	ebiten.SetWindowSize(1400, 950)
	ebiten.SetWindowTitle("Jelly :3")

	game := Game{particles: particles, prevTime: time.Now().UnixMilli(), numberOfParticles: numberOfParticles}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
