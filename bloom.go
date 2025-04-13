package gobloom

import (
	"errors"
	"fmt"
	hash "iperalta7/gobloom/internal"
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
	bf.getBitsSize(float64(expectedSize), float64(desiredFalseProbability))
	log.Printf("Initiated bit array size of %d", bf.bitSize)
	bf.initBitsArray()
	bf.calculateNumOfHashes()
	log.Printf("Calculated hash count for ops %d", bf.hashCount)
	return bf, nil
}

func (bf *BloomFilter) String() string {
	return fmt.Sprintf(
		"BloomFilter{bits: %v, bitSize: %d, ExpectedSize: %d, hashCount: %d, DesiredFalseProbability: %f}",
		bf.bits, bf.bitSize, bf.ExpectedSize, bf.hashCount, bf.DesiredFalseProbability,
	)
}

func (bf *BloomFilter) Insert(word string) {
	for i := range bf.hashCount {
		hash_int := hash.MurmurHash3(word, uint32(i)) // use i as seed
		index := int(math.Mod(float64(hash_int), float64(bf.bitSize)))
		bf.bits[index] = true
	}
	log.Println("inserted!")
}

// TODO
func (bf *BloomFilter) Lookup(word string) bool {
	for i := range bf.hashCount {
		hash_int := hash.MurmurHash3(word, uint32(i)) // use i as seed
		index := int(math.Mod(float64(hash_int), float64(bf.bitSize)))
		if !(bf.bits[index]) {
			return false
		}

	}
	log.Println("Exists!")
	return true
}

// Private
func (bf *BloomFilter) CalculateAFProbability() (float64, error) {
	if bf.bitSize <= 0 {
		return 0.0, errors.New("size of bit array must be >= zero to get false positive probability")
	}
	actualFp := math.Pow(1-math.Pow(math.Abs(float64(1-(1/float64(bf.bitSize)))), float64(bf.hashCount)*float64(bf.ExpectedSize)),
		float64(bf.hashCount))
	return actualFp, nil
}

func (bf *BloomFilter) getBitsSize(n float64, p float64) {
	m := -(n * math.Log(p)) /
		math.Pow(math.Log(2), 2)
	bf.bitSize = int(m)
}

func (bf *BloomFilter) initBitsArray() {
	array := make([]bool, bf.bitSize)
	bf.bits = array
}

func (bf *BloomFilter) calculateNumOfHashes() {
	k := float64(bf.bitSize/bf.ExpectedSize) * math.Log(2)
	bf.hashCount = int(k)
}
