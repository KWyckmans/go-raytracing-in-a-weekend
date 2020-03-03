package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-gl/mathgl/mgl32"
)

func color(r Ray) mgl32.Vec3 {
	var unitDirection = r.Direction().Normalize()
	var t float32 = float32(0.5) * (unitDirection.Y() + 1)
	return mgl32.Vec3{1, 1, 1}.Mul((1 - t)).Add(mgl32.Vec3{0.5, 0.7, 1.0}.Mul(t))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const nx int = 200
	const ny int = 100

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

	var A = mgl32.Vec3{0, 0, 0}
	var B = mgl32.Vec3{1, 1, 1}
	var r = Ray{A, B}

	fmt.Println(r.A)
}
