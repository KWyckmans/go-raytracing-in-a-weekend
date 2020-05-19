package main

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

func randInUnitDisk() mgl32.Vec3 {
	var p mgl32.Vec3 = mgl32.Vec3{rand.Float32(), rand.Float32(), 0}.Mul(2.0).Sub(mgl32.Vec3{1, 1, 0})

	// if p.Dot(p) < 1.0 {
	// 	return p
	// }

	for ok := true; ok; ok = p.Dot(p) < 1.0 {
		p = mgl32.Vec3{rand.Float32(), rand.Float32(), 0}.Mul(2.0).Sub(mgl32.Vec3{1, 1, 0})
	}

	return p
}

type Camera interface {
	GetRay(u float32, v float32) Ray
}

type camera struct {
	origin          mgl32.Vec3
	v               mgl32.Vec3
	u               mgl32.Vec3
	w               mgl32.Vec3
	lowerLeftCorner mgl32.Vec3
	horizontal      mgl32.Vec3
	vertical        mgl32.Vec3

	lensRadius float32

	theta      float32
	halfHeight float32
	halfWidth  float32
}

func BuildCamera(lookFrom mgl32.Vec3, lookAt mgl32.Vec3, vup mgl32.Vec3, vfov float32, aspect float32, aperture float32, focusDist float32) Camera {

	w := lookFrom.Sub(lookAt).Normalize()
	u := vup.Cross(w).Normalize()
	v := w.Cross(u)

	theta := vfov * math.Pi / 180.0
	halfHeight := float32(math.Tan(float64(theta / 2.0)))
	halfWidth := aspect * halfHeight
	origin := lookFrom

	return camera{
		lensRadius:      aperture / 2.0,
		origin:          lookFrom,
		w:               w,
		u:               u,
		v:               v,
		theta:           theta,
		halfHeight:      halfHeight,
		halfWidth:       halfWidth,
		lowerLeftCorner: origin.Sub(u.Mul(halfWidth * focusDist)).Sub(v.Mul(halfHeight * focusDist)).Sub(w.Mul(focusDist)),
		horizontal:      u.Mul(2 * halfWidth * focusDist),
		vertical:        v.Mul(2 * halfHeight * focusDist),
	}
}

func (c camera) GetRay(s float32, t float32) Ray {
	var rd mgl32.Vec3 = randInUnitDisk().Mul(c.lensRadius)
	var offset mgl32.Vec3 = c.u.Mul(rd.X()).Add(c.v.Mul(rd.Y()))

	return Ray{A: c.origin.Add(offset), B: c.lowerLeftCorner.Add(c.horizontal.Mul(s)).Add(c.vertical.Mul(t).Sub(c.origin).Sub(offset))}
}
