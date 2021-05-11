package rectangle

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//TestNew
func TestNew(t *testing.T) {

	r1 := New(Point{12, 14}, Point{4, 10})

	r1.Describe()

	got := []Point{r1.TopLeft, r1.BottomRight}

	want := []Point{{
		X: 4,
		Y: 14,
	}, {
		X: 12,
		Y: 10,
	}}

	if !assert.Equal(t, want, got) {
		t.Fail()
	}

}

//TestLoad
func TestLoad(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_1.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	if err != nil {
		t.Fail()
	}

	if len(rectangles) != 2 {
		t.Fail()
	}
}

func TestLoadFigure(t *testing.T) {

	LoadFigure(1)
}

func TestRectangle_Containment_1(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_1.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	got := rectangles[0].Containment(rectangles[1])

	want := false

	if !assert.Equal(t, want, got) {
		t.Fail()
	}

}

func TestRectangle_Containment_2(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_2.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	got := rectangles[0].Containment(rectangles[1])

	want := false

	if !assert.Equal(t, want, got) {
		t.Fail()
	}

}

func TestRectangle_Containment_4(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_4.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	got := rectangles[0].Containment(rectangles[1])

	want := false

	if !assert.Equal(t, want, got) {
		t.Fail()
	}

}
func TestRectangle_Containment_15(t *testing.T) {

	file, err := os.Open("test_data/rectangle_data_15.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	got := rectangles[0].Containment(rectangles[1])

	want := true

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_18(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_18.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := Proper

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_19(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_18.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[1].Adjacency(rectangles[0])

	got := adjacencyType

	want := Proper

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_20(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_20.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := Partial

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_21(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_21.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := Partial

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_22(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_22.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := Partial

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_23(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_23.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := Partial

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_24(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_24.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := SubLine

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_25(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_25.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := SubLine

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_26(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_26.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := SubLine

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_27(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_27.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := SubLine

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_28(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_15.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := NoAjacency

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestRectangle_Adjacency_29(t *testing.T) {
	file, err := os.Open("test_data/rectangle_data_2.json")

	if err != nil {
		t.Fail()
	}

	rectangles, err := Load(file)

	adjacencyType := rectangles[0].Adjacency(rectangles[1])

	got := adjacencyType

	want := NoAjacency

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}
