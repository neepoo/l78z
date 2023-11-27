package bloom

import (
	"fmt"
	"hash"
	"math"
	"sync"
)

// BloomFilter 表示单个 Bloom 筛选器结构。
type BloomFilter struct {
	bitSet []bool
	m      uint64        // len(bitSet)
	hashes []hash.Hash64 // 使用的hash函数
	k      uint64        // 使用的hash函数数量 len(hashes)
	mutex  sync.Mutex
}

func (bf *BloomFilter) Add(data []byte) {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()
	var i uint64
	for i = 0; i < bf.k; i++ {
		hs := bf.hashes[i]
		hs.Reset()
		hs.Write(data)
		hashVal := hs.Sum64() % bf.m
		bf.bitSet[hashVal] = true
	}
}

func (bf *BloomFilter) Test(data []byte) bool {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()

	var i uint64
	for i = 0; i < bf.k; i++ {
		hs := bf.hashes[i]
		hs.Reset()
		hs.Write(data)
		hashVal := hs.Sum64() % bf.m
		if !bf.bitSet[hashVal] {
			// 肯定不在集合中
			return false
		}
	}
	// 误报的来源。如果都是true,则认为该元素可能存在于集合中
	return true
}

// NewBloomFilterWithHasher 创建一个新的布隆过滤器，参数为给定的元素数量(n)和误报率(p。
func NewBloomFilterWithHasher(n uint64, p float64, h Hasher) (*BloomFilter, error) {
	if n == 0 {
		return nil, fmt.Errorf("number of elements cannot be 0")
	}
	if p <= 0 || p >= 1 {
		return nil, fmt.Errorf("false positive rate must be between 0 and 1")
	}
	if h == nil {
		return nil, fmt.Errorf("hasher cannot be nil")
	}
	m, k := getOptimalParams(n, p)
	return &BloomFilter{
		m:      m,
		k:      k,
		bitSet: make([]bool, m),
		hashes: h.GetHashes(k),
	}, nil
}

// getOptimalParams 函数计算布隆过滤器的最优参数，包括位集合中的位数(m)和哈希函数的数量(k)。
func getOptimalParams(n uint64, p float64) (uint64, uint64) {
	m := uint64(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	if m == 0 {
		m = 1
	}
	k := uint64(math.Ceil((float64(m) / float64(n)) * math.Log(2)))
	if k == 0 {
		k = 1
	}
	return m, k
}
