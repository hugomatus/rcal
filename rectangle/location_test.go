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

//TestLocation_4
//: use case #4
func TestLocation_4(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_4.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	rectangles[0].Describe()

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location4]

	want := []Point{{
		X: 8,
		Y: 10,
	}, {
		X: 10,
		Y: 10,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_5
//: use case #5
func TestLocation_5(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_5.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location5]

	want := []Point{{
		X: 10,
		Y: 10,
	}, {
		X: 14,
		Y: 8,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_6
//: use case #6
func TestLocation_6(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_6.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location6]

	want := []Point{{
		X: 14,
		Y: 6,
	}, {
		X: 14,
		Y: 8,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_7
//: use case #7
func TestLocation_7(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_7.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location7]

	want := []Point{{
		X: 10,
		Y: 4,
	}, {
		X: 14,
		Y: 6,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_8
//: use case #8
func TestLocation_8(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_8.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location8]

	want := []Point{{
		X: 10,
		Y: 4,
	}, {
		X: 12,
		Y: 4,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_9
//: use case #9
func TestLocation_9(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_9.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[0].Locate(rectangles[1])

	got := location[Location9]

	want := []Point{{
		X: 6,
		Y: 8,
	}, {
		X: 14,
		Y: 8,
	}, {
		X: 6,
		Y: 6,
	}, {
		X: 14,
		Y: 6,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_10
//: use case #10
func TestLocation_10(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_9.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[1].Locate(rectangles[0])

	got := location[Location10]

	want := []Point{{
		X: 6,
		Y: 8,
	}, {
		X: 14,
		Y: 8,
	}, {
		X: 6,
		Y: 6,
	}, {
		X: 14,
		Y: 6,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_11
//use case #11
//TODO : converse of #2
func TestLocation_11(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_2.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[1].Locate(rectangles[0])

	got := location[Location11]

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

//TestLocation_12
//: use case #12
//TODO converse #4
func TestLocation_12(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_4.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[1].Locate(rectangles[0])

	got := location[Location12]

	want := []Point{{
		X: 8,
		Y: 10,
	}, {
		X: 10,
		Y: 10,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_13
//: use case #13
//TODO converse #6
func TestLocation_13(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_6.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[1].Locate(rectangles[0])

	got := location[Location13]

	want := []Point{{
		X: 14,
		Y: 6,
	}, {
		X: 14,
		Y: 8,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

//TestLocation_14
//: use case #14
//TODO Converse #8
func TestLocation_14(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_8.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	location := rectangles[1].Locate(rectangles[0])

	got := location[Location14]

	want := []Point{{
		X: 10,
		Y: 4,
	}, {
		X: 12,
		Y: 4,
	}}

	fmt.Printf("Points of Intersection: %v\n", got)
	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}
