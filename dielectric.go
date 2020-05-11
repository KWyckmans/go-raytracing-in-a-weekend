package main

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

type Dielectric struct {
	refIdx float32
}

func (d Dielectric) shlick(cosine float32) float32 {
	var r0 = (1 - d.refIdx) / (1 + d.refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*float32(math.Pow((float64(1-cosine)), 5))
}

func refract(v mgl32.Vec3, n mgl32.Vec3, niOverNt float32, refracted *mgl32.Vec3) bool {
	var uv mgl32.Vec3 = v.Normalize()
	var dt float32 = uv.Dot(n)
	var discriminant float32 = 1.0 - niOverNt*(1-dt*dt)

	if discriminant > 0 {
		*refracted = uv.Sub(n.Mul(dt)).Mul(niOverNt).Sub(n.Mul(float32(math.Sqrt(float64(discriminant)))))
		return true
	}

	return false

}

func (d Dielectric) Scatter(rIn Ray, rec HitRecord, attenuation *mgl32.Vec3, scattered *Ray) bool {
	var outwardNormal mgl32.Vec3

	// TODO: This will "Magically use the reflect method defined in metal.go. Behaviour should be extracted?"
	var reflected = reflect(rIn.Direction(), rec.normal)

	var niOverNt float32
	*attenuation = mgl32.Vec3{1.0, 1.0, 1.0}
	var refracted mgl32.Vec3

	var reflectProb float32
	var cosine float32

	if rIn.Direction().Dot(rec.normal) > 0 {
		outwardNormal = mgl32.Vec3{-rec.normal.X(), -rec.normal.Y(), -rec.normal.Z()}
		niOverNt = d.refIdx
		cosine = d.refIdx * rIn.Direction().Dot(rec.normal) / rIn.Direction().Len()
	} else {
		outwardNormal = rec.normal
		niOverNt = 1.0 / d.refIdx
		cosine = -rIn.Direction().Dot(rec.normal) / rIn.Direction().Len()
	}

	if refract(rIn.Direction(), outwardNormal, niOverNt, &refracted) {
		reflectProb = d.shlick(cosine)
	} else {
		reflectProb = 1.0
	}

	if rand.Float32() < reflectProb {
		*scattered = Ray{rec.p, reflected}
	} else {
		*scattered = Ray{rec.p, refracted}
	}

	return true

}
