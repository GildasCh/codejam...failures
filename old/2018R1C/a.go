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
	N int
	L int
	W []string
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N = readInt()
		c.L = readInt()
		for j := 0; j < c.N; j++ {
			c.W = append(c.W, readString())
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...\n")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res := solve(c.N, c.L, c.W)
		fmt.Printf("Case #%d: %s\n", i+1, res)
	}
}

func solve(N int, L int, W []string) string {
	fmt.Fprintf(os.Stderr, "solving %d %v\n", N, W)

	for i := 0; i < L; i++ {
		letters := make(map[rune]int)

		for j := 0; j < N; j++ {
			letters[rune(W[j][i])] = letters[rune(W[j][i])] + 1
		}

		// if len(letters)*len(letters) < N {
		// 	continue
		// }

		rarest := '-'
		rCount := -1
		for l, n := range letters {
			if rCount == -1 || rCount > n {
				rarest = l
				rCount = n
			}
		}

		for j := 0; j < N; j++ {
			res := []rune(W[j])
			res[i] = rarest
			if check(W, string(res)) {
				return string(res)
			}
		}
	}

	return "-"
}

func check(W []string, res string) bool {
	for _, w := range W {
		if w == res {
			return false
		}
	}
	return true
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

func readRune() rune {
	var x rune
	fmt.Fscanf(input, "%c", &x)
	return x
}
