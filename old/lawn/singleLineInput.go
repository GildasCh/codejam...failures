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
	return strconv.Itoa(in.N)
}
