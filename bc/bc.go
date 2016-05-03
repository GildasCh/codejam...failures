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
	L []string
	R []string
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
		data := InputData{0,[]string{},[]string{}}
        // fmt.Println(scanner.Text())
		scanner.Scan()
		challenge := strings.Split(scanner.Text(), " ")
		data.N, _ = strconv.Atoi(challenge[0])

		for j := 0; j < data.N; j++ {
			scanner.Scan()
			readRow := strings.Split(scanner.Text(), " ")
			l := readRow[0]
			data.L = append(data.L, l)
			r := readRow[1]
			data.R = append(data.R, r)
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
	lmap := OccurMap(in.L)
	rmap := OccurMap(in.R)

	// score := 0
	// for i := 0; i < in.N; i++ {
	// 	li, lok := lmap[in.L[i]]
	// 	ri, rok := rmap[in.R[i]]

	// 	if lok && li > 1 && rok && ri > 1 {
	// 		score += 1
	// 		lmap[in.L[i]] -= 1
	// 		rmap[in.R[i]] -= 1
	// 	}
	// }

	lsum := 0
	rsum := 0
	for i := 0; i < in.N; i++ {
		lsum += lmap[in.L[i]] - 1
		rsum += rmap[in.R[i]] - 1
	}

	if lsum > rsum {
		return strconv.Itoa(rsum)
	}
	return strconv.Itoa(lsum)

	// return strconv.Itoa(score)
}

func OccurMap(words []string) map[string]int {
	ret := make(map[string]int)

	for _, w := range(words) {
		i, ok := ret[w]
		if ok {
			ret[w] = i + 1
		} else {
			ret[w] = 1
		}
	}

	return ret
}
