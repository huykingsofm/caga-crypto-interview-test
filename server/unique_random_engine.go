package server

import (
	"crypto/rand"
	"errors"
	"math/big"
	"sync"
)

// UniqueRandomEngine supports creating unique random big integers.
type UniqueRandomEngine struct {
	mutex sync.RWMutex

	// existed contains all random number which is generated before
	existed map[string]bool

	// max is the max value of the random value (exclusive).
	max *big.Int
}

// NewUniqueRandomEngine creates a default RandomEngine with max value is
// 2^130-1.
func NewUniqueRandomEngine() *UniqueRandomEngine {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	return &UniqueRandomEngine{
		mutex:   sync.RWMutex{},
		existed: make(map[string]bool),
		max:     max,
	}
}

// NewUniqueRandomEngineWithMax creates a customized max value RandomEngine.
// Return the default UniqueRandomEngine if max is zero.
func NewUniqueRandomEngineWithMax(max int64) *UniqueRandomEngine {
	// In the case max value is 0, we should return the default engine.
	if max == 0 {
		return NewUniqueRandomEngine()
	}

	bigMax := new(big.Int)
	bigMax.SetInt64(max)

	return &UniqueRandomEngine{
		mutex:   sync.RWMutex{},
		existed: make(map[string]bool),
		max:     bigMax,
	}
}

// Next generates a unique random integer in string.
func (e *UniqueRandomEngine) Next() (string, error) {
	if e.max.IsInt64() && e.max.Int64() == int64(len(e.existed)) {
		return "", errors.New("no more unique random value")
	}

	for {
		n, err := rand.Int(rand.Reader, e.max)
		if err != nil {
			return "", err
		}

		// Base32 number is for saving memory of existed map.
		base32 := n.Text(32)

		// In the case the number never exists before, it returns the number in
		// string intermediately; otherwise, it will re-generate the number.
		e.mutex.RLock()
		existed := e.existed[base32]
		e.mutex.RUnlock()

		if !existed {
			// Mark the number as existed to prevent it will be generated in the
			// future.
			e.mutex.Lock()
			e.existed[base32] = true
			e.mutex.Unlock()

			// The base 10 number is for representing to user.
			return n.Text(10), nil
		}
	}
}
