package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
	"strings"
	"math"
)

func main() {
	//fmt.Println(getDigitsFromInt(2847))
	// return

	file, err := os.Open(os.Args[1])
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
        // fmt.Println(scanner.Text())
		scanner.Scan()
		challenge := strings.Split(scanner.Text(), " ")
		k, _ := strconv.Atoi(challenge[0])
		c, _ := strconv.Atoi(challenge[1])
		s, _ := strconv.Atoi(challenge[2])
		res := calculate(k, c, s)
		fmt.Print("Case #", counter, ": ", res)
		fmt.Println()
		counter += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func calculate(k int, c int, s int) string {
	if c == 1 {
		if k > s {
			return "IMPOSSIBLE"
		}
		res := "1"
		for i := 2; i <= k; i++ {
			res += " " + strconv.Itoa(i)
		}
		return res
	}

	if k - 1 > s {
		return "IMPOSSIBLE"
	}

	blockSize := int64(math.Pow(float64(k), float64(c - 1)))
	res := strconv.FormatInt(blockSize + int64(1), 10)
	for i := int64(2); i <= int64(k); i++ {
		res += " " + strconv.FormatInt(i*blockSize + i, 10)
	}
	return res
}
