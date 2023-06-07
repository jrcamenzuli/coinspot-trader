package common

type Point struct {
	X, Y float64
}

func AverageSlope(points []Point) float64 {
	sumSlope := 0.0
	n := len(points)
	for i := 0; i < n-1; i++ {
		dx := points[i+1].X - points[i].X
		dy := points[i+1].Y - points[i].Y
		slope := dy / dx
		sumSlope += slope
	}
	return sumSlope / float64(n-1)
}
