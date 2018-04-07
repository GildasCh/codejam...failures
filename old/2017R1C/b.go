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
	Ac, Aj int
	C, D   []int
	J, K   []int
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
		c.Ac = readInt()
		c.Aj = readInt()
		c.C = make([]int, c.Ac)
		c.D = make([]int, c.Ac)
		for i := 0; i < c.Ac; i++ {
			c.C[i] = readInt()
			c.D[i] = readInt()
		}
		c.J = make([]int, c.Aj)
		c.K = make([]int, c.Aj)
		for i := 0; i < c.Aj; i++ {
			c.J[i] = readInt()
			c.K[i] = readInt()
		}
		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		fmt.Printf("Case #%d\n", i+1)
		res := solve(c.Ac, c.Aj, c.C, c.D, c.J, c.K)
		fmt.Fprintf(output, "Case #%d: %d\n", i+1, res)
	}
}

func solve(Ac, Aj int, C, D, J, K []int) int {
	// fmt.Println(R, C, L)

	// Generate acts
	acts := Acts{
		make([]bool, Ac+Aj),
		make([]int, Ac+Aj),
		make([]int, Ac+Aj),
		make([]bool, Ac+Aj),
	}
	for i := 0; i < Ac; i++ {
		acts.cam[i] = true
		acts.start[i] = C[i]
		acts.end[i] = D[i]
		acts.visited[i] = false
	}
	for i := 0; i < Aj; i++ {
		acts.cam[Ac+i] = false
		acts.start[Ac+i] = J[i]
		acts.end[Ac+i] = K[i]
		acts.visited[i] = false
	}
	sort.Sort(acts)
	// fmt.Println(acts)

	// Calculating low and high bounds for Cam
	low := 0
	high := 0
	freeCam := []int{}
	freeJam := []int{}
	lastWasCam := acts.cam[Ac+Aj-1]
	inversions := 0
	for i := 0; i < Ac+Aj; i++ {
		if acts.visited[i] {
			continue
		}
		acts.visited[i] = true

		// Cam's
		if acts.cam[i] {
			if !lastWasCam {
				inversions++
			}

			lowBack := acts.start[i]
			lowForward := acts.end[i]
			highBack := acts.end[i]
			highForward := acts.start[i]

			// Go as far back as possible
			for j := (i - 1 + Ac + Aj) % (Ac + Aj); ; j = (j - 1 + Ac + Aj) % (Ac + Aj) {
				if !acts.cam[j] {
					highBack = acts.end[j]
					break
				}
				freeCam = append(freeCam, (lowBack-acts.end[j]+1440)%1440)
				lowBack = acts.start[j]
				if acts.visited[j] {
					break
				}
				acts.visited[j] = true
			}

			// Go as far forward as possible
			for j := (i - 1 + Ac + Aj) % (Ac + Aj); ; j = (j - 1 + Ac + Aj) % (Ac + Aj) {
				if !acts.cam[j] {
					highForward = acts.start[j]
					break
				}
				freeCam = append(freeCam, (acts.start[j]-lowForward+1440)%1440)
				lowForward = acts.end[j]
				if acts.visited[j] {
					break
				}
				acts.visited[j] = true
			}

			fmt.Println(lowBack, lowForward, highBack, highForward)
			fmt.Println("Low+=", (lowForward-lowBack+1440)%1440)
			low += (lowForward - lowBack + 1440) % 1440
			fmt.Println("high +=", (highForward-highBack+1440)%1440)
			fmt.Println(highForward, highBack)
			high += (highForward - highBack + 1440) % 1440

			lastWasCam = true
			continue
		}

		// Jamie's
		if lastWasCam {
			inversions++
		}
		lowBack := acts.start[i]
		lowForward := acts.end[i]
		// Go as far back as possible
		for j := (i - 1 + Ac + Aj) % (Ac + Aj); j != i; j = (j - 1 + Ac + Aj) % (Ac + Aj) {
			if acts.cam[j] {
				break
			}
			freeJam = append(freeJam, (lowBack-acts.end[j]+1440)%1440)
			lowBack = acts.start[j]
			acts.visited[j] = true
		}

		// Go as far forward as possible
		for j := (i - 1 + Ac + Aj) % (Ac + Aj); j != i; j = (j - 1 + Ac + Aj) % (Ac + Aj) {
			if acts.cam[j] {
				break
			}
			freeJam = append(freeJam, (acts.start[j]-lowForward+1440)%1440)
			lowForward = acts.end[j]
			acts.visited[j] = true
		}

		lastWasCam = false
	}

	fmt.Println(low, high, freeCam, freeJam)

	if high < 720 {
		// Cam need to fit between Jamie's
		sort.Sort(ByF(freeJam))
		for _, fj := range freeJam {
			high += fj
			inversions += 2
			if high >= 720 {
				return inversions
			}
		}
	} else if low > 720 {
		// Jamie need to fit between Cam's
		sort.Sort(ByF(freeCam))
		for _, fj := range freeCam {
			fmt.Println("low -", fj)
			low -= fj
			fmt.Println("low =", low)
			inversions += 2
			if low <= 720 {
				return inversions
			}
		}
	}

	// OK
	return inversions
}

type ByF []int

func (a ByF) Len() int { return len(a) }
func (a ByF) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByF) Less(i, j int) bool { return a[i] > a[j] }

type Plage struct {
	low, high int
}

type Acts struct {
	cam     []bool
	start   []int
	end     []int
	visited []bool
}

func (a Acts) Len() int { return len(a.cam) }
func (a Acts) Swap(i, j int) {
	a.cam[i], a.cam[j] = a.cam[j], a.cam[i]
	a.start[i], a.start[j] = a.start[j], a.start[i]
	a.end[i], a.end[j] = a.end[j], a.end[i]
}
func (a Acts) Less(i, j int) bool { return a.start[i] < a.start[j] }

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
