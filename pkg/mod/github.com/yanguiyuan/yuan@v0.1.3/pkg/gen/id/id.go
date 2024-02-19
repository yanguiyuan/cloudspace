package id

import (
	"github.com/sony/sonyflake"
	"github.com/yanguiyuan/yuan/pkg/utils/convert"
)

var sf *sonyflake.Sonyflake

func init() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
}
func One() uint64 {
	id, err := sf.NextID()
	for err != nil {
		id, err = sf.NextID()
	}
	return id
}

func Base62() string {
	// 生成一个唯一的ID
	id := One()
	// 将ID转换为Base62编码的字符串
	return convert.Int64ToBase62(int64(id))
}
