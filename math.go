package main

type Vector struct {
	x float64
	y float64
}

func (v Vector) dividef(scalar float64) Vector {
	v.x /= scalar
	v.y /= scalar

	return v
}

func (v Vector) multiplyf(scalar float64) Vector {
	v.x *= scalar
	v.y *= scalar

	return v
}

func (v Vector) multiplyv(vec Vector) Vector {
	v.x *= vec.x
	v.y *= vec.y

	return v
}

func (v Vector) addf(scalar float64) Vector {
	v.x += scalar
	v.y += scalar

	return v
}

func (v Vector) addv(vec Vector) Vector {
	v.x += vec.x
	v.y += vec.y

	return v
}
