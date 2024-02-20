package ui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/yanguiyuan/yuan/pkg/render"
	"github.com/yanguiyuan/yuan/pkg/render/consts"
)

//	type Column struct {
//		DefaultView
//		builder func(ctx context.Context)
//	}
//
//	func (c *Column) Build(builder func(ctx context.Context)) *Option {
//		//TODO implement me
//		c.builder = builder
//		return &Option{}
//	}
//
//	func (c Column) String() string {
//		s := "Column{\n"
//		for _, view := range c.child {
//			s += "\t"
//			childString := view.String()
//			s += strings.Replace(childString, "\n", "\n\t", -1)
//			s += "\n"
//		}
//		s += "}"
//		return s
//	}
//
//	func (c *Column) Init() {
//		c.builder(c.ctx)
//		for _, view := range c.child {
//			view.Init()
//		}
//	}
//
//	func (v *Column) ToCanvasObject() fyne.CanvasObject {
//		c := container.New(layout.NewVBoxLayout())
//		c.Add(layout.NewSpacer())
//		for _, view := range v.child {
//			c.Add(view.ToCanvasObject())
//		}
//		c.Add(layout.NewSpacer())
//		return c
//	}
//
//	func NewColumn(ctx context.Context) *Column {
//		c := Column{}
//		Child(ctx, &c)
//		cv := context.WithValue(ctx, consts.ParentKey, &c)
//		c.ctx = cv
//		return &c
//	}

type DefaultColumn struct {
	DefaultView
	builder func(ctx context.Context)
}

func (c *DefaultColumn) Build(builder func(ctx context.Context)) render.Column {
	c.builder = builder
	return c
}

type FyneColumn struct {
	DefaultColumn
}

func (f *FyneColumn) Render(ctx context.Context) {
	if f.builder == nil {
		panic("column builder is nil")
	}
	f.builder(f.ctx)
	c := ctx.Value(consts.ContextKeyRenderParentObject).(*fyne.Container)
	box := container.New(layout.NewVBoxLayout())
	box.Add(layout.NewSpacer())
	for _, view := range f.GetChildViews() {
		ctx := context.WithValue(ctx, consts.ContextKeyRenderParentObject, box)
		view.Render(ctx)
	}
	box.Add(layout.NewSpacer())
	c.Add(box)
}
func NewColumn(ctx context.Context) render.Column {
	renderer := render.GetApplicationRenderer(ctx)
	switch renderer {
	case consts.RendererFyne:
		return newFyneColumn(ctx)
	default:
		panic("not support")
	}
}
func newFyneColumn(ctx context.Context) render.Column {
	c := FyneColumn{}
	parent := render.GetParentView(ctx)
	parent.Child(&c)
	cv := context.WithValue(ctx, consts.ContextKeyParentView, &c)
	c.ctx = cv
	return &c
}
