package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

var input io.Reader

func main() {
	input = bufio.NewReader(os.Stdin)

	n := readInt()
	fmt.Fprintf(os.Stderr, "%d test cases\n", n)

	for i := 0; i < n; i++ {
		a := readInt()
		fmt.Fprintf(os.Stderr, "test case %d: a=%d\n", i, a)

		h := int(math.Ceil(math.Sqrt(float64(a))))
		w := int(math.Ceil(math.Sqrt(float64(a))))

		x, y := 2, 2
		for {
			fmt.Printf("%d %d\n", x, y)

			i, j := readInt(), readInt()
			fmt.Fprintf(os.Stderr, "read %d %d\n", i, j)

			if i == 0 && j == 0 {
				fmt.Fprintf(os.Stderr, "OK :)\n")
				break
			}
			if i == -1 && j == -1 {
				fmt.Fprintf(os.Stderr, "failed :(\n")
				break
			}

			x++
			if x > w-1 {
				x = 2
				y++
				if y > h-1 {
					x, y = 2, 2
				}
			}
		}
	}
}

func readInt() int {
	var i int
	fmt.Fscanf(input, "%d\n", &i)
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
