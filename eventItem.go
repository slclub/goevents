package goevents

type EventItem interface {
	exec(args ...interface{})
}

type Event interface {
}

type eventItem struct {
	name     string
	fn       func(...interface{})
	param    []interface{}
	paralled bool
}

func NewEvent() *eventItem {
	return &eventItem{}
}

func (this *eventItem) exec(args ...interface{}) {
	this.fn(args...)
}
