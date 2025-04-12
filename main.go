package main

import "fmt"

type BloomFilter struct {
	bits             []bool
	expectedSize     int
	numberOfHashes   int
	falseProbability int
}

func NewBloomFilter(expectedSize int) (*BloomFilter, error) {
	return &BloomFilter{
		expectedSize: expectedSize,
	}, nil
}

func (bf *BloomFilter) String() string {
	return fmt.Sprintf("BloomFilter{expectedSize: %d, numberOfHashes: %d, falseProbability: %d}", bf.expectedSize, bf.numberOfHashes, bf.falseProbability)
}

// TODO
func (*BloomFilter) Insert(word string) {}

// TODO
func (*BloomFilter) Lookup(word string) {}

// TODO
func (*BloomFilter) GetSize() {}

// TODO
func (*BloomFilter) newBitsArray() {}

// TODO
func (*BloomFilter) calculateNumOfHashes() {}

// TODO
func (*BloomFilter) calculateFalseProbability() {}

func main() {
	bloomFilter, err := NewBloomFilter(10)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", bloomFilter.String())
}
