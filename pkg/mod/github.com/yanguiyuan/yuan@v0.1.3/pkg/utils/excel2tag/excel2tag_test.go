package excel2tag_test

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/utils/excel2tag"
	"os"
	"testing"
)

func TestExcel2Tag(t *testing.T) {
	file, err := os.Open("temp/lxy.xlsx")
	if err != nil {
		return
	}
	worker := excel2tag.New(file, excel2tag.WithCommentColumnIndex(11), excel2tag.WithPrimaryKeyColumnIndex(1))
	worker.Parse()
	tagTypeSet := worker.GetTagTypeList()
	fmt.Println(tagTypeSet)
}
