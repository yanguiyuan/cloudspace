package main

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/gen/id"
)

func main() {
	i := id.Base62()
	fmt.Println(i)
}
