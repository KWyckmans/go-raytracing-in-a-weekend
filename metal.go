package main

import "github.com/go-gl/mathgl/mgl32"

type Metal struct {
	albedo mgl32.Vec3
}

func reflect(v mgl32.Vec3, n mgl32.Vec3) mgl32.Vec3 {
	return v.Sub(n.Mul(2.0 * v.Dot(n)))
}

func (m Metal) Scatter(rIn Ray, rec HitRecord, attenuation *mgl32.Vec3, scattered *Ray) bool {
	var reflected mgl32.Vec3 = reflect(rIn.Direction().Mul(1/rIn.Direction().Len()), rec.normal)
	*scattered = Ray{rec.p, reflected}
	*attenuation = m.albedo
	return true
}
