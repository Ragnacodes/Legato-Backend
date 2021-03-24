package scenario

// Type interface specifies of Event or Handler
type Type interface {
	Type() string
	Name() string
}

// Event includes methods to trigger custom actions.
// Execute runs the actions in the main thread.
// Post runs the actions in the background thread.
type Event interface {
	Type
	Execute()
	Post()
}

// HandlerListener includes methods for watching and trigger.
// Observe checks the condition to trigger repeatedly.
// It can be for loop or custom time scheduling.
// Trigger uses the events and execute/post the one by one.
type Handler interface {
	Type
	Observe(events []Event)
	Trigger(events []Event)
}

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Events is an array of Event that includes actions.
// Handler controls the scenario actions and events.
type Scenario struct {
	Name    string
	Handler Handler
	Events  []Event
}

// To Start scenario
func (s *Scenario) Start() {
	s.Handler.Observe(s.Events)
}
