package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	run()
}

func run() {
	setSpace()
	a, b, c, d, e, k := readN(), readN(), readN(), readN(), readN(), readN()
	if abs(a, b) > k {
		fmt.Println(":(")
		return
	}
	if abs(a, c) > k {
		fmt.Println(":(")
		return
	}
	if abs(a, d) > k {
		fmt.Println(":(")
		return
	}
	if abs(a, e) > k {
		fmt.Println(":(")
		return
	}
	if abs(b, c) > k {
		fmt.Println(":(")
		return
	}
	if abs(b, d) > k {
		fmt.Println(":(")
		return
	}
	if abs(b, e) > k {
		fmt.Println(":(")
		return
	}
	if abs(c, d) > k {
		fmt.Println(":(")
		return
	}
	if abs(c, e) > k {
		fmt.Println(":(")
		return
	}
	if abs(d, e) > k {
		fmt.Println(":(")
		return
	}
	fmt.Println("Yay!")
}

func abs(x, y int) int {
	n := x - y
	if n > 0 {
		return n
	}
	return -n
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
