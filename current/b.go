package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	N int
	P int
	R []int
	Q [][]int
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
		c.P = readInt()

		c.R = make([]int, c.N)
		for i := 0; i < c.N; i++ {
			c.R[i] = readInt()
		}

		c.Q = make([][]int, c.N)
		for i := 0; i < c.N; i++ {
			c.Q[i] = make([]int, c.P)
			for j := 0; j < c.P; j++ {
				c.Q[i][j] = readInt()
			}
		}

		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		s := solve(c.N, c.P, c.R, c.Q)
		fmt.Fprintf(output, "Case #%d: %d\n", i+1, s)
	}
}

type Range struct {
	l, u int
}

func solve(N int, P int, R []int, Q [][]int) int {
	// fmt.Println(N, P, R, Q)

	// Ranges
	ranges := make([][]Range, N)
	for i := 0; i < N; i++ {
		ranges[i] = getRanges(R[i], Q[i])
	}

	// fmt.Println(ranges)

	counter := 0

	for {
		if oneIsEmpty(ranges) {
			return counter
		}

		firsts := make([]Range, N)
		for i := 0; i < N; i++ {
			firsts[i] = ranges[i][0]
		}

		works, toRemove := analyse(firsts)
		if works {
			counter++
			// Remove all the firsts
			for i := 0; i < len(ranges); i++ {
				ranges[i] = ranges[i][1:]
			}
		} else {
			ranges[toRemove] = ranges[toRemove][1:]
		}
	}

	return -1
}

func analyse(R []Range) (bool, int) {
	highestLower, smallestUpper := R[0].l, R[0].u

	for _, r := range R {
		if highestLower < r.l {
			highestLower = r.l
		}
		if smallestUpper > r.u {
			smallestUpper = r.u
		}
	}

	if highestLower <= smallestUpper {
		// fmt.Println(R, "Working!")
		return true, -1
	}

	// Remove the highestLower
	for i, r := range R {
		if r.l == highestLower {
			return false, i
		}
	}

	panic("No")
	return false, -1
}

func oneIsEmpty(Q [][]Range) bool {
	for _, q := range Q {
		if len(q) == 0 {
			return true
		}
	}
	return false
}

func getRanges(r int, q []int) []Range {
	var ret []Range

	for i := 0; i < len(q); i++ {
		ran := getRange(r, q[i])
		if ran.l != -1 {
			ret = append(ret, ran)
		}
	}

	// Order
	sort.Sort(ByUpper(ret))

	return ret
}

type ByUpper []Range

func (a ByUpper) Len() int           { return len(a) }
func (a ByUpper) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpper) Less(i, j int) bool { return a[i].u > a[j].u }

func getRange(r int, q int) Range {
	upper := int(float64(q) / (float64(r) * 0.9))
	lower := int(math.Ceil(float64(q) / (float64(r) * 1.1)))
	if lower > upper {
		return Range{-1, -1}
	}
	return Range{lower, upper}
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
