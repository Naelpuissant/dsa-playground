package bloom_test

import (
	"ds/bloom"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bloom := bloom.New(0.01, 4)

	keys := []string{"hello", "world", "foo", "bar"}
	for _, key := range keys {
		bloom.Add([]byte(key))
	}

	for _, key := range keys {
		if !bloom.Contains([]byte(key)) {
			t.Errorf("Expected BloomFilter to contain key %s, but it does not", key)
		}
	}
}

// TODO : add a concurrency test
