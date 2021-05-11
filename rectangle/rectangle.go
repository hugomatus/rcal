// Package rectangle represents Rectangle Object
package rectangle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"text/tabwriter"
)

//Point holds X and Y coordinates which represent a corner of a rectangle
type Point struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

//Rectangle has for corners defined by four points
type Rectangle struct {
	TopLeft     Point
	TopRight    Point
	BottomLeft  Point
	BottomRight Point
}

//xMax returns max X coordinate
func (r Rectangle) xMax() int {
	return r.TopRight.X
}

//xMin returns min Y coordinate
func (r Rectangle) xMin() int {
	return r.BottomLeft.X
}

//yMax returns max Y coordinate
func (r Rectangle) yMax() int {
	return r.TopRight.Y
}

//yMin returns min Y coordinate
func (r Rectangle) yMin() int {
	return r.BottomLeft.Y
}

// isValid represent a valid rectangle in the context of this problem set
// assumes axis aligned rectangles
// are normalized, they have only horizontal and vertical sides.
func (r Rectangle) isValid() bool {

	if r.TopLeft.X != r.BottomLeft.X {
		return false
	}

	if r.TopRight.X != r.BottomRight.X {
		return false
	}

	if r.TopLeft.Y != r.TopRight.Y {
		return false
	}

	if r.BottomLeft.Y != r.BottomRight.Y {
		return false
	}

	return true
}

// New create and returns a pointer to a Rectangle
func New(TopRight Point, BottomLeft Point) *Rectangle {
	r := new(Rectangle)
	r.TopLeft = Point{
		X: BottomLeft.X,
		Y: TopRight.Y,
	}
	r.TopRight = TopRight
	r.BottomRight = Point{
		X: TopRight.X,
		Y: BottomLeft.Y,
	}
	r.BottomLeft = BottomLeft

	valid := r.isValid()

	if !valid {
		return nil
	}

	return r
}

// Describe describes a Rectangle
func (r Rectangle) Describe() {

	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t\t%s\t%s\t%s\t%s\t%s\t%s\t", "TopLeft", "TopRight", "BottomLeft", "BottomRight", "xMin", "xMax", "yMin", "yMax")
	fmt.Fprintf(w, "\n %s\t%s\t\t%s\t%s\t%s\t%s\t%s\t%s\t", "--------", "--------", "--------", "--------", "--------", "--------", "--------", "--------")
	fmt.Fprintf(w, "\n %v\t%v\t\t%v\t%v\t%v\t%v\t%v\t%v\t\n", r.TopLeft, r.TopRight, r.BottomLeft, r.BottomRight, r.xMin(), r.xMax(), r.yMin(), r.yMax())
}

func (r Rectangle) toJson() (string, error) {

	data, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

//Load rectangles r1 and r2 from file for feature analysis
func Load(file *os.File) ([]Rectangle, error) {

	var results []Rectangle

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &results)

	for _, r := range results {

		r.TopLeft = Point{
			X: r.BottomLeft.X,
			Y: r.TopRight.Y,
		}
		r.BottomRight = Point{
			X: r.TopRight.X,
			Y: r.BottomLeft.Y,
		}

		r.Describe()
	}

	return results, nil
}

//LoadFigure loads rectangle diagram from file
func LoadFigure(figure int) error {

	fig := "rectangle/figures/fig" + strconv.Itoa(figure) + ".txt"
	file, err := os.Open(fig)

	if err != nil {
		fmt.Println(err)
		return err
	}

	data, _ := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("\n+--------Reference--------+\n%v \n", string(data))

	return nil
}

//Intersection confirms if rectangles r1 and r2 intersect
func (r1 Rectangle) Intersection(r2 Rectangle) PointsOfIntersection {

	if r1.xIntersect(r2) && r1.yIntersect(r2) {

		points := r1.Locate(r2)

		for location, point := range points {
			if len(point) == 0 {
				delete(points, location)
			}

		}

		return points
	}

	return nil
}

func (r1 Rectangle) xIntersect(r2 Rectangle) bool {

	if r1.xMin() < r2.xMax() && r1.xMax() > r2.xMin() {
		return true
	}
	return false
}

//yIntersect
func (r1 Rectangle) yIntersect(r2 Rectangle) bool {

	return r1.yMin() < r2.yMax() && r1.yMax() > r2.yMin()
}

/*
+---------------+
|R1  +------+   |
|	 |	R2 	|	|
|	 +------+	|
+---------------+
*/

// Containment determines if Rectangle R1 contains Rectangle r2
func (r1 Rectangle) Containment(r2 Rectangle) bool {

	if len(r1.Intersection(r2)) == 0 {
		if r2.xMin() > r1.xMin() && r2.xMax() < r1.xMax() {
			if r2.yMax() < r1.yMax() {
				if r2.yMin() > r1.yMin() {
					LoadFigure(15)
					return true
				}
			}
		}
	}
	return false
}

// Adjacency determines if Rectangle R1 and R2 are adjacent either
// Proper, Partial or sub-line
func (r1 Rectangle) Adjacency(r2 Rectangle) bool {

	return r1.isAdjacent(r2)
}
