package bloomfilter

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := New(1000, 3)

	data1 := []byte("apple")
	data2 := []byte("banana")
	data3 := []byte("orange")

	bf.Add(data1)
	bf.Add(data2)

	fmt.Println(bf.Contains(data1)) // Output: true
	fmt.Println(bf.Contains(data2)) // Output: true
	fmt.Println(bf.Contains(data3)) // Output: false
}
