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

	seq := [][]int{}
	for i := 0; i < M; i++ {
		seq = append(seq, lineToInts(2))
	}

	path := []int{seq[0][0], seq[0][1]}
	seq = seq[1:]
	ok := buildPath(path, seq)

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func buildPath(path []int, seq [][]int) bool {
	if len(seq) == 0 {
		return true
	}

	first := path[0]
	last := path[len(path)-1]
	removed := false

	for i, v := range seq {
		if v[0] == first {
			path = insertFirst(path, v[1])
			seq = removeAt(seq, i)
			removed = true
			break
		}

		if v[1] == first {
			path = insertFirst(path, v[0])
			seq = removeAt(seq, i)
			removed = true
			break
		}

		if v[0] == last {
			path = append(path, v[1])
			seq = removeAt(seq, i)
			removed = true
			break
		}

		if v[1] == last {
			path = append(path, v[0])
			seq = removeAt(seq, i)
			removed = true
			break
		}
	}

	if !removed {
		return false
	}

	return buildPath(path, seq)
}

func insertFirst(path []int, value int) []int {
	path = append(path[:1], path[0:]...)
	path[0] = value
	return path
}

func removeAt(seq [][]int, pos int) [][]int {
	seq = append(seq[:pos], seq[pos+1:]...)
	return seq
}
