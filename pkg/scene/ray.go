package scene

import "github.com/golang/geo/r3"

type Ray struct {
	Origin    r3.Vector
	Direction r3.Vector
}

func NewRay(origin r3.Vector, direction r3.Vector) Ray {
	return Ray{Origin: origin, Direction: direction}
}
