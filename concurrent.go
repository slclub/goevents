package goevents

//import "fmt"

const (
	//Default concurrented numbers of channels
	defaultChannelNumber = 10
)

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
	waited bool
	endFn  *eventItem
}

//
type channelManager struct {
	chNumber int
	ch       chan int
}

func NewConcurrent() *concurrent {
	loop := make([]*eventItem, 0)
	return &concurrent{NewChannelManager(), loop, true, nil}
}

func NewChannelManager() *channelManager {
	ch := make(chan int, defaultChannelNumber)
	return &channelManager{defaultChannelNumber, ch}
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
	return this
}

//Concurrent run the events that has been defined
func (this *concurrent) gofunc(fn func(args ...interface{}), args ...interface{}) {
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

	for _, e := range this.loop {
		param := e.param
		go this.gofunc(e.fn, param...)
	}

	if this.waited {
		this.wait()
	}

	return nil
}
