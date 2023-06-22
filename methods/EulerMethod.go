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
	yh2uncut := countEuler(f, x0, xn, y0, h/2, yh.X[len(yh.X)-1]+1e-9)
	yh2 := Cutout(origin, yh2uncut)
	for i := 0; i < n; i++ {
		if cnt >= 100 {
			break
		}
		for index, _ := range yh.Y {
			if math.Abs(yh.Y[index]-yh2.Y[index])/15 > math.Pow(10, -accuracy) {
				cnt++
				return EulerMethod(eqn, x0, y0, xo, xn, h/2, accuracy, cnt, origin, false)
			}
		}
		break
	}
	println("EULER ITERATIONS: ", cnt)
	return yh2uncut
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
	// O(n*n) gains reliability and simplicity
	if len(origin.X) != len(origin.Y) {
		fmt.Println("X and Y len don't match in the origin, function Cutout")
		os.Exit(1)
	}
	if len(new.X) != len(new.Y) {
		fmt.Println("X and Y len don't match in the new struct, function Cutout")
		os.Exit(1)
	}
	filteredX := []float64{}
	filteredY := []float64{}
	for _, looked := range origin.X {
		closest := math.MaxFloat64
		currentY := math.MaxFloat64
		for lookingIndex, looking := range new.X {
			currentCloseness := math.Abs(looked - looking)
			if currentCloseness < closest {
				closest = currentCloseness
				currentY = new.Y[lookingIndex]
			}
		}
		filteredX = append(filteredX, looked)
		filteredY = append(filteredY, currentY)
	}

	if len(filteredY) != len(origin.Y) {
		fmt.Println("Cutout function messed up, length don't match up")
		if false {
			fmt.Println("Origin.X:")
			fmt.Println(origin.X)
			fmt.Println("Origin.Y:")
			fmt.Println(origin.Y)
			fmt.Println("new.X")
			fmt.Println(new.X)
			fmt.Println("filteredY:")
			fmt.Println(filteredY)
		}
		os.Exit(1)
	}
	return utils.XY{X: filteredX, Y: filteredY}
}
