package main

import "github.com/go-gl/mathgl/mgl32"

type Lambertian struct {
	albedo mgl32.Vec3
}

func (l Lambertian) Scatter(rIn Ray, rec HitRecord, attenuation *mgl32.Vec3, scattered *Ray) bool {
	var target mgl32.Vec3 = rec.p.Add(rec.normal).Add(randInUnitSpehre())
	*scattered = Ray{rec.p, target.Sub(rec.p)}
	*attenuation = l.albedo
	return true
}
