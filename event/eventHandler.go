package event

// EventHandler Event Should Implements by other handler
type EventHandler interface {
	ProcessEvent(msg *Massage)
}
