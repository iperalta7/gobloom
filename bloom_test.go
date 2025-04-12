package gobloom

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bloomFilter, err := New(10, 0.01)
	if err != nil {
		t.Fatalf("Failed to create bloom filter: %v", err)
	}
	fpProbability, err := bloomFilter.CalculateAFProbability()
	if err != nil {
		t.Fatalf("Error calculating false positive probability: %v", err)
	}
	t.Logf("False positive probability: %f", fpProbability)
}
