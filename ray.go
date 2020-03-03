package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Ray represents a ray originating at A and pointing to B
type Ray struct {
	A mgl32.Vec3
	B mgl32.Vec3
}

// NewRay Creates a ray that originates at A and points to B
func NewRay(A mgl32.Vec3, B mgl32.Vec3) Ray {
	r := Ray{A: A, B: B}
	return r
}

// Origin returns the origin of the ray
func (r Ray) Origin() mgl32.Vec3 {
	return r.A
}

// Direction returns the direction the ray points at
func (r Ray) Direction() mgl32.Vec3 {
	return r.B
}

// P returns a point B on the ray
func (r Ray) P(t float32) mgl32.Vec3 {
	return r.A.Add(r.B.Mul(t))
}
