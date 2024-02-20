package render

import "context"

type Button interface {
	View
	OnClick(click func())
}
type Column interface {
	View
	Build(builder func(ctx context.Context)) Column
}
type Text interface {
	View
	GetText() string
	Text(text string) Text
}
type Row interface {
	View
	Build(builder func(ctx context.Context)) Row
}
type View interface {
	GetContext() context.Context
	Render(ctx context.Context)
	Update(ctx context.Context)
	GetChildViews() []View
	Child(view View)
}

type WindowOption interface {
	Apply(w Window)
}
type Application interface {
	NewWindow(options ...WindowOption) Window
	Run()
}
type Window interface {
	View
	Title(title string) Window
	Size(width, height int) Window
	ShowAndRun()
}
