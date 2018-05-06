package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var input io.Reader

func main() {
	input = bufio.NewReader(os.Stdin)

	T := readInt()
	// fmt.Fprintf(os.Stderr, "%d test cases\n", T)

	for i := 0; i < T; i++ {
		flavors = make(map[int]int)
		sold = make(map[int]bool)

		N := readInt()
		// fmt.Fprintf(os.Stderr, "%d customers\n", N)

		for j := 0; j < N; j++ {
			F := readInt()
			// fmt.Fprintf(os.Stderr, "%d flavors\n", F)

			var fs []int
			for k := 0; k < F; k++ {
				ff := readInt()
				Seen(ff)
				fs = append(fs, ff)
			}

			fmt.Println(Choose(fs))
			// fmt.Fprintf(os.Stderr, "Choose(%v) => %d", fs, Choose(fs))
		}
	}
}

var flavors map[int]int
var sold map[int]bool

func Seen(i int) {
	flavors[i]++
}

func Choose(fs []int) int {
	best := -1
	bestSeen := -1
	for _, i := range fs {
		if sold[i] {
			continue
		}

		if bestSeen == -1 || flavors[i] < bestSeen {
			best = i
			bestSeen = flavors[i]
		}
	}
	sold[best] = true
	return best
}

func readInt() int {
	var i int
	fmt.Fscanf(input, "%d ", &i)
	return i
}
