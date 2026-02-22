package bf_test

import (
	"crypto/sha1"
	bf "ds/bloom"
	"hash"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := bf.NewBloomFilter(0.01, 4, func() hash.Hash {
		return sha1.New()
	})

	keys := []string{"hello", "world", "foo", "bar"}
	for _, key := range keys {
		bf.Add([]byte(key))
	}

	for _, key := range keys {
		if !bf.Contains([]byte(key)) {
			t.Errorf("Expected BloomFilter to contain key %s, but it does not", key)
		}
	}
}

// TODO : add a concurrency test
