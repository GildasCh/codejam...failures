package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
	"strings"
	// "math"
)

type InputData struct {
	N int
	A []int
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
		data := InputData{0,[]int{}}
        // fmt.Println(scanner.Text())
		scanner.Scan()
		challenge := strings.Split(scanner.Text(), " ")
		data.N, _ = strconv.Atoi(challenge[0])

		scanner.Scan()
		readRow := strings.Split(scanner.Text(), " ")
		for j := 0; j < data.N; j++ {
			a, _ := strconv.Atoi(readRow[j])
			data.A = append(data.A, a)
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
	max := 0
	for i := 0; i < len(in.A); i++ {
		score := rec([]int{i}, []int{0}, in.A)
		if score > max {
			max = score
		}
	}

	return strconv.Itoa(max)
}

func rec(circle []int, anchors []int, bff []int) int {
	if len(anchors) <= 0 {
		return len(circle)
	}

	max := 0
	for index, a := range anchors {
		tries := bffOf(circle[a], bff)
		for _, t := range tries {
			if !in(t, circle) {
				newAnchors := remove(index, anchors)
				newAnchors = append(newAnchors, t)
				score := rec(append(circle, t), newAnchors, bff)
				if score > max {
					max = score
				}
			}
		}
	}

	return max
}

func bffOf(a int, bff []int) []int {
	ret := []int{}

	for b, f := range bff {
		if f == a {
			ret = append(ret, b)
		}
	}

	return ret
}

func in(a int, list []int) bool {
	for _, v := range list {
		if v == a {
			return false
		}
	}
	return true
}

func remove(index int, list []int) []int {
	if index+1 >= len(list) {
		return list[:index]
	}
	return append(list[:index], list[index+1:]...)
}
