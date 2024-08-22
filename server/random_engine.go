package server

import (
	"crypto/rand"
	"math/big"
)

type RandomEngine struct {
	existed map[string]bool
	max     *big.Int
}

func NewRandomEngine() *RandomEngine {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	return &RandomEngine{
		existed: make(map[string]bool),
		max:     max,
	}
}

func (e *RandomEngine) Next() (string, error) {
	for {
		n, err := rand.Int(rand.Reader, e.max)
		if err != nil {
			return "", err
		}

		// Base32 number is for saving memory in map.
		base32 := n.Text(32)
		if !e.existed[base32] {
			e.existed[base32] = true

			// The base 10 number is for representing to user.
			return n.Text(10), nil
		}
	}
}
