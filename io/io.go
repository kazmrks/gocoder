package io

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var Scanner = bufio.NewScanner(os.Stdin)

func init() {
	Scanner.Buffer([]byte{}, math.MaxInt64)
}

func NextLine() string {
	Scanner.Scan()
	return Scanner.Text()
}

func ScanInt() int {
	i, err := strconv.Atoi(NextLine())
	if err != nil {
		panic(err)
	}
	return i
}

func LineToInts() []int {
	slice := []int{}
	text := strings.Split(NextLine(), " ")

	for _, t := range text {
		v, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		slice = append(slice, v)
	}
	return slice
}

func PrintInts(s []int) {
	p := []string{}
	for _, v := range s {
		p = append(p, strconv.Itoa(v))
	}
	fmt.Println(strings.Join(p, " "))
}

func PrintStrings(s []string) {
	fmt.Println(strings.Join(s, " "))
}

func Contains(elements []int, v int) bool {
	for _, e := range elements {
		if e == v {
			return true
		}
	}
	return false
}
