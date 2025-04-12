package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type BloomFilter struct {
	Bits                    []bool
	ExpectedSize            int // n
	HashCount               int // k
	DesiredFalseProbability float32
}

func New(expectedSize int, desiredFalseProbability float32) (*BloomFilter, error) {
	bf := &BloomFilter{
		ExpectedSize:            expectedSize,
		DesiredFalseProbability: desiredFalseProbability,
	}

	bf.initBitsArray(float64(expectedSize), float64(desiredFalseProbability))

	return bf, nil
}

func (bf *BloomFilter) String() {
	jsonBF, err := json.Marshal(bf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", string(jsonBF))
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
	bf.Bits = array
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
	bloomFilter.String()
}
