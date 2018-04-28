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
	N, P int
	W    []int
	H    []int
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N, c.P = readInt(), readInt()
		c.W, c.H = make([]int, c.N), make([]int, c.N)
		for i := 0; i < c.N; i++ {
			c.H[i] = readInt()
			c.W[i] = readInt()
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		fmt.Fprintf(os.Stderr, "Case #%v\n", c)
		res := solve(c.N, c.P, c.H, c.W)
		fmt.Printf("Case #%d: %s\n", i+1, res)
	}
}

func solve(N, P int, H, W []int) string {
	// Starting point
	start := 0
	for i := 0; i < N; i++ {
		start += 2*H[i] + 2*W[i]
	}

	ivs := &Intervals{}
	ivs.Add(float64(start), float64(start))

	// Add more intervals for each potential cut
	for i := 0; i < N; i++ {
		ivs.AddToAll(
			minCut(H[i], W[i]),
			maxCut(H[i], W[i]),
		)
	}

	fmt.Fprintf(os.Stderr, "All possible intervals: %v\n", ivs.a)

	return "IMPOSSIBLE"
}

func minCut(a, b int) float64 {
	if a < b {
		return 2 * float64(a)
	}
	return 2 * float64(b)
}

func maxCut(a, b int) float64 {
	return 2 * math.Sqrt(float64(a*a+b*b))
}

type Intervals struct {
	a []Interval
}

type Interval struct {
	L, H float64
}

func (i *Intervals) Add(l, h float64) {
	i.a = append(i.a, Interval{l, h})

	i.SortAndMerge()
}

func (i *Intervals) AddToAll(l, h float64) {
	toAdd := []Interval{}
	for k := 0; k < len(i.a); k++ {
		toAdd = append(toAdd, Interval{
			L: i.a[k].L + l,
			H: i.a[k].H + h,
		})
	}

	for _, ii := range toAdd {
		i.a = append(i.a, ii)
	}

	i.SortAndMerge()
}

func (i *Intervals) SortAndMerge() {
	sort.Slice(i.a, func(a, b int) bool {
		return i.a[a].L < i.a[b].L
	})

	for k := 0; k < len(i.a)-1; k++ {
		if i.a[k].H >= i.a[k+1].L {
			i.a[k].H = i.a[k+1].H
			i.a = append(i.a[:k+1], i.a[k+2:]...)
			k--
		}
	}
}

func readInt() int {
	var i int
	fmt.Fscanf(os.Stdin, "%d", &i)
	return i
}

func readString() string {
	var str string
	fmt.Fscanf(os.Stdin, "%s", &str)
	return str
}

func readChar() rune {
	var c rune
	fmt.Fscanf(os.Stdin, "%c", &c)
	return c
}
