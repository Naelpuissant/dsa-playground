package bloom

import (
	"ds/bitmap"
	"encoding/binary"
	"errors"
	"math"
)

var (
	ErrWrongFalsePositive = errors.New("falsePositive should be between 0 and 1")
	ErrWrongNItems        = errors.New("nItems should be greater than zero")
)

type BloomFilter struct {
	bitmap    *bitmap.BitMap
	nbits     int
	nhashes   int
	digestBuf [4]byte
}

func getNBits(falsePositive float64, nItems float64) int {
	nbits := int(-(nItems * math.Log(falsePositive)) / math.Pow(math.Ln2, 2))

	// round up to nearest multiple of 64
	// divid and round up to get number of 64 bit blocs
	// multiply by 64 to get total bits
	return (nbits + 64 - 1) / 64 * 64
}

func getNHashes(nbits float64, nItems float64) int {
	nhashes := nbits / nItems * math.Ln2
	return int(math.Ceil(nhashes))
}

// Create a new BloomFilte
func New(falsePositive float64, nItems int) *BloomFilter {
	if falsePositive <= 0 || falsePositive >= 1 {
		panic(ErrWrongFalsePositive)
	}

	if nItems <= 0 {
		panic(ErrWrongNItems)
	}

	nbits := getNBits(falsePositive, float64(nItems))
	nhashes := getNHashes(float64(nbits), float64(nItems))

	bitmap, err := bitmap.New(nbits)
	if err != nil {
		panic(err)
	}

	return &BloomFilter{
		bitmap:  bitmap,
		nbits:   nbits,
		nhashes: int(nhashes),
	}
}

func (bf *BloomFilter) Add(key []byte) {
	digest := binary.BigEndian.AppendUint32(bf.digestBuf[:0], Hash(key))

	h1 := binary.BigEndian.Uint16(digest[:2])
	h2 := binary.BigEndian.Uint16(digest[2:4])
	h2 |= 1 // avoid 0 h2

	nbits := uint16(bf.nbits)
	for i := range bf.nhashes {
		idx := (h1 + uint16(i)*h2) % nbits
		bf.bitmap.Set(int(idx))
	}
}

// Contains returns :
// true if a key is "maybe" in the bloom filter,
// false if a key is not the bloom filter
func (bf *BloomFilter) Contains(key []byte) bool {
	digest := binary.BigEndian.AppendUint32(bf.digestBuf[:0], Hash(key))

	h1 := binary.BigEndian.Uint16(digest[:2])
	h2 := binary.BigEndian.Uint16(digest[2:4])
	h2 |= 1 // avoid 0 h2

	nbits := uint16(bf.nbits)
	for i := range bf.nhashes {
		idx := (h1 + uint16(i)*h2) % nbits
		if !bf.bitmap.IsSet(int(idx)) {
			return false
		}
	}

	return true
}

// hash implements a hashing algorithm similar to the Murmur hash.
// https://github.com/dgraph-io/badger/blob/796cb85f662c06cb4660f0143af176ca7cf1a857/y/bloom.go
func Hash(b []byte) uint32 {
	const (
		seed = 0xbc9f1d34
		m    = 0xc6a4a793
	)
	h := uint32(seed) ^ uint32(len(b))*m
	for ; len(b) >= 4; b = b[4:] {
		h += uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
		h *= m
		h ^= h >> 16
	}
	switch len(b) {
	case 3:
		h += uint32(b[2]) << 16
		fallthrough
	case 2:
		h += uint32(b[1]) << 8
		fallthrough
	case 1:
		h += uint32(b[0])
		h *= m
		h ^= h >> 24
	}
	return h
}
