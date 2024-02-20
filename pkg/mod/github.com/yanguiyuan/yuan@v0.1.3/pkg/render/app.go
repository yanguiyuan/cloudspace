package render

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/yanguiyuan/yuan/pkg/iter"
	"github.com/yanguiyuan/yuan/pkg/render/consts"
	"log"
)

type FyneApplication struct {
	rtx    context.Context
	handle fyne.App
	window []*FyneWindow
}

func (a *FyneApplication) NewWindow(options ...WindowOption) Window {
	w := FyneWindow{}
	ctx1 := context.WithValue(a.rtx, consts.ContextKeyApplication, a)
	ctx2 := context.WithValue(ctx1, consts.ContextKeyWindow, &w)
	ctx3 := context.WithValue(ctx2, consts.ContextKeyParentView, &w)
	w.ctx = ctx3
	for _, option := range options {
		option.Apply(&w)
	}
	if w.title == "" {
		w.title = "Yuan"
	}
	w.handle = a.handle.NewWindow(w.title)
	a.window = append(a.window, &w)
	return &w
}
func (a *FyneApplication) Run() {
	iter.FromSlice(a.window).ForEach(func(w *FyneWindow) {
		w.ShowAndRun()
	})
}

type OpenGLApplication struct {
	rtx    context.Context
	window []*OpenGLWindow
}

func (a *OpenGLApplication) NewWindow(options ...WindowOption) Window {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	w := OpenGLWindow{}
	for _, option := range options {
		option.Apply(&w)
	}
	// 设置窗口参数
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(100, 100, w.title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	w.handle = window
	a.window = append(a.window, &w)

	return &w
}
func (a *OpenGLApplication) Run() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
	iter.FromSlice(a.window).ForEach(func(w *OpenGLWindow) {
		w.ShowAndRun()
	})
	glfw.Terminate()
}
func NewApplication(renderer consts.Renderer) Application {
	switch renderer {
	case consts.RendererFyne:
		return &FyneApplication{
			handle: app.New(),
			rtx:    context.WithValue(context.Background(), consts.ContextKeyRenderer, consts.RendererFyne),
		}
	default:

		return &OpenGLApplication{
			rtx: context.WithValue(context.Background(), consts.ContextKeyRenderer, consts.RendererOpenGL),
		}
	}
}
func GetApplicationRenderer(ctx context.Context) consts.Renderer {
	return ctx.Value(consts.ContextKeyRenderer).(consts.Renderer)
}
func GetParentView(ctx context.Context) View {
	return ctx.Value(consts.ContextKeyParentView).(View)
}
