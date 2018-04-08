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
	V []int
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N = readInt()
		c.V = make([]int, c.N)
		for i := 0; i < c.N; i++ {
			c.V[i] = readInt()
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res, ok := solve(c.N, c.V)
		if ok {
			fmt.Printf("Case #%d: OK\n", i+1)
		} else {
			fmt.Printf("Case #%d: %d\n", i+1, res)
		}
	}
}

func solve(D int, V []int) (int, bool) {
	// fmt.Fprintf(os.Stderr, "solving %d %v\n", D, V)

	V = trouble(V)
	// fmt.Fprintf(os.Stderr, "after trouble %v\n", V)

	for i := 0; i < len(V)-1; i++ {
		if V[i] > V[i+1] {
			return i, false
		}
	}

	return 0, true
}

func trouble(V []int) []int {
	sort.Sort(OddSort(V))
	sort.Sort(EvenSort(V))
	return V
}

type OddSort []int
type EvenSort []int

func (a OddSort) Len() int {
	return len(a) / 2
}
func oddIndex(i int) int {
	return 1 + 2*i
}
func (a OddSort) Swap(i, j int) {
	i, j = oddIndex(i), oddIndex(j)
	a[i], a[j] = a[j], a[i]
}
func (a OddSort) Less(i, j int) bool {
	i, j = oddIndex(i), oddIndex(j)
	return a[i] < a[j]
}

func (a EvenSort) Len() int {
	if len(a)%2 == 0 {
		return len(a) / 2
	}
	return len(a)/2 + 1
}
func evenIndex(i int) int {
	return 2 * i
}
func (a EvenSort) Swap(i, j int) {
	i, j = evenIndex(i), evenIndex(j)
	a[i], a[j] = a[j], a[i]
}
func (a EvenSort) Less(i, j int) bool {
	i, j = evenIndex(i), evenIndex(j)
	return a[i] < a[j]
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
