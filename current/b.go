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
	// fmt.Println(R, C, L)
	prims := [][]int{
		[]int{int('R'), R + O + V},
		[]int{int('Y'), Y + O + G},
		[]int{int('B'), B + G + V}}
	sort.Sort(ByN(prims))

	if prims[0][1] > N/2 {
		return "IMPOSSIBLE"
	}

	ret := make([]rune, N)

	// Place most frequent
	i := 0
	for prims[0][1] > 0 {
		ret[i] = rune(prims[0][0])
		prims[0][1]--
		i += 2
	}

	// Place others
	swit := false
	for i := N-1; i >=0; i-- {
		if ret[i] != rune(0) {
			continue
		}
		if swit || prims[2][1] <= 0 {
			ret[i] = rune(prims[1][0])
			prims[1][1]--
			swit = false
		} else {
			ret[i] = rune(prims[2][0])
			prims[2][1]--
			swit = true
		}
	}

	if prims[1][1] > 0 || prims[2][1] > 0 {
		fmt.Println("Problem:", prims[1][1], prims[2][1], string(ret), N, R, O, Y, G, B, V)
	}

	return string(ret)
}

func compose(a, b rune) rune {
	switch a {
	case 'R':
		switch b {
		case 'Y':
			return 'O'
		case 'B':
			return 'V'
		}
	case 'Y':
		switch b {
		case 'R':
			return 'O'
		case 'B':
			return 'G'
		}
	case 'B':
		switch b {
		case 'R':
			return 'V'
		case 'Y':
			return 'G'
		}
	}

	return rune(0)
}

func conflicts(a, b rune) bool {
	if a == rune(0) || b == rune(0) {
		return false
	}

	if a == b {
		return true
	}

	switch b {
	case 'O':
		return a == 'R' || a == 'Y'
	case 'G':
		return a == 'B' || a == 'Y'
	case 'V':
		return a == 'R' || a == 'B'
	}

	return false
}

type ByN [][]int

func (a ByN) Len() int           { return len(a) }
func (a ByN) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByN) Less(i, j int) bool { return a[i][1] > a[j][1] }

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
