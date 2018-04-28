package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	data := map[int]Interval{
		0: Interval{-192, -191},
		1: Interval{192, 193},
		2: Interval{10, 50},
		3: Interval{49, 70},
		4: Interval{-1, 1},
		5: Interval{12239, 33322},
		6: Interval{1234, 13259},
	}

	actual := &Intervals{}

	for _, d := range data {
		actual.Add(d.L, d.H)
	}

	expected := &Intervals{
		a: []Interval{
			Interval{L: -192, H: -191},
			Interval{L: -1, H: 1},
			Interval{L: 10, H: 70},
			Interval{L: 192, H: 193},
			Interval{L: 1234, H: 33322}}}

	assert.Equal(t, expected, actual)
}

func TestAddToAll(t *testing.T) {
	actual := &Intervals{
		a: []Interval{
			Interval{L: -192, H: -191},
			Interval{L: -1, H: 1},
			Interval{L: 10, H: 70},
			Interval{L: 192, H: 193},
			Interval{L: 1234, H: 33322}}}

	actual.AddToAll(10, 20)

	expected := &Intervals{
		a: []Interval{
			Interval{L: -192, H: -191},
			Interval{L: -182, H: -171},
			Interval{L: -1, H: 1},
			Interval{L: 9, H: 90},
			Interval{L: 192, H: 193},
			Interval{L: 202, H: 213},
			Interval{L: 1234, H: 33342}}}

	assert.Equal(t, expected, actual)
}
