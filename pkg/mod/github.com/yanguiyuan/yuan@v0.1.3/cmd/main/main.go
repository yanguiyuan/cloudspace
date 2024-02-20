package main

import (
	"context"
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/render"
	"github.com/yanguiyuan/yuan/pkg/render/consts"
	"github.com/yanguiyuan/yuan/pkg/render/options"
	"github.com/yanguiyuan/yuan/pkg/render/ui"
	"github.com/yanguiyuan/yuan/pkg/render/ui/log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}
func main() {
	application := render.NewApplication(consts.RendererFyne)
	w := application.NewWindow(options.WithTitle("TestWindow")).Title("YuanWindow").Size(400, 200)
	ctx := w.GetContext()
	ui.NewColumn(ctx).Build(func(ctx context.Context) {
		ui.NewButton(ctx, "Hello")
		ui.NewRow(ctx).Build(func(ctx context.Context) {
			text := ui.NewText(ctx, "Hello,World")
			ui.NewButton(ctx, "ClickMe").OnClick(func() {
				log.Info(ctx, "ClickMe")
				text.Text("Haha")
			})
		})
	})
	application.Run()
	fmt.Println(application)
}
