package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
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
		fmt.Println("Case #%d", i+1)
		res := solve(c.N, c.R, c.O, c.Y, c.G, c.B, c.V)
		fmt.Fprintf(output, "Case #%d: %s\n", i+1, res)
	}
}

func solve(N, R, O, Y, G, B, V int) string {
	colors := []*Color{
		&Color{'R', R},
		&Color{'O', O},
		&Color{'Y', Y},
		&Color{'G', G},
		&Color{'B', B},
		&Color{'V', V},
	}

	couples := []*ColorCouple{
		&ColorCouple{"VRO", V + R + O},
		&ColorCouple{"OYG", O + Y + G},
		&ColorCouple{"YGB", Y + G + B},
		&ColorCouple{"ROY", R + O + Y},
		&ColorCouple{"GBV", G + B + V},
		&ColorCouple{"BVR", B + V + R},
	}

	sort.Sort(ByN(colors))
	sort.Sort(ByNc(couples))

	// Impossible
	if couples[0].n > N/2 {
		return "IMPOSSIBLE"
	}

	// Possible
	ret := []rune{}
	// veryFirst = bestColor(colors, couples, 'X')
	previous := &Color{}
	for i := 0; i < N; i++ {
		place := bestColor(colors, couples, previous.letter)
		if place == nil {
			fmt.Println("No best color found:", string(ret))
			return "FAIL"
		}
		ret = append(ret, place.letter)
		previous = place
		place.n--
		for _, cc := range couples {
			if strings.Contains(cc.def, string(previous.letter)) {
				cc.n--
			}
		}
		sort.Sort(ByN(colors))
		sort.Sort(ByNc(couples))
	}

	return string(ret)
}

// colors and couples need to be sorted
func bestColor(colors []*Color, couples []*ColorCouple, previous rune) *Color {
	var bestCouple string
	for _, cc := range couples {
		fmt.Println("Trying", cc.def, ", previous:", string(previous))
		conflict := false
		for _, r := range []rune(cc.def) {
			if conflicts(r, previous) {
				fmt.Println("Conflit found:", string(r))
				conflict = true
			}
		}
		if !conflict {
			bestCouple = cc.def
			break
		}
	}

	for _, c := range colors {
		if strings.Contains(bestCouple, string(c.letter)) {
			fmt.Println("bestCouple:", bestCouple)
			fmt.Println("bestColor:", string(c.letter))
			return c
		}
	}

	return nil
}

func conflicts(a, b rune) bool {
	if a == b {
		return true
	}
	switch a {
	case 'R':
		if b == 'O' || b == 'V' {
			return true
		}
	case 'Y':
		if b == 'O' || b == 'G' {
			return true
		}
	case 'B':
		if b == 'G' || b == 'V' {
			return true
		}
	case 'O':
		if b == 'R' || b == 'Y' || b == 'G' || b == 'V' {
			return true
		}
	case 'G':
		if b == 'Y' || b == 'B' || b == 'V' || b == 'O' {
			return true
		}
	case 'V':
		if b == 'R' || b == 'G' || b == 'B' || b == 'O' {
			return true
		}
	}
	return false
}

// var veryFirst rune

type Color struct {
	letter rune
	n      int
}

type ByN []*Color

func (a ByN) Len() int           { return len(a) }
func (a ByN) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByN) Less(i, j int) bool { return a[i].n > a[j].n }

type ColorCouple struct {
	def string
	n   int
}

type ByNc []*ColorCouple

func (a ByNc) Len() int           { return len(a) }
func (a ByNc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNc) Less(i, j int) bool { return a[i].n > a[j].n }

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
