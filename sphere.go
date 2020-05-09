package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Sphere struct {
	center mgl32.Vec3
	radius float32
}

// NewSphere instantiates a sphere with a center and radius r
func NewSphere(center mgl32.Vec3, r float32) Sphere {
	s := Sphere{center: center, radius: r}
	return s
}

func (s Sphere) hit(r Ray, tMin float32, tMax float32, rec *HitRecord) bool {
	var oc mgl32.Vec3 = r.Origin().Sub(s.center)

	var a float32 = r.Direction().Dot(r.Direction())
	var b float32 = oc.Dot(r.Direction())
	var c float32 = oc.Dot(oc) - s.radius*s.radius

	var discriminant float32 = b*b - a*c

	if discriminant > 0 {
		var temp float32 = float32((float64(-b) - math.Sqrt(float64(b*b-a*c))) / float64(a))

		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.P(rec.t)
			rec.normal = (rec.p.Sub(s.center)).Mul(1 / s.radius)
			return true
		}

		temp = float32((float64(-b) + math.Sqrt(float64(b*b-a*c))) / float64(a))
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.P(rec.t)
			rec.normal = (rec.p.Sub(s.center)).Mul(1 / s.radius)
			return true
		}
	}
	return false
}
