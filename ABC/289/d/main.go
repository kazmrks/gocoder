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
	N, _ := strconv.Atoi(nextLine())
	A := lineToSlice()
	M, _ := strconv.Atoi(nextLine())
	B := lineToSlice()
	X, _ := strconv.Atoi(nextLine())

	steps := make([]bool, X+1)
	steps[0] = true

	mochi := map[int]bool{}
	for i := 0; i < M; i++ {
		mochi[B[i]] = true
	}

	for i := 0; i < X; i++ {
		if !steps[i] {
			continue
		}

		for j := 0; j < N; j++ {
			next := i + A[j]
			_, m := mochi[next]
			if next <= X && !m {
				steps[next] = true
			}
		}
	}

	if steps[X] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
