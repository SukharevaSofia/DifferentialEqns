package methods

import (
	"DifferentialEqns/utils"
	"math"
)

func RungeKuttaMethod(eqn utils.Equation, x0, y0, xo, xn, h, accuracy float64, origin utils.XY, start bool) utils.XY {
	var cnt = 0
	f := eqn.F
	yh := countRungeKutta(f, x0, xn, y0, h)
	if start {
		origin = yh
	}
	yh = Cutout(origin, yh)
	yh2 := Cutout(origin, countRungeKutta(f, x0, xn, y0, h/2))
	for i := 0; i < len(yh.Y); i++ {
		if cnt >= 100 {
			break
		}
		if math.Abs((yh.Y[i]-yh2.Y[i])/15) > math.Pow(10, accuracy) {
			return RungeKuttaMethod(eqn, x0, y0, xo, xn, h/2, accuracy, origin, false)
		}
		cnt++
	}
	println("RUNGE ITERATIONS: ", cnt)
	return yh2
}

func countRungeKutta(f func(x float64, y float64) float64, x0, xn, y0, h float64) utils.XY {
	X := []float64{}
	Y := []float64{}
	x := x0
	y := y0
	for x <= xn+h {
		X = append(X, x)
		Y = append(Y, y)
		k1 := h * f(x, y)
		k2 := h * f(x+h/2, y+k1/2)
		k3 := h * f(x+h/2, y+k2/2)
		k4 := h * f(x+h, y+k3)
		x += h
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
	}
	return utils.XY{
		X: X,
		Y: Y,
	}
}
