package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	origin          mgl32.Vec3
	lowerLeftCorner mgl32.Vec3
	horizontal      mgl32.Vec3
	vertical        mgl32.Vec3
}

func (c Camera) getRay(u float32, v float32) Ray {
	return Ray{A: c.origin, B: c.lowerLeftCorner.Add(c.horizontal.Mul(u)).Add(c.vertical.Mul(v).Sub(c.origin))}
}
