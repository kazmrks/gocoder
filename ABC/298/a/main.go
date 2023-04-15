package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	i, err := strconv.Atoi(nextLine())
	if err != nil {
		panic(err)
	}
	return i
}

func lineToInts(n int) []int {
	slice := make([]int, n)
	text := strings.Split(nextLine(), " ")

	for i := 0; i < n; i++ {
		v, err := strconv.Atoi(text[i])
		if err != nil {
			panic(err)
		}
		slice[i] = v
	}
	return slice
}

func contains(elements []int, v int) bool {
	for _, e := range elements {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	sc.Buffer([]byte{}, math.MaxInt64)

	var N int
	fmt.Scan(&N)
	S := nextLine()

	ok := false

	for _, c := range S {
		if c == 'x' {
			fmt.Println("No")
			return
		} else if c == 'o' {
			ok = true
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
