package main

import (
	"DifferentialEqns/IO"
	"DifferentialEqns/methods"
	"DifferentialEqns/utils"
	"math"
)

// Одношаговые методы:
// 1. Метод Эйлера,
// 3. Метод Рунге-Кутта 4- го порядка.
// Многошаговые методы:
// 4. Адамса
func main() {
	var eqn1 = utils.Equation{
		F:              func(x, y float64) float64 { return y + (1+x)*y*y },
		RealF:          func(x, C float64) float64 { return -(math.Exp(x) / (x*math.Exp(x) + C)) },
		C:              func(x, y float64) float64 { return y + math.Exp(x)/(x*math.Exp(x)) },
		NameOfFunction: "y' = y + (1 + x) * y^2",
	}
	var eqn2 = utils.Equation{
		F:              func(x, y float64) float64 { return -(2*y + 1) * math.Cos(x) },
		RealF:          func(x, C float64) float64 { return C/math.Exp(2*math.Sin(x)) - 0.5 },
		C:              func(x, y float64) float64 { return (y + 0.5) * math.Exp(2*math.Sin(x)) },
		NameOfFunction: "y' = -(2y + 1) * cos(x)",
	}
	var eqn3 = utils.Equation{
		F:              func(x, y float64) float64 { return math.Pow(x, 5) },
		RealF:          func(x, C float64) float64 { return math.Pow(x, 6)/6 + C },
		C:              func(x, y float64) float64 { return (y*6 - math.Pow(x, 6)) / 6 },
		NameOfFunction: "y' = x^5",
	}
	var fToBeEvaluated utils.Equation
	x0, y0, xo, xn, h, accuracy, eqn := IO.UserInput(eqn1, eqn2, eqn3)
	switch eqn {
	case 1:
		fToBeEvaluated = eqn1
	case 2:
		fToBeEvaluated = eqn2
	case 3:
		fToBeEvaluated = eqn3
	}
	//fmt.Println(x0, y0, xo, xn, h, accuracy)

	euler := methods.EulerMethod(fToBeEvaluated, x0, y0, xo, xn, h, accuracy, 0, utils.XY{}, true)
	println("euler ready")
	runge := methods.RungeKuttaMethod(fToBeEvaluated, x0, y0, xo, xn, h, accuracy, utils.XY{}, true)
	println("runge ready")
	adams := methods.AdamsMethod(fToBeEvaluated, x0, y0, xo, xn, h, accuracy, 0, utils.XY{}, true)
	println("adams ready")
	precise := methods.PreciseAns(fToBeEvaluated, x0, y0, xn, h)
	IO.OutputResults(euler, runge, adams, precise, accuracy)
}
