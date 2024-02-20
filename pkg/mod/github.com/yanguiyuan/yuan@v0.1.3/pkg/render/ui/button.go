package ui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/yanguiyuan/yuan/pkg/render"
	"github.com/yanguiyuan/yuan/pkg/render/consts"
	"github.com/yanguiyuan/yuan/pkg/render/consts/color"
)

//	type Button struct {
//		DefaultView
//		text string
//	}
//
//	func (b *Button) OnClick(f func(e Event)) {
//		f(Event{})
//	}
//
//	func (b Button) String() string {
//		return "Button(text=\"" + b.text + "\")"
//	}
//
//	func NewButton(ctx context.Context, text string) View {
//		c := Button{text: text}
//		Child(ctx, &c)
//		cv := context.WithValue(ctx, consts.ParentKey, &c)
//		c.ctx = cv
//		return &c
//	}
//
//	func (b *Button) ToCanvasObject() fyne.CanvasObject {
//		btn := widget.NewButton(b.text, b.click)
//		btn.Importance = widget.HighImportance
//		return btn
//	}

type DefaultButton struct {
	DefaultView
	text  string
	click func()
}

func (b *DefaultButton) OnClick(click func()) {
	b.click = click
}

type ButtonStyle interface {
	ViewStyle
	Text(text string) ButtonStyle
	TextSize(size float32) ButtonStyle
	TextAlign(align fyne.TextAlign) ButtonStyle
	BackgroundColor(color color.Color) ButtonStyle
	TextColor(color color.Color) ButtonStyle
	TextBold(bold bool) ButtonStyle
	TextItalic(italic bool) ButtonStyle
	TextUnderline(underline bool) ButtonStyle
}
type FyneButton struct {
	DefaultButton
}

func (b *FyneButton) Render(ctx context.Context) {
	c := ctx.Value(consts.ContextKeyRenderParentObject).(*fyne.Container)
	btn := widget.NewButton(b.text, b.click)
	btn.Importance = widget.HighImportance
	c.Add(btn)
}
func NewButton(ctx context.Context, text string) render.Button {
	renderer := render.GetApplicationRenderer(ctx)
	switch renderer {
	case consts.RendererFyne:
		return newFyneButton(ctx, text)
	default:
		panic("not support")
	}
}
func newFyneButton(ctx context.Context, text string) render.Button {
	c := FyneButton{}
	cv := context.WithValue(ctx, consts.ContextKeyParentView, &c)
	parent := render.GetParentView(ctx)
	parent.Child(&c)
	c.ctx = cv
	c.text = text
	return &c
}
