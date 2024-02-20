package ui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/yanguiyuan/yuan/pkg/render"
	"github.com/yanguiyuan/yuan/pkg/render/consts"
)

//	type Row struct {
//		DefaultView
//		builder func(ctx context.Context)
//	}
//
//	func (r *Row) Build(builder func(ctx context.Context)) *Option {
//		//TODO implement me
//		r.builder = builder
//		return &Option{}
//	}
//
//	func (r Row) String() string {
//		s := "Row{\n"
//		for _, view := range r.child {
//			s += "\t"
//			childString := view.String()
//			s += strings.Replace(childString, "\n", "\n\t", -1)
//			s += "\n"
//		}
//		s += "}"
//		return s
//	}
//
//	func (r *Row) Init() {
//		r.builder(r.ctx)
//		for _, view := range r.child {
//			view.Init()
//		}
//	}
//
//	func (r *Row) ToCanvasObject() fyne.CanvasObject {
//		c := container.New(layout.NewHBoxLayout())
//		c.Add(layout.NewSpacer())
//		for _, view := range r.child {
//			c.Add(view.ToCanvasObject())
//		}
//		c.Add(layout.NewSpacer())
//		return c
//	}
//
//	func NewRow(ctx context.Context) View {
//		c := Row{}
//		Child(ctx, &c)
//		cv := context.WithValue(ctx, consts.ParentKey, &c)
//		c.ctx = cv
//		return &c
//	}

type DefaultRow struct {
	DefaultView
	builder func(ctx context.Context)
}

func (r *DefaultRow) Build(builder func(ctx context.Context)) render.Row {
	r.builder = builder
	return r
}

type FyneRow struct {
	DefaultRow
}

func (f *FyneRow) Render(ctx context.Context) {
	if f.builder == nil {
		panic("column builder is nil")
	}
	f.builder(f.ctx)
	c := ctx.Value(consts.ContextKeyRenderParentObject).(*fyne.Container)
	box := container.New(layout.NewHBoxLayout())
	box.Add(layout.NewSpacer())
	for _, view := range f.GetChildViews() {
		ctx := context.WithValue(ctx, consts.ContextKeyRenderParentObject, box)
		view.Render(ctx)
	}
	box.Add(layout.NewSpacer())
	c.Add(box)
}
func NewRow(ctx context.Context) render.Row {
	renderer := render.GetApplicationRenderer(ctx)
	switch renderer {
	case consts.RendererFyne:
		return newFyneRow(ctx)
	default:
		return nil
	}
}
func newFyneRow(ctx context.Context) render.Row {
	c := FyneRow{}
	cv := context.WithValue(ctx, consts.ContextKeyParentView, &c)
	parent := render.GetParentView(ctx)
	parent.Child(&c)
	c.ctx = cv
	return &c
}
