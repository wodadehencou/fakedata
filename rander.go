package fakedata

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	mrand "math/rand"
)

var globalRand *mrand.Rand

func init() {
	buf := make([]byte, 8)
	crand.Read(buf)
	globalRand = rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(buf))))
}
