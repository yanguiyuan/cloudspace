package ui

import (
	"context"
	"github.com/yanguiyuan/yuan/pkg/render"
)

//	type DefaultView struct {
//		options Option
//		ctx     context.Context
//		click   func()
//		child   []View
//	}
//
// type Event struct {
// }
//
//	type View interface {
//		fmt.Stringer
//		Build(builder func(ctx context.Context)) *Option
//		Init()
//		Child(view View)
//		Onclick(func())
//		ToCanvasObject() fyne.CanvasObject
//	}
//
//	func (v *DefaultView) Build(builder func(ctx context.Context)) *Option {
//		return &Option{}
//	}
//
//	func (v *DefaultView) Child(view View) {
//		v.child = append(v.child, view)
//	}
//
//	func (v *DefaultView) Onclick(click func()) {
//		v.click = click
//	}
//
//	func (v DefaultView) String() string {
//		return "DefaultView"
//	}
//
//	func (v *DefaultView) Init() {
//		for _, view := range v.child {
//			view.Init()
//		}
//	}
//
//	func (v *DefaultView) ToCanvasObject() fyne.CanvasObject {
//		c := container.New(layout.NewCenterLayout())
//		for _, view := range v.child {
//			c.Add(view.ToCanvasObject())
//		}
//		return c
//	}
//
//	func (d DefaultView) GetContext() context.Context {
//		return d.ctx
//	}
//
//	func (d DefaultView) GetChildViews() []View {
//		return d.child
//	}
//
//	func (d *DefaultView) SetContext(ctx context.Context) {
//		d.ctx = ctx
//	}

type ViewStyle interface {
	Width(w int) ViewStyle
	Height(h int) ViewStyle
}
type DefaultView struct {
	ctx   context.Context
	child []render.View
}

func (v *DefaultView) GetContext() context.Context {
	return v.ctx
}
func (v *DefaultView) Render(ctx context.Context) {

}
func (v *DefaultView) Update(ctx context.Context) {

}
func (v *DefaultView) GetChildViews() []render.View {
	return v.child
}
func (v *DefaultView) Child(view render.View) {
	v.child = append(v.child, view)
}
