package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	run()
}

func run() {
	setSpace()
}

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
