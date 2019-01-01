package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(run())
}

func run() int {
	setSpace()
	start := readN()
	end := readN()

	m := make(map[int]bool, 31)
	for i := start - 1; i < end; i++ {
		m[i] = true
	}

	cnt := 0
	if m[22] {
		cnt++
	}
	if m[23] {
		cnt++
	}
	if m[24] {
		cnt++
	}
	return 3 - cnt
}

var sc = bufio.NewScanner(os.Stdin)

func setSpace() {
	sc.Split(bufio.ScanWords)
}

func setLine() {
	sc.Split(bufio.ScanLines)
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
