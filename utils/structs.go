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

func XYtoMap(xy XY, precision float64) map[float64]float64 {
	xymap := map[float64]float64{}
	for i := 0; i < len(xy.X); i++ {
		xymap[Truncate(xy.X[i], precision)] = xy.Y[i]
	}
	return xymap
}
