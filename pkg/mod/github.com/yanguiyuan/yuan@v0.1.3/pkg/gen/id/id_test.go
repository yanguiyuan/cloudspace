package id

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		id := One()
		fmt.Println(i, " : ", id)
	}
}
