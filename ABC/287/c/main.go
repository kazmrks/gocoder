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

	var N, M int
	fmt.Scan(&N, &M)

	if N != M+1 {
		fmt.Println("No")
		return
	}

	path := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		path[i] = make([]int, 2)
	}

	for i := 0; i < M; i++ {
		node := lineToInts(2)
		u := node[0]
		v := node[1]

		if path[u][1] != 0 || path[v][1] != 0 {
			fmt.Println("No")
			return
		}

		if path[u][0] == 0 {
			path[u][0] = v
		} else if path[u][1] == 0 {
			if path[u][0] == v {
				fmt.Println("No")
				return
			}
			path[u][1] = v
		}

		if path[v][0] == 0 {
			path[v][0] = u
		} else if path[v][1] == 0 {
			if path[v][0] == u {
				fmt.Println("No")
				return
			}
			path[v][1] = u
		}
	}

	edge := 0
	for i := 1; i <= N; i++ {
		u := path[i][0]
		v := path[i][1]

		if u == 0 && v == 0 {
			fmt.Println("No")
			return
		}

		if (u == 0 && v != 0) || (u != 0 && v == 0) {
			edge++
		}
	}
	if edge == 2 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
