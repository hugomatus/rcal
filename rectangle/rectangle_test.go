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
