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
	W []int
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N = readInt()
		for j := 0; j < c.N; j++ {
			c.W = append(c.W, readInt())
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...\n")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res := solve(c.D, c.P)
		fmt.Printf("Case #%d: %d\n", i+1, res)
	}
}

func solve(N int, W []int) int {
	fmt.Fprintf(os.Stderr, "solving %d %v\n", N, W)

	var ahs AntHeaps
	for _, w := range W {
		for _, aa := range ahs {
			ahs = append(ahs, AntHeap{
				w:    append(aa.w, w),
				sumW: aa.sumW + w,
				l:    aa.l + 1,
			})
		}

		clean(&ahs)
	}

	return max
}

type AntHeap struct {
	w    []int
	sumW int
	l    int
}

type AntHeaps []AntHeap

func clean(ahs []AntHeap) {
	sort.Sort(ByWeight(ahs))

	longest := -1
	for _, aa := range ahs {
		if aa.l < longest {

		}
	}
}

type ByWeight []AntHeap

func (i ByWeight) Len() int           { return len(i.a) }
func (i ByWeight) Swap(a, b int)      { i.a[a], i.a[b] = i.a[b], i.a[a] }
func (i ByWeight) Less(a, b int) bool { return i.a[a].sumW < i.a[b].sumW }

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
