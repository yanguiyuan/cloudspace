package consts

type contextKey int

const (
	ContextKeyApplication contextKey = iota
	ContextKeyWindow
	ContextKeyParentView
	ContextKeyRenderer
	ContextKeyRenderParentObject
)
