package services

// Service contains details about provided service.
// Execute runs the related action in the main thread.
// Post runs the related actions in the background thread.
// Next runs the next node(s)
type Service interface {
	Execute(*Pipe)
	Post(*Pipe)
	Resume(data ...interface{})
	Next(*Pipe)
}
