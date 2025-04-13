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

	bloomFilter.Insert("test")
	if !bloomFilter.Lookup("test") {
		t.Errorf("Expected 'test' to be found in the bloom filter")
	}

	if bloomFilter.Lookup("go") == true {
		t.Errorf("Expected 'go' to not be found in bloom filter")
	}
}
