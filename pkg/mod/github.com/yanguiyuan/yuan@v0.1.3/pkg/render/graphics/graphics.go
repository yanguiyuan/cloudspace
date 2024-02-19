package graphics

import (
	"context"
	"github.com/yanguiyuan/yuan/pkg/render/consts/color"
)

func DrawLine2D(ctx context.Context, x1, y1, x2, y2 float64) {
}
func DrawRect2D(ctx context.Context, x, y, w, h float64, color color.Color) {
	for i := 0; i < 10; i++ {
		DrawLine2D(ctx, x, y, x+w, y)
		DrawLine2D(ctx, x, y, x, y+h)
		DrawLine2D(ctx, x+w, y, x+w, y+h)
		DrawLine2D(ctx, x, y+h, x+w, y+h)
	}
}
func DrawCircle(ctx context.Context) {

}
func DrawTriangle(ctx context.Context) {

}

// DrawArc 画弧线
func DrawArc(ctx context.Context) {

}
