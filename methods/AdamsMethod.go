package methods

//
//import (
//	"DifferentialEqns/utils"
//	"math"
//)
//
//func AdamsMethod() {
//
//}
//
//func rungeKutaForAdams(eqn utils.Equation, x0, y0, xo, xn, h, accuracy float64) utils.XY {
//	var newY float64
//	f := eqn.F
//	percise := PreciseAns(eqn, x0, y0, xn, h)
//	n := (xn - x0) / h
//	for {
//		runge := countRungeKutta(f, x0, x0+(xn-x0)/n*3, y0, h)
//		n := len(runge.X)
//		for i := 0; i < n-3; i++ {
//			f1 := f(runge.X[n-1], runge.Y[n-1])
//			f2 := f(runge.X[n-2], runge.Y[n-2])
//			f3 := f(runge.X[n-3], runge.Y[n-3])
//			f4 := f(runge.X[n-4], runge.Y[n-4])
//			dF1 := f1 - f2
//			dF2 := f1 - 2*f2 + f3
//			dF3 := f1 - 3*f2 + 3*f3 - f4
//			newY = runge.Y[n-1] + h*f1 + h*h*dF1/2 + 5/12*math.Pow(h, 3)*dF2 + 3/8*math.Pow(h, 4)*dF3
//		}
//		runge.X = append(runge.X)
//		percise = PreciseAns(eqn, runge.X[n-1]+h, y0, xn, h)
//		if math.Abs(percise.X[n-1]-newY) < math.Pow(10, accuracy) {
//			break
//		}
//
//		h = h / 2
//	}
//
//	return utils.XY{
//		X: X,
//		Y: Y,
//	}
//
//}
