package set

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSetAdd(t *testing.T) {
	set := New[string]()
	set.Add("Hello")
	set.Add("World")
	//set.Delete("Hello")
	assert.Equal(t, set.Contain("Hello"), true)
}

func ExampleForEach() {
	set := New[string]()
	set.Add("zhangsan")
	set.Add("lisi")
	set.Add("wangqiang")
	set.ForEach(func(v string) {
		fmt.Println("Hello", v)
	})

	//Output:
	//Hello zhangsan
	//Hello lisi
	//Hello wangqiang
}
