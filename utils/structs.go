package utils

type Equation struct {
	F              func(x, y float64) float64
	RealF          func(x, C float64) float64
	C              func(x, y float64) float64
	NameOfFunction string
}

type XY struct {
	X []float64
	Y []float64
}
