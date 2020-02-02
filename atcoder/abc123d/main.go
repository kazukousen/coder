package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	run()
}

func run() {
	setSpace()
	xs, ys, zs, k := readAll()
	solve(1, xs, ys, zs, k)
}

func solve(p int, xs, ys, zs []int, k int) bool {
	var cnt int
	for _, x := range xs {
		for _, y := range ys {
			for _, z := range zs {
				if x+y+z < p {
					break
				}
				cnt++
				if cnt > k {
					return true
				}
			}
		}
	}
	return false
}

func readAll() ([]int, []int, []int, int) {
	x, y, z, k := readN(), readN(), readN(), readN()
	xs := make([]int, x)
	for i := 0; i < x; i++ {
		xs[i] = readN()
	}
	sort.Slice(xs, func(a, b int) bool {
		return a < b
	})
	ys := make([]int, y)
	for i := 0; i < y; i++ {
		ys[i] = readN()
	}
	sort.Slice(ys, func(a, b int) bool {
		return a < b
	})
	zs := make([]int, z)
	for i := 0; i < z; i++ {
		zs[i] = readN()
	}
	sort.Slice(zs, func(a, b int) bool {
		return a < b
	})
	return xs, ys, zs, k
}

// ------以下、ユーティリティ------

var sc = bufio.NewScanner(os.Stdin)

func setSpace() {
	sc.Split(bufio.ScanWords)
}

func read() string {
	sc.Scan()
	return sc.Text()
}

func readN() int {
	n, err := strconv.Atoi(read())
	if err != nil {
		panic(err)
	}
	return n
}
