package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	run()
}

type state struct {
	n        int
	players  map[int]int
	who      int
	prevDraw int
	prev     string
	d        int
}

func (s *state) next() {
	s.who = (s.who + s.d + s.n) % s.n
}

func (s *state) put() {
	s.players[s.who]++
}

func (s *state) draw() {
	s.players[s.who] -= s.prevDraw
	s.prevDraw = 0
}

func run() {
	setSpace()
	n := readN()
	m := readN()

	s := &state{
		n:       n,
		players: make(map[int]int, n),
		d:       1,
	}
	for i := 0; i < m; i++ {
		card := read()
		if s.prev != card && strings.HasPrefix(s.prev, "draw") {
			s.draw()
			s.next()
		}
		s.put()
		switch card {
		case "drawtwo":
			s.prevDraw += 2
		case "drawfour":
			s.prevDraw += 4
		case "skip":
			s.next()
		case "reverse":
			s.d *= -1
		}
		s.prev = card
		s.next()
	}
	// 一人前に戻る
	s.who = (s.who - s.d + s.n) % s.n
	fmt.Printf("%d %d\n", s.who+1, s.players[s.who])
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
