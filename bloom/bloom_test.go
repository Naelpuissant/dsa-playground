package bloom_test

import (
	"crypto/sha1"
	"ds/bloom"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := bloom.NewBloomFilter(sha1.New(), 0.01, 4)
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
