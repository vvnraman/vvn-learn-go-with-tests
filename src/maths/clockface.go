package clockface

import (
	"math"
	"time"
)

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

// A Point represents a 2D cartesian co-ordinate
type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30) // won't work

	// because this is what we're doing in the tests
	return (math.Pi / (30 / (float64(t.Second()))))

}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

// SecondHand is the unit vector of the second hand on an analog clock at time
// `t` represented as a Point
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}
