package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

func color(r Ray, world Hitable, depth int) mgl32.Vec3 {
	var rec HitRecord
	if world.hit(r, 0.001, math.MaxFloat32, &rec) {
		var scattered Ray
		var attenuation mgl32.Vec3

		if depth < 50 && rec.material.Scatter(r, rec, &attenuation, &scattered) {
			// fmt.Println("Scattering rays")
			var col = color(scattered, world, depth+1)
			return mgl32.Vec3{attenuation.X() * col.X(), attenuation.Y() * col.Y(), attenuation.Z() * col.Z()}
		} else {
			return mgl32.Vec3{0, 0, 0}
		}
	}

	var unitDirection = r.Direction().Normalize()
	var t float32 = float32(0.5) * (unitDirection.Y() + 1)
	return mgl32.Vec3{1, 1, 1}.Mul((1 - t)).Add(mgl32.Vec3{0.5, 0.7, 1.0}.Mul(t))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func randInUnitSpehre() mgl32.Vec3 {
	var p mgl32.Vec3 = mgl32.Vec3{rand.Float32(), rand.Float32(), rand.Float32()}.Mul(2.0).Sub(mgl32.Vec3{1, 1, 1})

	for ok := true; ok; ok = p.LenSqr() < 1.0 {
		p = mgl32.Vec3{rand.Float32(), rand.Float32(), rand.Float32()}.Mul(2.0).Sub(mgl32.Vec3{1, 1, 1})
	}

	return p
}

func main() {
	const nx int = 600
	const ny int = 300
	const ns int = 100

	var camera Camera = Camera{
		origin:          mgl32.Vec3{0, 0, 0},
		lowerLeftCorner: mgl32.Vec3{-2, -1, -1},
		horizontal:      mgl32.Vec3{4, 0, 0},
		vertical:        mgl32.Vec3{0, 2, 0},
	}

	var contents string = ""
	contents += "P3"
	contents += fmt.Sprintln(nx, ny)
	contents += fmt.Sprintln(255)

	var list []Hitable
	list = append(list, Sphere{center: mgl32.Vec3{0, 0, -1}, radius: 0.5, material: Lambertian{albedo: mgl32.Vec3{0.8, 0.3, 0.3}}})
	list = append(list, Sphere{center: mgl32.Vec3{0, -100.5, -1}, radius: 100, material: Lambertian{albedo: mgl32.Vec3{0.8, 0.8, 0.0}}})
	list = append(list, Sphere{center: mgl32.Vec3{1, 0, -1}, radius: 0.5, material: Metal{albedo: mgl32.Vec3{0.8, 0.6, 0.2}, fuzz: 0.3}})
	list = append(list, Sphere{center: mgl32.Vec3{-1, 0, -1}, radius: 0.5, material: Metal{albedo: mgl32.Vec3{0.8, 0.8, 0.8}, fuzz: 1.0}})

	var world HitableList = HitableList{list: list}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			var col = mgl32.Vec3{0, 0, 0}
			for s := 0; s < ns; s++ {
				var u = (float32(i) + rand.Float32()) / float32(nx)
				var v = (float32(j) + rand.Float32()) / float32(ny)

				var r = camera.getRay(u, v)
				// var p mgl32.Vec3 = r.P(2.0)
				col = col.Add(color(r, world, 0))
			}

			col = col.Mul(1 / float32(ns))
			col = mgl32.Vec3{float32(math.Sqrt(float64(col.X()))), float32(math.Sqrt(float64(col.Y()))), float32(math.Sqrt(float64(col.Z())))}
			var ir = int(255.99 * col[0])
			var ig = int(255.99 * col[1])
			var ib = int(255.99 * col[2])

			contents += fmt.Sprintln(ir, ig, ib)
		}
	}

	err := ioutil.WriteFile("dump.ppm", []byte(contents), 0644)
	check(err)

	fmt.Println("Finished generating image")
}
