package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type BloomFilter struct {
	bits                    []bool
	bitSize                 int // m
	ExpectedSize            int // n
	hashCount               int // k
	DesiredFalseProbability float32
}

func New(expectedSize int, desiredFalseProbability float32) (*BloomFilter, error) {
	bf := &BloomFilter{
		ExpectedSize:            expectedSize,
		DesiredFalseProbability: desiredFalseProbability,
	}

	bf.getBitsSize(float64(expectedSize), float64(desiredFalseProbability))
	bf.initBitsArray()
	bf.calculateNumOfHashes()
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
func (bf *BloomFilter) getBitsSize(n float64, p float64) {
	m := -(n * math.Log(p)) /
		math.Pow(math.Log(2), 2)
	bf.bitSize = int(m)
}

// TODO
func (bf *BloomFilter) initBitsArray() {
	array := make([]bool, bf.bitSize)
	bf.bits = array
}

// TODO
func (bf *BloomFilter) calculateNumOfHashes() {
	k := float64(bf.bitSize/bf.ExpectedSize) * math.Log(2)
	bf.hashCount = int(k)
}

// TODO
func (bf *BloomFilter) calculateFalseProbability() {}

func main() {
	bloomFilter, err := New(10, 0.05)
	if err != nil {
		return
	}
	bloomFilter.String()
}
