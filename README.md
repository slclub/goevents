<link rel="stylesheet" href="markdown.css">

# Event description

Event loop.

Serial and parallel events supported.

You can use them seperately. Also you can also mix using serial and parallel .

####  [简体中文](https://github.com/slclub/goevents/blob/master/doc/README.zh.md)
####  <a href="https://github.com/slclub/goevents#api">API documents</a>

- <a href="#classic">[Classic]</a>

# Get code to your local development.

    go get github.com/slclub/goevents

# Usage


## first import the package
  
    import github.com/slclub/goevents
  
## Instance
  
    events := goevents.Classic()


## Defined your self event func

```
    //defined func follow the type.It is the beanchmark,don't need to appear in your code
    Define functions "func" rather than method according to golang standards

    //There is the function of event defined by yourself. The following definitions are allowed.Maybe you can say that as long as the golang function is allowed
    
    func event1(st ...interface{}) { }//allowed
   
    func event(){}//allowed.
    
    func event(n int, str string){}//allowed.
    
```
    
    
## Example of serial events

    func main() {

        self := goevents.Classic()

        self.On("a", func() {
        print("serial event added:e1")
        })

        self.On("b", func(n int) {
        print("serial event added:e2; args:", n)
        })
        self.Bind(22222)

        self.Bind("I am string").On("b", func(str string) {
        print("serial event added:e3; args:", str)
        })

        //Can trigger the events that was named belong to On function.
        //events.Bind("e2", "do some things by bind").Trigger("b")
        
        //Trigger all the event functions.
        events.Emit()
    }
    
## Example of parallel event supported

    func main() {

        self := goevents.Classic()
        self.GoOn(func() {
        print("parallel event added e1")
        }, "no", 1)

        self.GoOn(func(str string, n int) {
        print("parallel event added e2", str, n)
        }, "no", 2)

        self.Emit()

        //trigger all of the events that has not been emited.
        self.Emit()

## Example mutil

    package main

    import (
        "fmt"
        "github.com/slclub/goevents"
        "time"
    )

    var print = fmt.Println

    func Test() {
        Timer := time.Now()
        self := goevents.Classic()

        self.On("a", func() {
            print("serial event added:e1")
        })

        self.On("b", func(n int) {
            print("serial event added:e2; args:", n)
        })
        self.Bind(22222)

        self.Bind("I am string").On("b", func(str string) {
            print("serial event added:e3; args:", str)
        })

        self.GoOn(func() {
            print("parallel event added e1")
        }, "no", 1)

        self.GoOn(func(str string, n int) {
            print("parallel event added e2", str, n)
        }, "no", 2)

        self.Emit()
        excuTimes := time.Since(Timer)
        print("event running time:", excuTimes)
    }

# <a name="api">API document</a>

### <a name="classic" >[events.Classic()]</a>

    //instance
    ev := events.Classic()
    
### .On(name string, fn func(...interface{}))

Bind event to the object of events

    ev.On("message", func(args ...interface{}){
        //do some things
    })

### .Bind(args ...interface{})*events

Bind param to the current event func
```
    //use mod one
    ev.Bind("abc",123,&struct1{1,2})
    
    //you can also coherent usage.
    ev.Bind(...).On("message", func(args ...interface{}){
        //do something
    })
```    

### .Trigger(args ...string)

Trigger the events by the first element of the args. it will Emit all events if no argments.
    
    //Just trigger partion of the events by first argment.
    ev.Trigger("message")
    //If no params it will emit all the serial events 
    ev.Trigger()

### .GoOn(fn func(...interface{}), args ...interface{})

Bind event that need to parallel execution.
    
    ev.GoOn(func(...interface{}){
        //event do something
    }, args)
    
    
### .Emit()

Emit all the events.
Parallel events execution included.
    
    ev.Emit().

### .Conf(chNum int, safeMod int)

Setting events object running mod.
Param:chNum parallel gorountine numbers.
Param:safeMod events object running mod.

    ev.Conf(3,0)
    
more contact with slclub
