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

type pair struct {
	a byte
	b byte
}

func f(x int) pair {
	xs := strconv.Itoa(x)
	a, b := xs[0], xs[len(xs)-1]
	return pair{a: a, b: b}
}

func run() {
	setSpace()
	n := readN()
	freq := map[pair]int{}
	for i := 1; i <= n; i++ {
		n, ok := freq[f(i)]
		if !ok {
			n = 0
		}
		n++
		freq[f(i)] = n
	}
	ans := 0
	for i := 1; i <= n; i++ {
		p := f(i)
		q := pair{a: p.b, b: p.a}
		ans += freq[q]
	}
	fmt.Println(ans)
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
