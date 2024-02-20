package ui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/yanguiyuan/yuan/pkg/render"
	"github.com/yanguiyuan/yuan/pkg/render/consts"
)

//	type Text struct {
//		DefaultView
//		text string
//	}
//
//	func (t Text) String() string {
//		return "Text(\"" + t.text + "\")"
//	}
//
//	func (t *Text) ToCanvasObject() fyne.CanvasObject {
//		text := widget.NewLabel(t.text)
//		return text
//	}
//
//	func NewText(ctx context.Context, text string) {
//		c := Text{text: text}
//		Child(ctx, &c)
//		cv := context.WithValue(ctx, consts.ParentKey, &c)
//		c.ctx = cv
//	}

type DefaultText struct {
	DefaultView
	text string
}

func (t *DefaultText) GetText() string {
	return t.text
}
func (t *DefaultText) Text(text string) render.Text {
	t.text = text
	return t
}
func (t *FyneText) Text(text string) render.Text {
	t.text = text
	if t.handle != nil {
		t.handle.SetText(text)
	}
	return t
}

type FyneText struct {
	DefaultText
	handle *widget.Label
}

func (b *FyneText) Render(ctx context.Context) {
	c := ctx.Value(consts.ContextKeyRenderParentObject).(*fyne.Container)
	label := widget.NewLabel(b.text)
	b.handle = label
	c.Add(label)
}
func newFyneText(ctx context.Context, text string) render.Text {
	c := FyneText{
		DefaultText: DefaultText{text: text},
	}
	cv := context.WithValue(ctx, consts.ContextKeyParentView, &c)
	parent := render.GetParentView(ctx)
	parent.Child(&c)
	c.ctx = cv
	return &c
}
func NewText(ctx context.Context, text string) render.Text {
	renderer := render.GetApplicationRenderer(ctx)
	switch renderer {
	case consts.RendererFyne:
		return newFyneText(ctx, text)
	default:
		panic("not support")
	}
}
