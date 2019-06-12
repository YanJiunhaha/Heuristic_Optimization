package object

import "math"

func Result(X float64, Y float64) float64 {
	x := math.Pow(X+2, 2)
	y := math.Pow(Y+2, 2)
	_x := math.Pow(X-2, 2)
	_y := math.Pow(Y-2, 2)

	return 4/(_x+_y+1) + 3/(_x+y+1) + 2/(x+_y+1)
}
