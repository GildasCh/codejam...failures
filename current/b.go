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
	ret := make([]rune, N)

	totals := map[rune]int{
		'R': R,
		'O': O,
		'Y': Y,
		'G': G,
		'B': B,
		'V': V,
	}

	prims := [][]int{
		[]int{0, R + O + V},
		[]int{1, Y + O + G},
		[]int{2, B + G + V}}
	sort.Sort(ByN(prims))

	for _, p := range prims {
		curr := '0'
		switch p[0] {
		case 0: // R
			curr = 'R'
		case 1: // Y
			curr = 'Y'
		case 2: // B
			curr = 'B'
		}

		// we need to place p[1] unicorns
		for k := 0; k < N; k++ {
			i := spot(k, N)
			fmt.Println("spot", k, N, "=", i)

			if p[1] <= 0 {
				break
			}

			left := ret[(i-1+N)%N]
			right := ret[(i+1)%N]

			if conflicts(curr, left) || conflicts(curr, right) {
				continue
			}

			if ret[i] == rune(0) {
				ret[i] = curr
				p[1]--
				totals[curr]--
			} else {
				inplace := ret[i]
				compo := compose(curr, inplace)
				if totals[compo] > 0 {
					ret[i] = compo
					p[1]--
					totals[compo]++
					totals[inplace]--
				}
			}
		}

		if p[1] > 0 {
			fmt.Println("IMPOSSIBLE:" + string(ret))
			return "IMPOSSIBLE"
		}
	}

	return string(ret)
}

func spot(k int, N int) int {
	if k == 0 {
		return 0
	}

	prev := prevPowerOfTwo(k)
	flag := N / (2 * prev)
	if flag == 0 {
		return 1 + spot(prev/2+(k-prev), N)
	}
	return (1 + 2*(k-prev)) * flag
}

func prevPowerOfTwo(v int) int {
	ret := 1
	for v>>1 > 0 {
		v = v >> 1
		ret *= 2
	}
	return ret
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
func (a ByN) Less(i, j int) bool { return a[i][1] < a[j][1] }

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
