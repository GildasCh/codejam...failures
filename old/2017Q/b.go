package main

import (
	"fmt"
	"os"
	"strings"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	N []rune
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
		c.N = []rune(readString())
		C = append(C, c)
	}

	for i, c := range C {
		s := solve(c.N)
		fmt.Fprintf(output, "Case #%d: %s\n", i+1, s)
	}
}

func solve(N []rune) string {
	out := make([]rune, len(N), len(N))

	i := 0

	fmt.Println("First part: tidy")
	for {
		out[i] = N[i]
		if i == len(N)-1 {
			return string(out)
		}

		current := int(N[i] - '0')
		next := int(N[i+1] - '0')
		if current > next {
			break
		}

		i++
	}

	fmt.Println("Second part: fallback:", i, string(out))
	for {
		current := int(N[i] - '0')
		current--
		out[i] = rune('0' + current)

		if i <= 0 {
			break
		}

		previous := int(out[i-1] - '0')
		if previous <= current {
			break
		}
		i--
	}

	i++

	fmt.Println("Third part: append nines:", i, string(out))
	for k := i; k < len(N); k++ {
		out[k] = '9'
	}

	fmt.Println(out)
	fmt.Println(strings.TrimLeft(string(out), "0"))
	return strings.TrimLeft(string(out), "0")
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
