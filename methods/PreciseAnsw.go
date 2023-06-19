package methods

import "DifferentialEqns/utils"

func PreciseAns(eqn utils.Equation, x0, y0, xn, h float64) utils.XY {
	rf := eqn.RealF
	c := eqn.C(x0, y0)
	X := []float64{}
	Y := []float64{}
	x := x0
	y := y0
	for x <= xn+h {
		X = append(X, x)
		Y = append(Y, y)
		x += h
		y = rf(x, c)
	}
	return utils.XY{
		X: X,
		Y: Y,
	}
}
