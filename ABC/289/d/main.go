package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func lineToSlice() []int {
	slice := []int{}
	text := strings.Split(nextLine(), " ")
	for _, v := range text {
		i, _ := strconv.Atoi(v)
		slice = append(slice, i)
	}
	return slice
}

func main() {
	// N, _ := strconv.Atoi(nextLine())
	nextLine()
	A := lineToSlice()
	// M, _ := strconv.Atoi(nextLine())
	nextLine()
	B := lineToSlice()
	X, _ := strconv.Atoi(nextLine())

	if stepup(X, A, B, 0, false) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func stepup(X int, A []int, B []int, step int, reached bool) bool {
	if reached {
		return true
	}

	for _, a := range A {
		next := step + a
		if next == X {
			return true
		} else if next < X && !contains(B, next) {
			reached = stepup(X, A, B, next, reached)
		}
	}
	return reached
}

func contains(elements []int, v int) bool {
	for _, e := range elements {
		if e == v {
			return true
		}
	}
	return false
}
