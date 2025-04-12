package main

import (
	"fmt"
	"math"
)

type BloomFilter struct {
	bits                    []bool
	expectedSize            int // n
	numberOfHashes          int // k
	desiredFalseProbability float32
}

func New(expectedSize int, desiredFalseProbability float32) (*BloomFilter, error) {
	bf := &BloomFilter{
		expectedSize:            expectedSize,
		desiredFalseProbability: desiredFalseProbability,
	}

	bf.initBitsArray(float64(expectedSize), float64(desiredFalseProbability))

	return bf, nil
}

func (bf *BloomFilter) String() string {
	return fmt.Sprintf("BloomFilter{expectedSize: %d, numberOfHashes: %d, falseProbability: %f}", bf.expectedSize, bf.numberOfHashes, bf.desiredFalseProbability)
}

// TODO
func (bf *BloomFilter) Insert(word string) {}

// TODO
func (bf *BloomFilter) Lookup(word string) {}

// TODO: aka m
func (bf *BloomFilter) getSize(n float64, p float64) int {
	m := -(n * math.Log(p)) /
		math.Pow(math.Log(2), 2)
	return int(m)
}

// TODO
func (bf *BloomFilter) initBitsArray(n float64, p float64) {
	m := bf.getSize(n, p)
	array := make([]bool, m)
	bf.bits = array
}

// TODO
func (bf *BloomFilter) calculateNumOfHashes() {}

// TODO
func (bf *BloomFilter) calculateFalseProbability() {}

func main() {
	bloomFilter, err := New(10, 0.05)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", bloomFilter.String())
}
