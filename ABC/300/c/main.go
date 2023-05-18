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

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
}

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

func lineToInts() []int {
	slice := []int{}
	text := strings.Split(nextLine(), " ")

	for _, t := range text {
		v, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		slice = append(slice, v)
	}
	return slice
}

func printInts(s []int) {
	p := []string{}
	for _, v := range s {
		p = append(p, strconv.Itoa(v))
	}
	fmt.Println(strings.Join(p, " "))
}

func printStrings(s []string) {
	fmt.Println(strings.Join(s, " "))
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
	line := lineToInts()
	H := line[0]
	W := line[1]

	C := [][]rune{}

	for i := 0; i < H; i++ {
		C = append(C, []rune(nextLine()))
	}

	S := make([]int, min(H, W))

	for i := 0; i < H; i++ {
		n := 0
		cross := false
		for j := 0; j < W; j++ {
			if cross {
				n++
				if C[i][j] == '#' {
					cross = false
					n = n / 2
					S[n-1]++

					a := i + n
					b := j - n
					C[a][b] = '.'
					for k := n; k > 0; k-- {
						C[a+k][b+k] = '.'
						C[a+k][b-k] = '.'
						C[a-k][b+k] = '.'
						C[a-k][b-k] = '.'
					}
					n = 0
				}
			} else {
				if C[i][j] == '#' {
					cross = true
				}
			}
		}
	}
	printInts(S)
}

func min(H, W int) int {
	if H < W {
		return H
	} else {
		return W
	}
}
