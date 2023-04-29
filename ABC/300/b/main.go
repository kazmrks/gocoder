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
	sc.Buffer([]byte{}, math.MaxInt64)

	l := lineToInts()

	A := [][]rune{}
	B := [][]rune{}

	CountA := []int{}
	CountB := []int{}

	for i := 0; i < l[0]; i++ {
		line := nextLine()
		r := []rune(line)

		count := 0
		for _, v := range r {
			if v == '.' {
				count++
			}
		}
		CountA = append(CountA, count)
		A = append(A, r)
	}

	for i := 0; i < l[0]; i++ {
		line := nextLine()
		r := []rune(line)

		count := 0
		for _, v := range r {
			if v == '.' {
				count++
			}
		}
		CountB = append(CountB, count)
		B = append(B, r)
	}

	shift := 0
	for i := 1; i < l[0]; i++ {
		if isMatchCount(CountA, shiftH(CountB, i)) {
			shift = i
			break
		}
	}

	if shift == 0 {
		fmt.Println("No")
		return
	}

	s := ""
	for _, v := range A {
		s = s + string(v[0])
	}

	for i := 0; i < l[1]; i++ {
		ss := shiftW(B, shift, i)
		if s == ss {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}

func shiftH(s []int, n int) []int {
	res := s[n:]
	res = append(res, s[:n]...)
	return res
}

func shiftW(r [][]rune, n int, pos int) string {
	res := r[n:]
	res = append(res, r[:n]...)

	s := ""
	for i := 0; i < len(res); i++ {
		s = s + string(res[i][pos])
	}
	return s
}

func isMatchCount(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
