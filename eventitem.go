package goevents

//Event interface
type EventItem interface {
	exec(args ...interface{})
}

type Event interface {
}

//The struct of event
type eventItem struct {
	name  string
	fn    func(...interface{})
	param []interface{}
	//Whether the event has run
	emited bool
}

//Create a new event
func NewEvent() *eventItem {
	fn := func(args ...interface{}) {}
	param := make([]interface{}, 0)
	return &eventItem{"", fn, param, false}
}

//Excute the current event
func (this *eventItem) exec(args ...interface{}) {
	if this.emited {
		return
	}
	this.fn(args...)
	this.emited = true
}
