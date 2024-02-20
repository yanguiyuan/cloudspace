package ui_test

//
//import (
//	"context"
//	"fmt"
//	"github.com/yanguiyuan/yuan/pkg/render/app"
//	ui2 "github.com/yanguiyuan/yuan/pkg/render/ui"
//	"testing"
//)
//
//func TestUI(t *testing.T) {
//	application := app.NewApplication()
//	w := application.NewWindow()
//	w.SetTitle("Test Window")
//	ctx := w.GetWindowConText()
//	ui2.NewColumn(ctx).Build(func(ctx context.Context) {
//		ui2.NewButton(ctx, "Hello")
//		ui2.NewRow(ctx).Build(func(ctx context.Context) {
//			ui2.NewText(ctx, "Hello,World")
//			ui2.NewButton(ctx, "ClickMe").OnClick(func() {
//				fmt.Println("ClickMe")
//			})
//		})
//	})
//	application.Run()
//	fmt.Println(application)
//}
//func TestContext(t *testing.T) {
//	ctx := context.Background()
//	ctx1 := context.WithValue(ctx, "app", "1")
//	ctx2 := context.WithValue(ctx1, "app", "2")
//	fmt.Println(ctx2.Value("app"))
//}
