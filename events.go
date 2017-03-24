package goevents

import "sync"
import "fmt"

var print = fmt.Println

type Arguments interface{}

type eventFunc func(args ...Arguments)
type EventFunc interface{}

//Events interface who operate the events struct
//On:bind event to the queue of events
//Emit all the event in the queue
type Events interface {
	On(on string, fn func(...Arguments) error)
	Emit()
	//Can run the special event defined based on the first Arguments
	Trigger()
}

//Open parallel interface
//GoOn:bind event to the queue and run run in gorountine
//End:bind the last event
type EventsConcurrent interface {
	GoOn(fn EventFunc, args ...Arguments)
	End(fn EventFunc, args ...Arguments)
}

//goevents master struct
type events struct {
	*sync.RWMutex
	//Structured event queue
	queue map[string][]*eventItem
	//Liner struct event queue
	loop []*eventItem

	// Temp Arguments for the curruent event
	curParam []Arguments
	// Current event
	curEvent *eventItem
	//eventsModule is running the evnets that added by devoloper
	running bool
	//concurrent running object
	concurrent *concurrent
	config     *config
}

//All the events instances
//It is the entry
func Classic() (this *events) {
	this = new(events)
	this.queue = make(map[string][]*eventItem)
	this.loop = make([]*eventItem, 0)
	this.running = false
	this.config = newConf(0, 0, false)
	this.concurrent = NewConcurrent(this.config.chNumber)
	return
}

//Bind event to the attribute queue of events struct
func (this *events) On(name string, fn EventFunc, args ...Arguments) {
	if fn == nil {
		return
	}

	if len(name) == 0 {
		name = "all"
	}

	if len(args) == 0 {
		args = this.curParam
	}
	item := NewEvent(fn, args)
	this.curEvent = item
	this.queue[name] = append(this.queue[name], item)
	this.loop = append(this.loop, item)
	this.curParam = make([]Arguments, 0)

	return
}

/**
 * Bind param to the variable curparam
 * Return events master struct
 * It would be clear the variable curParam when invoke On function.
 */
func (this *events) Bind(args ...Arguments) *events {
	this.curParam = args
	this.concurrent.currentParam = args
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

	this.running = true
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
	this.concurrent.emit()
}

// Add concurrent event
func (this *events) GoOn(fn EventFunc, args ...Arguments) error {
	this.curParam = args
	this.concurrent.on(fn, args...)
	return nil
}

// Add the last event
func (this *events) End(fn EventFunc, args ...Arguments) {
	this.concurrent.end(fn, args...)
}
