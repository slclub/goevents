package goevents

import "sync"

//Events interface who operate the events struct
//On:bind event to the queue of events
//GoOn:bind event to the queue and run run in gorountine
//Emit all the event in the queue
type Events interface {
	On(on string, fn func(interface{}) error)
	GoOn(on string, fn func(interface{}) error)
	Emit()
	//Can run the special event defined based on the first arguments
	Trigger()
}

//goevents master struct
type events struct {
	*sync.RWMutex
	//Structured event queue
	queue map[string][]*eventItem
	//Liner struct event queue
	loop []*eventItem

	// Temp arguments for the curruent event
	curParam []interface{}
	// Current event
	curEvent *eventItem
}

//All the events instances
//It is the entry
func Classic() (this *events) {
	this = new(events)
	this.queue = make(map[string][]*eventItem)
	this.loop = make([]*eventItem, 0)
	return
}

//Bind event to the attribute queue of events struct
func (this *events) On(name string, fn func(...interface{}), args ...interface{}) {
	item := NewEvent()
	item.fn = fn

	if len(name) == 0 {
		name = "all"
	}

	if len(args) == 0 {
		args = this.curParam
	}
	item.param = args

	this.curEvent = item
	this.queue[name] = append(this.queue[name], item)
	this.loop = append(this.loop, item)
	this.curParam = make([]interface{}, 0)

	return
}

/**
 * Bind param to the variable curparam
 * Return events master struct
 * It would be clear the variable curParam when invoke On function.
 */
func (this *events) Bind(args ...interface{}) *events {
	this.curParam = args
	return this
}

/**
 * Tiigger events
 * If the lenght of names bigger than one It will trigger the group events that named.
 */
func (this *events) Trigger(names ...string) {

	if len(names) > 0 && len(this.queue[names[0]]) == 0 {
		return
	}

	loop := this.loop
	if len(names) > 0 && len(this.queue[names[0]]) > 0 {
		loop = this.queue[names[0]]
	}

	for _, e := range loop {
		param := e.param
		if len(param) == 0 {
			param = this.curParam
		}
		e.exec(param...)
	}
}

// Trigger all the events
func (this *events) Emit() {
	this.Trigger()
}

func (this *events) GoOn() {
}
