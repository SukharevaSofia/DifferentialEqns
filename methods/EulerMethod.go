package methods

import (
	"DifferentialEqns/utils"
	"fmt"
	"math"
	"os"
)

func EulerMethod(eqn utils.Equation, x0, y0, xo, xn, h, accuracy float64, cnt int, origin utils.XY, starting bool) utils.XY {
	n := int(math.Floor((xn-xo)/h) + 1)
	f := eqn.F
	yh := countEuler(f, x0, xn, y0, h, xn+h)
	if starting {
		origin = yh
	}
	yh = Cutout(origin, yh)
	yh2 := Cutout(origin, countEuler(f, x0, xn, y0, h/2, yh.X[len(yh.X)-1]+1e-9))
	for i := 0; i < n; i++ {
		if cnt >= 100 {
			break
		}
		fmt.Println("ITER# ", cnt)
		fmt.Println(" yh: ", yh.Y)
		fmt.Println(" yh2:; ", yh2.Y)
		for index, _ := range yh.Y {
			if math.Abs(yh.Y[index]-yh2.Y[index])/15 > math.Pow(10, -accuracy) {
				println("E ITERATION №", cnt)
				fmt.Println(yh.Y[1] - yh2.Y[1])
				cnt++
				return EulerMethod(eqn, x0, y0, xo, xn, h/2, accuracy, cnt, origin, false)
			}
		}
		break
	}
	println("EULER ITERATIONS: ", cnt)
	return yh2
}

func countEuler(f func(x float64, y float64) float64, x0, xn, y0, h, goal float64) utils.XY {
	X := []float64{}
	Y := []float64{}
	x := x0
	y := y0
	for x <= goal {
		X = append(X, x)
		Y = append(Y, y)
		y += h * f(x, y)
		x += h
	}
	return utils.XY{
		X: X,
		Y: Y,
	}
}

func Cutout(origin, new utils.XY) utils.XY {
	if len(origin.X) != len(origin.Y) {
		fmt.Println("X and Y len don't match in the origin, function Cutout")
		os.Exit(1)
	}
	if len(new.X) != len(new.Y) {
		fmt.Println("X and Y len don't match in the new struct, function Cutout")
		os.Exit(1)
	}
	filteredY := []float64{}
	i := 0
	j := 0
	for i < len(origin.X) && j < len(new.X) {
		if utils.Float64Equals(origin.X[i], new.X[j], 1e-9) {
			filteredY = append(filteredY, new.Y[j])
			i++
		} else if new.X[j] > origin.X[i] {
			i++
		} else if origin.X[i] > new.X[j] {
			j++
		}
	}
	if len(filteredY) != len(origin.Y) {
		fmt.Println("Cutout function messed up, length don't match up")
		os.Exit(1)
	}
	return utils.XY{X: origin.X, Y: filteredY}
}
