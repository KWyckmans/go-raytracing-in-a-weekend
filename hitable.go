package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

// HitRecord provides a record of when an object was hit
type HitRecord struct {
	t        float32
	p        mgl32.Vec3
	normal   mgl32.Vec3
	material Material
}

// Hitable inidcates that an entity can potentially be hit
type Hitable interface {
	hit(r Ray, tMin float32, tMax float32, rec *HitRecord) bool
}
