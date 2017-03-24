package goevents

import "reflect"

type Concurrent interface {
	gofunc(args ...Arguments)
	wait()
	end(fn EventFunc, args ...Arguments)
	emit()
}

type concurrent struct {
	*channelManager
	//Concurrent events queue
	loop []*eventItem
	//need wait gorountine
	waited       bool
	endFn        *eventItem
	currentParam []Arguments
}

//
type channelManager struct {
	chNumber int
	ch       chan int
}

func NewConcurrent(chNum int) *concurrent {
	loop := make([]*eventItem, 0)
	cur := make([]Arguments, 0)
	endFn := NewEvent(initFn, cur)
	return &concurrent{NewChannelManager(chNum), loop, true, endFn, cur}
}

func NewChannelManager(chNum int) *channelManager {
	ch := make(chan int, chNum)
	return &channelManager{chNum, ch}
}

//bind event to concurrent queue.
func (this *concurrent) on(fn EventFunc, args ...Arguments) *concurrent {
	if fn == nil {
		return this
	}

	item := NewEvent(fn, args)

	this.loop = append(this.loop, item)
	this.currentParam = args
	return this
}

//Concurrent run the events that has been defined
func (this *concurrent) gofunc(item *eventItem) {

	//Set argument by current param If lenght of args equal zero.
	if len(item.param) == 0 {
		item.param = this.currentParam
	}
	argvs := getArgs(item.param...)
	reflect.ValueOf(item.fn).Call(argvs[:item.len])

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
func (this *concurrent) end(fn EventFunc, args ...Arguments) {
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
		go this.gofunc(e)
	}

	if this.waited {
		this.wait()
	}

	//p(this.chNumber)

	//running the last event
	params := this.endFn.param

	fn, ok := this.endFn.fn.(eventFunc)
	if !ok {
		return nil
	}
	if len(params) == 0 {
		fn()
	} else {
		fn(params...)
	}

	return nil
}
