package core

type Entity interface {
	Startup()
	Update()
	OnDestroy()
	Destroy()
	Dead() bool
}
type EmptyEntity struct {
	dead bool
}

func (e *EmptyEntity) Dead() bool {
	return e.dead
}
func (e *EmptyEntity) Destroy() {
	e.dead = true
}
func (e *EmptyEntity) OnDestroy() {

}
func (e *EmptyEntity) Update() {

}
func (e *EmptyEntity) Startup() {

}
