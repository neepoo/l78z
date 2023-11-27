package bloom

import "hash"

// Interface 接口定义了布隆过滤器的核心功能：
// 添加一个元素
// 判断元素是否存在
type Interface interface {
	Add([]byte)
	Test([]byte) bool
}

// Hasher 用于为我们的布隆过滤器生成hash值
type Hasher interface {
	GetHashes(n uint64) []hash.Hash64
}
