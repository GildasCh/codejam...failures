package main

import (
	"fmt"
	"os"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	R, C, H, V int
	Chips      [][]bool
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.R, c.C, c.H, c.V = readInt(), readInt(), readInt(), readInt()
		c.Chips = make([][]bool, c.R)
		for i := 0; i < c.R; i++ {
			c.Chips[i] = make([]bool, c.C)
			for j := 0; j < c.C; j++ {
				c.Chips[i][j] = readChar() == '@'
			}
			readChar()
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		fmt.Fprintf(os.Stderr, "Case #%v\n", c)
		res := solve(c.R, c.C, c.H, c.V, c.Chips)
		fmt.Printf("Case #%d: %s\n", i+1, res)
	}
}

func solve(R, C, H, V int, Chips [][]bool) string {
	if solveRec(R, C, H, V, Chips) {
		return "POSSIBLE"
	}
	return "IMPOSSIBLE"
}

func solveRec(R, C, H, V int, Chips [][]bool) bool {
	fmt.Fprintf(os.Stderr, "solveRec %v %v %v %v\n", R, C, H, V)

	if H == 0 && V == 0 {
		return true
	}

	NChips := count(Chips, 0, 0, R, C)

	if NChips == 0 {
		return true
	}

	divs := (H + 1) * (V + 1)
	if NChips%divs != 0 || NChips%(H+1) != 0 || NChips%(V+1) != 0 {
		return false
	}
	target := NChips / divs
	hTarget := NChips / (H + 1)
	vTarget := NChips / (V + 1)
	fmt.Fprintf(os.Stderr, "target: %d\n", target)

	// H divs
	h1, h2 := 0, 0
	hCuts := []int{0}
	for len(hCuts) < H+1 {
		for count(Chips, h1, 0, h2, C) < hTarget {
			h2++
		}

		if count(Chips, h1, 0, h2, C) != hTarget {
			return false
		}

		hCuts = append(hCuts, h2)
		h1 = h2
	}

	// V divs
	v1, v2 := 0, 0
	vCuts := []int{0}
	for len(vCuts) < V+1 {
		for count(Chips, 0, v1, R, v2) < vTarget {
			v2++
		}

		if count(Chips, 0, v1, R, v2) != vTarget {
			return false
		}

		vCuts = append(vCuts, v2)
		v1 = v2
	}

	fmt.Fprintf(os.Stderr, "cutting at %v,%v\n", hCuts, vCuts)

	// Verify solution
	for i := 0; i < H; i++ {
		for j := 0; j < V; j++ {
			if count(Chips, hCuts[i], vCuts[j], hCuts[i+1], vCuts[j+1]) != target {
				return false
			}
		}
	}

	return true
}

func count(Chips [][]bool, r1, c1, r2, c2 int) int {
	NChips := 0
	for i := r1; i < r2; i++ {
		for j := c1; j < c2; j++ {
			if Chips[i][j] {
				NChips++
			}
		}
	}

	// fmt.Fprintf(os.Stderr, "cut(%d,%d,%d,%d) = %d\n", r1, c1, r2, c2, NChips)
	return NChips
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
