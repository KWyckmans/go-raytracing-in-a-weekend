package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

func color(r Ray, world Hitable, depth int) mgl32.Vec3 {
	var rec HitRecord
	if world.hit(r, 0.001, math.MaxFloat32, &rec) {
		var scattered Ray
		var attenuation mgl32.Vec3

		if depth < 50 && rec.material.Scatter(r, rec, &attenuation, &scattered) {
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

func randomWorld() []Hitable {
	const n int = 500

	var list []Hitable
	list = append(list, Sphere{center: mgl32.Vec3{0, -1000, 0}, radius: 1000, material: Lambertian{albedo: mgl32.Vec3{0.5, 0.5, 0.5}}})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			var chooseMat float32 = rand.Float32()
			var center mgl32.Vec3 = mgl32.Vec3{float32(a) + float32(0.9)*rand.Float32(), 0.2, float32(b) + float32(0.9)*rand.Float32()}

			if (center.Sub(mgl32.Vec3{4, 0.2, 0}).Len() > 0.9) {
				if chooseMat < 0.8 {
					list = append(list, Sphere{center: center, radius: 0.2, material: Lambertian{mgl32.Vec3{rand.Float32() * rand.Float32(), rand.Float32() * rand.Float32(), rand.Float32() * rand.Float32()}}})
				} else if chooseMat < 0.95 {
					list = append(list, Sphere{center: center, radius: 0.2, material: Metal{albedo: mgl32.Vec3{float32(0.5) * (rand.Float32() + 1.0), float32(0.5) * (rand.Float32() + 1.0), float32(0.5) * (rand.Float32() + 1.0)}, fuzz: float32(0.5) * rand.Float32()}})
				} else {
					list = append(list, Sphere{center: center, radius: 0.2, material: Dielectric{refIdx: 1.5}})
				}
			}
		}
	}

	list = append(list, Sphere{center: mgl32.Vec3{0, 1, 0}, radius: 1.0, material: Dielectric{refIdx: 1.5}})
	list = append(list, Sphere{center: mgl32.Vec3{-4, 1, 0}, radius: 1.0, material: Lambertian{mgl32.Vec3{0.4, 0.2, 0.1}}})
	list = append(list, Sphere{center: mgl32.Vec3{4, 1, 0}, radius: 1.0, material: Metal{mgl32.Vec3{0.7, 0.6, 0.5}, 0.0}})

	return list
}

func main() {
	const nx int = 500
	const ny int = 400
	const ns int = 100

	var lookFrom = mgl32.Vec3{3, 3, 2}
	var lookAt = mgl32.Vec3{0, 0, -1}
	var distToFocus float32 = (lookFrom.Sub(lookAt)).Len()
	var aperture float32 = 2.0

	var camera = BuildCamera(lookFrom, lookAt, mgl32.Vec3{0, 1, 0}, 20, float32(nx)/float32(ny), aperture, distToFocus)

	var contents string = ""
	contents += "P3"
	contents += fmt.Sprintln(nx, ny)
	contents += fmt.Sprintln(255)

	var world HitableList = HitableList{list: randomWorld()}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			var col = mgl32.Vec3{0, 0, 0}
			for s := 0; s < ns; s++ {
				var u = (float32(i) + rand.Float32()) / float32(nx)
				var v = (float32(j) + rand.Float32()) / float32(ny)

				var r = camera.GetRay(u, v)
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

	filename := time.Now().Format("20060201-1504") + ".ppm"
	err := ioutil.WriteFile(filename, []byte(contents), 0644)
	check(err)

	fmt.Println("Finished generating image")
}
