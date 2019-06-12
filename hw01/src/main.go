package main

import (
    "fmt"
    my "../../object_function"
)

func main() {
	var bx, by float64
	var f float64
	f = -1e99

	const scale float64 = 1e-6
	for x := 1.993; x <= 1.994; x += scale {
		for y := 1.989; y <= 1.990; y += scale {
			if my.Result(x, y) > f {
				f = my.Result(x, y)
				bx = x
				by = y
			}
		}
	}

	fmt.Println("The biggest value for object function :")
	fmt.Printf("F(%f, %f) = %f\n", bx, by, f)
}
