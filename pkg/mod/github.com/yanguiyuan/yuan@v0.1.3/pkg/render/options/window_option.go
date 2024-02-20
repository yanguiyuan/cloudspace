package options

import "github.com/yanguiyuan/yuan/pkg/render"

type WindowTitleOption struct {
	title string
}

func (o WindowTitleOption) Apply(w render.Window) {
	w.Title(o.title)
}
func WithTitle(title string) render.WindowOption {
	return WindowTitleOption{title: title}
}
