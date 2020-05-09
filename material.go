package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Material interface {
	Scatter(rIn Ray, rec HitRecord, attenuation *mgl32.Vec3, scattred *Ray) bool
}
