package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	prevTime     int64
	testParticle Particle
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
	numberOfParticles := 12

	var particles []*Particle = make([]*Particle, numberOfParticles)

	for i := range particles {
		particles[i] = new(Particle)
	}

	g.testParticle.draw(screen)
	g.testParticle.update(dt)

	g.prevTime = now
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Jelly :3")

	game := Game{testParticle: Particle{force: Vector{0, 9.8}, mass: 1, pos: Vector{160, 0}}, prevTime: time.Now().UnixMilli()}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
