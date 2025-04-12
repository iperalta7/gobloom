package main

import "fmt"

type BloomFilter struct {
	bits                    []bool
	expectedSize            int // n
	numberOfHashes          int // k
	desiredFalseProbability float32
}

func New(expectedSize int, desiredFalseProbability float32) (*BloomFilter, error) {
	return &BloomFilter{
		expectedSize:            expectedSize,
		desiredFalseProbability: desiredFalseProbability,
	}, nil
}

func (bf *BloomFilter) String() string {
	return fmt.Sprintf("BloomFilter{expectedSize: %d, numberOfHashes: %d, falseProbability: %f}", bf.expectedSize, bf.numberOfHashes, bf.desiredFalseProbability)
}

// TODO
func (*BloomFilter) Insert(word string) {}

// TODO
func (*BloomFilter) Lookup(word string) {}

// TODO: aka m
func (*BloomFilter) GetSize() {}

// TODO
func (*BloomFilter) newBitsArray() []bool {}

// TODO
func (*BloomFilter) calculateNumOfHashes() {}

// TODO
func (*BloomFilter) calculateFalseProbability() {}

func main() {
	bloomFilter, err := New(10, 0.05)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", bloomFilter.String())
}
