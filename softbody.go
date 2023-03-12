package main

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
