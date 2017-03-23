package goevents

import "fmt"

type Concurrent interface {
	gofunc(args ...interface{})
	wait()
	end(fn func(args ...interface{}), args ...interface{})
	emit()
}

type concurrent struct {
	*channelManager
	//Concurrent events queue
	loop []*eventItem
	//need wait gorountine
	waited       bool
	endFn        *eventItem
	currentParam []interface{}
}

//
type channelManager struct {
	chNumber int
	ch       chan int
}

var p = fmt.Println

func NewConcurrent(chNum int) *concurrent {
	loop := make([]*eventItem, 0)
	endFn := NewEvent()
	cur := make([]interface{}, 0)
	return &concurrent{NewChannelManager(chNum), loop, true, endFn, cur}
}

func NewChannelManager(chNum int) *channelManager {
	ch := make(chan int, chNum)
	return &channelManager{chNum, ch}
}

//bind event to concurrent queue.
func (this *concurrent) on(fn func(args ...interface{}), args ...interface{}) *concurrent {
	if fn == nil {
		return this
	}

	item := NewEvent()
	item.fn = fn

	item.param = args

	this.loop = append(this.loop, item)
	this.currentParam = args
	return this
}

//Concurrent run the events that has been defined
func (this *concurrent) gofunc(fn func(args ...interface{}), args ...interface{}) {

	//Set argument by current param If lenght of args equal zero.
	if len(args) == 0 {
		args = this.currentParam
	}

	fn(args...)

	this.ch <- 1
}

//Wait all the goruountine finished.
//If not finished will wait here.
func (this *concurrent) wait() {
	len := len(this.loop)
	for i := 0; i < len; i++ {
		<-this.ch
	}
}

//Add the last event function called.
func (this *concurrent) end(fn func(args ...interface{}), args ...interface{}) {
	if fn == nil {
		return
	}

	this.endFn.fn = fn
	this.endFn.param = args
}

//Emit all concurrent events
func (this *concurrent) emit() error {

	if len(this.loop) == 0 {
		return nil
	}

	//invoke the events that was in the queue
	for _, e := range this.loop {
		param := e.param
		go this.gofunc(e.fn, param...)
	}

	if this.waited {
		this.wait()
	}

	//p(this.chNumber)

	//running the last event
	params := this.endFn.param
	if len(params) == 0 {
		this.endFn.fn()
	} else {
		this.endFn.fn(params...)
	}

	return nil
}
