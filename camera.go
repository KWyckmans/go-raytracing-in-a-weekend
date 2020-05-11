package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

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

	theta      float32
	halfHeight float32
	halfWidth  float32
}

func BuildCamera(lookFrom mgl32.Vec3, lookAt mgl32.Vec3, vup mgl32.Vec3, vfov float32, aspect float32) Camera {
	w := lookFrom.Sub(lookAt).Normalize()
	u := vup.Cross(w).Normalize()
	v := w.Cross(u)

	theta := vfov * math.Pi / 180.0
	halfHeight := float32(math.Tan(float64(theta / 2.0)))
	halfWidth := aspect * halfHeight
	origin := lookFrom

	return camera{
		origin:          lookFrom,
		w:               w,
		u:               u,
		v:               v,
		theta:           theta,
		halfHeight:      halfHeight,
		halfWidth:       halfWidth,
		lowerLeftCorner: origin.Sub(u.Mul(halfWidth)).Sub(v.Mul(halfHeight)).Sub(w),
		horizontal:      u.Mul(2 * halfWidth),
		vertical:        v.Mul(2 * halfHeight),
	}
}

func (c camera) GetRay(u float32, v float32) Ray {
	return Ray{A: c.origin, B: c.lowerLeftCorner.Add(c.horizontal.Mul(u)).Add(c.vertical.Mul(v).Sub(c.origin))}
}
