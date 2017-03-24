package goevents

import "reflect"

//Event interface
type EventItem interface {
	exec(args ...Arguments)
}

type Event interface {
}

//The struct of event
type eventItem struct {
	fn    EventFunc
	param []Arguments
	//Whether the event has run
	emited bool
	len    int
}

//Create a new event
func NewEvent(fn EventFunc, param []Arguments) *eventItem {
	l := reflect.TypeOf(fn).NumIn()
	return &eventItem{fn, param, false, l}
}

//Excute the current event
func (this *eventItem) exec(args ...Arguments) {
	if this.emited {
		return
	}
	argvs := getArgs(args...)
	if len(argvs) > this.len {
		argvs = argvs[:this.len]
	}
	reflect.ValueOf(this.fn).Call(argvs)
	this.emited = true
}

//Convert  type of ...Arguments to reflect.Value slice type
func getArgs(args ...Arguments) []reflect.Value {
	var ma = make([]reflect.Value, len(args))
	for k, v := range args {
		ma[k] = reflect.ValueOf(v)
	}
	return ma
}

//Init function
func initFn() {
}
