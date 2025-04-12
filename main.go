package main

import (
	"errors"
	"fmt"
	"log"
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
	logger := log.Default()
	bf.getBitsSize(float64(expectedSize), float64(desiredFalseProbability))
	logger.Printf("Initiated bit array size of %d", bf.bitSize)
	bf.initBitsArray()
	bf.calculateNumOfHashes()
	logger.Printf("Calculated hash count for ops %d", bf.hashCount)
	return bf, nil
}

func (bf *BloomFilter) String() string {
	return fmt.Sprintf(
		"BloomFilter{bits: %v, bitSize: %d, ExpectedSize: %d, hashCount: %d, DesiredFalseProbability: %f}",
		bf.bits, bf.bitSize, bf.ExpectedSize, bf.hashCount, bf.DesiredFalseProbability,
	)
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

func (bf *BloomFilter) CalculateAFProbability() (float64, error) {
	if bf.bitSize <= 0 {
		return 0.0, errors.New("size of bit array must be >= zero to get false positive probability")
	}
	actualFp := math.Pow(1-math.Pow(math.Abs(float64(1-(1/float64(bf.bitSize)))), float64(bf.hashCount)*float64(bf.ExpectedSize)),
		float64(bf.hashCount))
	return actualFp, nil
}

func main() {
	bloomFilter, err := New(10, 0.05)
	if err != nil {
		return
	}
	//fmt.Println(bloomFilter.String())
	fpProbability, err := bloomFilter.CalculateAFProbability()
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Printf("%f\n", fpProbability)
}
