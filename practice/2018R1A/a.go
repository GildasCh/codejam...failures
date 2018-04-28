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

	NChips := count(Chips, R, C)

	if NChips == 0 {
		return true
	}

	divs := (H + 1) * (V + 1)
	if NChips%divs != 0 {
		return false
	}
	target := NChips / divs
	fmt.Fprintf(os.Stderr, "target: %d\n", target)

	h, v := 1, 1
	for {
		if count(Chips, h, C) > target*(H+1) ||
			count(Chips, R, v) > target*(V+1) ||
			count(Chips, h, v) > target {
			return false
		}

		if count(Chips, h, C) < target*(H+1) {
			h++
			continue
		}
		if count(Chips, R, v) < target*(V+1) {
			v++
			continue
		}

		break
	}

	fmt.Fprintf(os.Stderr,
		"cutting at %d,%d; counts %d, %d, %d\n",
		h, v, count(Chips, h, v), count(Chips, h, C), count(Chips, R, v),
	)

	if count(Chips, h, v) != target {
		return false
	}

	Chips = Chips[h:]
	for i := 0; i < R-h; i++ {
		Chips[i] = Chips[i][v:]
	}

	return solveRec(R-h, C-v, H-1, V-1, Chips)
}

func count(Chips [][]bool, r, c int) int {
	NChips := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if Chips[i][j] {
				NChips++
			}
		}
	}
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
