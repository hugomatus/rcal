package rectangle

import (
	"fmt"
	"math"
)

type Location int

type PointsOfIntersection map[Location][]Point

const (
	Location1 Location = iota + 1
	Location2
	Location3
)

type inLocation func(r1 Rectangle, r2 Rectangle) []Point

var blockOfLocation []inLocation

func init() {
	blockOfLocation = []inLocation{inLocation1, inLocation2, inLocation3}
}

func (r1 Rectangle) doLocate(r2 Rectangle) PointsOfIntersection {

	intersection := PointsOfIntersection{}

	for location, f := range blockOfLocation {
		location += 1
		fmt.Printf("scanning for intersection(s) @ block: %v\n", location)

		tmp_ := f(r1, r2)

		if len(tmp_) > 0 {
			intersection[Location(location)] = tmp_
			return intersection
		}
	}

	return intersection
}

// Locate identify on the assumption that the rectangles intersect.
func (r1 Rectangle) Locate(r2 Rectangle) PointsOfIntersection {

	return r1.doLocate(r2)
}

/*		+-----------+
		|        R1 |
+-----------+       |
| R2        |       |
|           |-------+
+-----------+
*/

// inLocation1 is R2 left and Location8 of R1
func inLocation1(r1 Rectangle, r2 Rectangle) []Point {

	if r2.xMin() < r1.xMin() {
		if r2.xMax() > r1.xMin() && r2.xMax() < r1.xMax() {
			if r2.yMax() < r1.yMax() {
				if r2.yMin() < r1.yMin() {
					LoadFigure(1)

					points := r1.Location1(r2)
					return points
				}
			}
		}
	}

	return nil
}

/*		+-----------+
		|        R1 |
+-----------+       |
| R2		|       |
+-----------+       |
		|           |
		+-----------+
*/

//inLocation2 is R2 left and partially within of R1
func inLocation2(r1 Rectangle, r2 Rectangle) []Point {

	if r2.xMin() < r1.xMin() {
		if r2.xMax() > r1.xMin() && r2.xMax() < r1.xMax() {
			if r2.yMax() < r1.yMax() {
				if r2.yMin() > r1.yMin() {
					LoadFigure(2)
					return r1.Location2(r2)
				}
			}
		}
	}
	return nil
}

/*
+-----------+
| R2        |
|           |-------+
|           |    R1 |
+-----------+       |
		|           |
		+-----------+
*/

// inLocation3
func inLocation3(r1 Rectangle, r2 Rectangle) []Point {

	if r2.xMin() < r1.xMin() {
		if r2.xMax() > r1.xMin() && r2.xMax() < r1.xMax() {
			if r2.yMax() > r1.yMax() {
				if r2.yMin() > r1.yMin() {
					LoadFigure(3)
					return r1.PointsOfIntersection(r2)
				}
			}
		}
	}

	return nil
}

// Location1 #1 rule for identifying points of intersection in this block
func (r Rectangle) Location1(r2 Rectangle) []Point {

	var points []Point

	points = append(points, Point{
		X: int(math.Max(float64(r.xMin()), float64(r2.xMin()))),
		Y: int(math.Min(float64(r.yMax()), float64(r2.yMax()))),
	},
		Point{
			X: int(math.Min(float64(r.xMax()), float64(r2.xMax()))),
			Y: int(math.Max(float64(r.yMin()), float64(r2.yMin()))),
		})
	return points
}

// Location2 #2 rule for identifying points of intersection in this block
func (r Rectangle) Location2(r2 Rectangle) []Point {

	var points []Point

	points = append(points, Point{
		X: int(math.Max(float64(r.xMin()), float64(r2.xMin()))),
		Y: int(math.Min(float64(r.yMax()), float64(r2.yMax()))),
	},
		Point{
			X: int(math.Max(float64(r.xMin()), float64(r2.xMin()))),
			Y: int(math.Max(float64(r.yMin()), float64(r2.yMin()))),
		})

	return points
}

// PointsOfIntersection Location rule for identifying points of intersection in this block
func (r Rectangle) PointsOfIntersection(r2 Rectangle) []Point {

	var points []Point

	points = append(points, Point{
		X: int(math.Max(float64(r.xMin()), float64(r2.xMin()))),
		Y: int(math.Max(float64(r.yMin()), float64(r2.yMin()))),
	},
		Point{
			X: int(math.Min(float64(r.xMax()), float64(r2.xMax()))),
			Y: int(math.Min(float64(r.yMax()), float64(r2.yMax()))),
		})

	return points

}
