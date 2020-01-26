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
	a, b := readN(), readN()
	s, l := 0, 0
	if a < b {
		s, l = a, b
	} else {
		s, l = b, a
	}
	for i := 0; i < l; i++ {
		fmt.Printf(strconv.Itoa(s))
	}
	fmt.Println("")
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
