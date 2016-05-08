package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
	"strings"
	"sort"
	// "math"
)

type InputData struct {
	N int
	A [][]int
}

// Helpers
func read(input string) (int, []InputData) {
	ret := []InputData{}

	file, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	// fmt.Println("Total:", total)
	counter := 1
    for counter <= total {
		data := InputData{0,[][]int{}}
        // fmt.Println(scanner.Text())
		scanner.Scan()
		challenge := strings.Split(scanner.Text(), " ")
		data.N, _ = strconv.Atoi(challenge[0])

		for i := 0; i < 2*data.N - 1; i++ {
			scanner.Scan()
			readRow := strings.Split(scanner.Text(), " ")
			row := []int{}
			for j := 0; j < data.N; j++ {
				a, _ := strconv.Atoi(readRow[j])
				row = append(row, a)
			}
			data.A = append(data.A, row)
		}

		ret = append(ret, data)
		counter += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return total, ret
}

func main() {
	total, inputData := read(os.Args[1])

	i := 0
    for i < total {
		res := calculate(inputData[i])
		fmt.Print("Case #", (i+1), ": ", res)
		fmt.Println()
		i += 1
    }
}

func calculate(in InputData) string {
	// sorted := sort(in.A)
	sort.Sort(RowList(in.A))

	fmt.Println(in.A)
	chosen := [][]int{in.A[0]}
	all := elect(chosen, in.A, in.N, 1)

	fmt.Println("Size of all:", len(all))
	fmt.Println(all)
	for i := 0; i < len(all); i++ {
		try := all[i]
		ok, res := verify(try, in.A, in.N)
		if ok {
			return format(res)
		}
	}

	return "Failed..."
}

func format(ret []int) string {
	s := ""

	for i := range ret {
		s = s + strconv.Itoa(i)
	}

	return s
}

func removeAll(try [][]int, total [][]int) [][]int {
	ret := [][]int{}

	for i := 0; i < len(total); i++ {
		t := total[i]
		if !columnInArray(t, try) {
			ret = append(ret, t)
		}
	}

	return ret
}

func verify(try [][]int, remaining [][]int, n int) (bool, []int) {
	fmt.Println("Try:", try)
	fmt.Println("Remaining:", remaining)

	missing := [][]int{}

	for i := 0; i < n; i++ {
		column := []int{}
		for j := 0; j < n; j++ {
			column = append(column, try[j][i])
		}

		if !columnInArray(column, remaining) {
			missing = append(missing, column)
		}

		if len(missing) >= 2 {
			return false, []int{}
		}
	}

	fmt.Println("True, missing:", missing)
	return true,missing[0]
}

func columnInArray(column []int, array[][]int) bool {
	for i := 0; i < len(array); i++ {
		found := true
		for k := 0; found && k < len(column); k++ {
			if array[i][k] != column[k] {
				found = false
			}
		}
		return found
	}
	return false
}

func elect(chosen [][]int, sorted [][]int, n int, i int) [][][]int {
	if len(chosen) >= n {
		return [][][]int{chosen}
	}

	if i >= len(sorted) {
		return [][][]int{}
	}

	next := sorted[i]
	for strictCompare(chosen[len(chosen)-1], next) >= 0 {
		i += 1
		if i >= len(sorted) {
			return [][][]int{}
		}

		next = sorted[i]
	}
	newChosen := append(chosen, next)

	return append(elect(newChosen, sorted, n, i+1), elect(chosen, sorted, n, i+1)...)
}

//func sort(in [][]int) [][]int {
	

	// out := [][]int{}

	// for i := 0; i < len(in); i++ {
	// 	if len(out) == 0 || compare(out[len(out) - 1], in[i]) <= 0 {
	// 		if len(out) > 0 {
	// 			fmt.Println(out[len(out) - 1], "vs", in[i])
	// 		}
	// 		out = append(out, in[i])
	// 	} else {
	// 		out = append([][]int{in[i]}, out...)
	// 	}
	// }

	// return out
//}

type RowList [][]int

func (r RowList) Len() int {
	return len(r)
}

func (r RowList) Swap(i, j int) {
    r[i], r[j] = r[j], r[i]
}
func (r RowList) Less(i, j int) bool {
    return compare(r[i], r[j]) <= 0
}

func compare(a []int, b []int) int {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	return 0
}

func strictCompare(a []int, b []int) int {
	for i := 0; i < len(a); i++ {
		if a[i] >= b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	return 0
}
