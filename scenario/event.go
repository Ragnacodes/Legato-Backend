package scenario

type Action interface {
	Execute()
	Post()
}

type Event struct {
	Action
	Name string
}
