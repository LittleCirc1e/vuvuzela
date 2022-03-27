package vuvuzela

import (
	"encoding/binary"
	"io"
)

type Shuffler []int

func NewShuffler(rand io.Reader, n int) Shuffler {
	p := make(Shuffler, n)
	for i := range p {
		p[i] = Intn(rand, i+1)
	}
	return p
}

func (s Shuffler) Shuffle(x [][]byte) {
	for i := range x {
		j := s[i]
		x[i], x[j] = x[j], x[i]
	}
}

func (s Shuffler) Unshuffle(x [][]byte) {
	for i := len(x) - 1; i >= 0; i-- {
		j := s[i]
		x[i], x[j] = x[j], x[i]
	}
}

func Intn(rand io.Reader, n int) int {
	max := ^uint32(0)
	m := max % uint32(n)
	r := make([]byte, 4)
	for {
		if _, err := rand.Read(r); err != nil {
			panic(err)
		}
		x := binary.BigEndian.Uint32(r)
		if x < max-m {
			return int(x % uint32(n))
		}
	}
}
