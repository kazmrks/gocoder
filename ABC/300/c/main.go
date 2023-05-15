package main

import (
	"github.com/kazmrks/gocoder/io"
)

func main() {
	line := io.LineToInts()
	H := line[0]
	W := line[1]

	C := [][]rune{}

	for i := 0; i < H; i++ {
		C = append(C, []rune(io.NextLine()))
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
				}
			} else {
				if C[i][j] == '#' {
					cross = true
				}
			}
		}
	}
	io.PrintInts(S)
}

func min(H, W int) int {
	if H < W {
		return H
	} else {
		return W
	}
}
