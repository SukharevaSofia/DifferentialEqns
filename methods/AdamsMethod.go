package methods

import (
	"DifferentialEqns/utils"
	"fmt"
	"math"
)

func AdamsMethod(eqn utils.Equation, x0, y0, xo, xn, h, accuracy float64, cnt int, origin utils.XY, starting bool) utils.XY {
	precise := PreciseAns(eqn, x0, y0, xn, h)
	f := eqn.F
	i := 0
	x := x0
	y := y0
	X := []float64{}
	Y := []float64{}
	var xyNew = utils.XY{
		X: X,
		Y: Y,
	}
	xy1 := countRungeKutta(f, x0, xn, y0, h)
	if starting {
		origin = precise
	}
	xy1 = Cutout(origin, xy1)
	precise = Cutout(origin, precise)
	for i != 4 {
		xyNew.X = append(xyNew.X, xy1.X[i])
		xyNew.Y = append(xyNew.Y, xy1.Y[i])
		i++
	}
	i-- // to reset i back to the last element, from 4 to 3

	x = origin.X[3]
	y = xyNew.Y[i]
	for x <= xn {
		f1 := f(xyNew.X[len(xyNew.X)-1], xyNew.Y[len(xyNew.Y)-1])
		f2 := f(xyNew.X[len(xyNew.X)-2], xyNew.Y[len(xyNew.Y)-2])
		f3 := f(xyNew.X[len(xyNew.X)-3], xyNew.Y[len(xyNew.Y)-3])
		f4 := f(xyNew.X[len(xyNew.X)-4], xyNew.Y[len(xyNew.Y)-4])

		d1 := f1 - f2
		d2 := f1 - 2*f2 + f3
		d3 := f1 - 3*f2 + 3*f3 - f4
		y += h*f1 + h*h*d1/2 + 5*h*h*h*d2/12 + 3*h*h*h*h*d3/8
		x += h
		xyNew.X = append(xyNew.X, x)
		xyNew.Y = append(xyNew.Y, y)
		i++
	}

	result := xyNew
	xyNew = Cutout(origin, xyNew)
	for index, _ := range xyNew.Y {
		precisionDiff := math.Pow(10, -accuracy)
		if math.Abs(xyNew.Y[index]-precise.Y[index]) > precisionDiff {
			cnt++
			if utils.DEBUG_STATEMENTS {
				fmt.Println("IT: ", cnt, "  IND: ", index, " : ", math.Abs(xyNew.Y[index]-precise.Y[index])-precisionDiff)
				fmt.Println(xyNew.Y[index], precise.Y[index])
			}
			return AdamsMethod(eqn, x0, y0, xo, xn, h/2, accuracy, cnt, origin, false)
		}
	}
	fmt.Println("ADAMS ITERATIONS: ", cnt)
	return result
}
