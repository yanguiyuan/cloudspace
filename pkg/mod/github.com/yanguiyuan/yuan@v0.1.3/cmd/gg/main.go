package main

import (
	"github.com/fogleman/gg"
)

func main() {
	const (
		width        = 400
		height       = 300
		cornerRadius = 20
		text         = "Hello"
	)

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 设置背景颜色为白色
	dc.Clear()

	dc.SetRGB(0, 0, 0) // 设置圆角矩形的颜色为黑色
	dc.SetLineWidth(2) // 设置线宽

	// 绘制圆角矩形
	dc.DrawRoundedRectangle(50, 50, float64(width-100), float64(height-100), cornerRadius)
	dc.Stroke()

	// 添加文字
	fontSize := 40
	dc.SetRGB(0, 0, 0)                          // 设置文字颜色为黑色
	dc.LoadFontFace("Arial", float64(fontSize)) // 加载字体和字号
	textWidth, textHeight := dc.MeasureString(text)

	// 计算放大比例
	scale := 2.0
	dc.Push()
	dc.Scale(scale, scale) // 放大文字
	textWidth *= scale
	textHeight *= scale
	textX := (float64(width) - textWidth) / 2
	textY := (float64(height) + textHeight) / 2
	dc.DrawStringAnchored(text, textX, textY, 0.5, 0.5)
	dc.Pop()

	dc.SavePNG("temp/rounded_rectangle_with_scaled_text.png") // 将绘制的圆角矩形和文字保存为图片文件
}
