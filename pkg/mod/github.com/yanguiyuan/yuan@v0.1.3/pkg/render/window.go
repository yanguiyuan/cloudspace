package render

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/yanguiyuan/yuan/pkg/render/consts"
)

type FyneWindow struct {
	ctx    context.Context
	title  string
	handle fyne.Window
	views  []View
}
type OpenGLWindow struct {
	ctx    context.Context
	handle *glfw.Window
	title  string
}

func (o *OpenGLWindow) GetContext() context.Context {
	return o.ctx
}

func (o *OpenGLWindow) Render(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (o *OpenGLWindow) Update(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (o *OpenGLWindow) GetChildViews() []View {
	//TODO implement me
	panic("implement me")
}

func (o *OpenGLWindow) Child(view View) {
	//TODO implement me
	panic("implement me")
}

func (o *OpenGLWindow) Title(title string) Window {
	o.title = title
	if o.handle != nil {
		o.handle.SetTitle(title)
	}
	return o
}

func (o *OpenGLWindow) Size(width, height int) Window {
	if o.handle != nil {
		o.handle.SetSize(width, height)
	}
	return o
}

func (o *OpenGLWindow) ShowAndRun() {
	gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	for !o.handle.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		o.handle.SwapBuffers()
		glfw.PollEvents()
	}
}

func (w *FyneWindow) Title(title string) Window {
	w.title = title
	if w.handle != nil {
		w.handle.SetTitle(title)
	}
	return w
}
func (w *FyneWindow) Size(width, height int) Window {
	if w.handle != nil {
		w.handle.Resize(fyne.NewSize(float32(width), float32(height)))
	}
	return w
}
func (w *FyneWindow) GetContext() context.Context {
	return w.ctx
}
func (w *FyneWindow) Child(view View) {
	w.views = append(w.views, view)
}
func (w *FyneWindow) Render(ctx context.Context) {
	c := container.New(layout.NewCenterLayout())
	for _, v := range w.views {
		ctx0 := context.WithValue(ctx, consts.ContextKeyRenderParentObject, c)
		v.Render(ctx0)
	}
	w.handle.SetContent(c)
}
func (w *FyneWindow) Update(ctx context.Context) {

}
func (w *FyneWindow) GetChildViews() []View {
	return w.views
}
func (w *FyneWindow) ShowAndRun() {
	w.Render(w.ctx)
	w.handle.ShowAndRun()
}
