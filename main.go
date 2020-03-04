package main

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

func hitSphere(center mgl32.Vec3, radius float32, r Ray) float32 {
	var oc mgl32.Vec3 = r.Origin().Sub(center)

	var a float32 = r.Direction().Dot(r.Direction())
	var b float32 = 2 * oc.Dot(r.Direction())
	var c float32 = oc.Dot(oc) - radius*radius

	var discriminant float32 = b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - float32(math.Sqrt(float64(discriminant)))) / (2 * a)
	}
}

func color(r Ray) mgl32.Vec3 {
	var t float32 = hitSphere(mgl32.Vec3{0, 0, -1}, 0.5, r)

	if t > 0 {
		var N mgl32.Vec3 = r.P(t).Sub(mgl32.Vec3{0, 0, -1}).Normalize()
		return mgl32.Vec3{N.X() + 1, N.Y() + 1, N.Z() + 1}.Mul(0.5)
	}

	var unitDirection = r.Direction().Normalize()
	t = float32(0.5) * (unitDirection.Y() + 1)
	return mgl32.Vec3{1, 1, 1}.Mul((1 - t)).Add(mgl32.Vec3{0.5, 0.7, 1.0}.Mul(t))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const nx int = 600
	const ny int = 300

	var lowerLeftCorner = mgl32.Vec3{-2, -1, -1}
	var horizontal = mgl32.Vec3{4, 0, 0}
	var vertical = mgl32.Vec3{0, 2, 0}
	var origin = mgl32.Vec3{0, 0, 0}

	var contents string = ""
	contents += "P3"
	contents += fmt.Sprintln(nx, ny)
	contents += fmt.Sprintln(255)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			var u = float32(i) / float32(nx)
			var v = float32(j) / float32(ny)

			var r = Ray{origin, lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v))}
			var col = color(r)

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
