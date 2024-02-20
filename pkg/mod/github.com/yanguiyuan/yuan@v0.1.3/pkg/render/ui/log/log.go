package log

import (
	"context"
	"fmt"
)

func Info(ctx context.Context, msg string) {
	fmt.Println(msg)
}
