package main

import (
	"fmt"
	"os"
	"sort"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	N int
	R int
	O int
	Y int
	G int
	B int
	V int
}

func main() {
	sample := os.Args[1]
	fileIn := sample + ".in"
	fileOut := sample + ".out"

	var err error
	input, err = os.Open(fileIn)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", fileIn, err))
	}
	output, err = os.Create(fileOut)
	if err != nil {
		panic(fmt.Sprintf("creating %s: %v", fileOut, err))
	}
	defer input.Close()
	defer output.Close()

	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N = readInt()
		c.R = readInt()
		c.O = readInt()
		c.Y = readInt()
		c.G = readInt()
		c.B = readInt()
		c.V = readInt()
		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		res := solve(c.N, c.R, c.O, c.Y, c.G, c.B, c.V)
		fmt.Fprintf(output, "Case #%d: %s\n", i+1, res)
	}
}

func solve(N, R, O, Y, G, B, V int) string {
	colors := []*Color{
		&Color{'R', R},
		&Color{'Y', Y},
		&Color{'B', B},
	}

	sort.Sort(ByN(colors))

	// Impossible
	if colors[0].n > colors[1].n+colors[2].n {
		return "IMPOSSIBLE"
	}

	// Possible
	ret := []rune{}
	veryFirst = colors[0].letter
	var previous *Color
	for i := 0; i < N; i++ {
		place := colors[0]
		if place == previous {
			place = colors[1]
		}
		ret = append(ret, place.letter)
		previous = place
		place.n--
		sort.Sort(ByN(colors))
	}

	return string(ret)
}

var veryFirst rune

type Color struct {
	letter rune
	n      int
}

type ByN []*Color

func (a ByN) Len() int      { return len(a) }
func (a ByN) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByN) Less(i, j int) bool {
	return a[i].n > a[j].n || (a[i].n == a[j].n && a[i].letter == veryFirst)
}

func readInt() int {
	var i int
	fmt.Fscanf(input, "%d", &i)
	return i
}

func readString() string {
	var str string
	fmt.Fscanf(input, "%s", &str)
	return str
}

func readFloat() float64 {
	var x float64
	fmt.Fscanf(input, "%f", &x)
	return x
}

func readRune() rune {
	var x rune
	fmt.Fscanf(input, "%c", &x)
	return x
}
