package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

var input *os.File
var output *os.File

var T int64
var C []Case

type Case struct {
	N int64
	K int64
	R []int64
	H []int64
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

	for i := int64(0); i < T; i++ {
		c := Case{}
		c.N = readInt()
		c.K = readInt()
		c.R = make([]int64, c.N)
		c.H = make([]int64, c.N)
		for i := int64(0); i < c.N; i++ {
			c.R[i] = readInt()
			c.H[i] = readInt()
		}
		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		ret := solve(c.N, c.K, c.R, c.H)
		fmt.Fprintf(output, "Case #%d: %.9f\n", i+1, ret)
	}
}

func solve(N, K int64, R, H []int64) float64 {
	// fmt.Println(R, C, L)
	p := createPank(R, H)
	sort.Sort(p)

	pByR := ByR(createPank(R, H))
	sort.Sort(pByR)

	max := float64(0)
	// Select best largest pancake
	for i := int64(0); i <= N-K; i++ {
		largest := i
		count := 1
		exposed :=
			math.Pi*float64(pByR.R[largest])*float64(pByR.R[largest]) +
				2*math.Pi*float64(pByR.R[largest])*float64(pByR.H[largest])

		// Select all the others
		for j := 0; int64(count) < K && j < len(R); j++ {
			if p.R[j] > pByR.R[largest] || p.I[j] == pByR.I[largest] {
				continue
			}
			count++
			exposed += 2 * math.Pi * float64(p.R[j]) * float64(p.H[j])
		}

		if int64(count) < K {
			fmt.Println("ERROR!")
		}

		if exposed > max {
			max = exposed
		}
	}

	return max
}

type Pank struct {
	I    []int
	R, H []int64
}

func createPank(R, H []int64) Pank {
	ret := Pank{
		make([]int, len(R)),
		make([]int64, len(R)),
		make([]int64, len(R)),
	}
	for i := 0; i < len(R); i++ {
		ret.I[i] = i
		ret.R[i] = R[i]
		ret.H[i] = H[i]
	}
	return ret
}

func (a Pank) Len() int { return len(a.R) }
func (a Pank) Swap(i, j int) {
	a.I[i], a.I[j] = a.I[j], a.I[i]
	a.R[i], a.R[j] = a.R[j], a.R[i]
	a.H[i], a.H[j] = a.H[j], a.H[i]
}
func (a Pank) Less(i, j int) bool {
	return float64(a.H[i])*float64(a.R[i]) > float64(a.H[j])*float64(a.R[j])
}

type ByR Pank

func (a ByR) Len() int { return len(a.R) }
func (a ByR) Swap(i, j int) {
	a.I[i], a.I[j] = a.I[j], a.I[i]
	a.R[i], a.R[j] = a.R[j], a.R[i]
	a.H[i], a.H[j] = a.H[j], a.H[i]
}
func (a ByR) Less(i, j int) bool { return a.R[i] > a.R[j] }

func readInt() int64 {
	var i int64
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
