package scenario

type Handler interface {
	Observe()
	Trigger()
}
