package rectangle

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRectangle_PointsOfIntersection_Location_1(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_1.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Intersection(rectangles[1])

	got := location[Location1]

	want := []Point{{
		X: 6,
		Y: 6,
	}, {
		X: 7,
		Y: 4,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}

}

//TestLocation_1
//: use case #1
func TestLocation_1(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_1.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Intersection(rectangles[1])

	got := location[Location1]

	want := []Point{{
		X: 6,
		Y: 6,
	}, {
		X: 7,
		Y: 4,
	}}

	fmt.Printf("Total Points of Intersection: %d\n", len(got))
	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}

}

//TestLocation_2
//: use case #2
func TestLocation_2(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_2.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location2]

	want := []Point{{
		X: 6,
		Y: 8,
	}, {
		X: 6,
		Y: 6,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_3
//: use case #3
func TestLocation_3(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_3.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location3]

	want := []Point{{
		X: 6,
		Y: 8,
	}, {
		X: 8,
		Y: 10,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}
