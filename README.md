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
    type func(...interface{})

    //There is the function of event defined by yourself. 
    func event1(st ...interface{}) {
        str := ""
        for _, value := range st {
            val, _ := value.(string)
            str += string(val)
        }
        fmt.Println("event1 emit", str)
    }
```
    
    
## Example of serial events

    func main() {
        //beego.Run()

        var2 := "132342333322222222"
        events := goevents.Classic()

        events.Bind("123", "dfd").On("a", event1)

        evnets.Bind("e2","do some things")
        events.On("b", event2:=func(...interface{}){
            //event2 that you need to define by yourself
        })

        //Can trigger the events that was named belong to On function.
        //events.Bind("e2", "do some things by bind").Trigger("b")
        
        //Trigger all the event functions.
        events.Emit()
    }
    
## Example of parallel event supported

      events.GoOn(coEvent1, "p1")
      events.GoOn(coEvent1, "p2")
      events.GoOn(coEvent1, "p3")
      events.GoOn(coEvent1, "p4")
      events.GoOn(coEvent1, "p5")

      //First trigger the events named by B
      events.Bind(var2).Trigger("b")

      //trigger all of the events that has not been emited.
      events.Emit()

## Example mutil

    package main

    import (
        "fmt"
        "github.com/slclub/goevents"
    )

    func event1(st ...interface{}) {
        str := ""
        for _, value := range st {
            val, _ := value.(string)
            str += string(val)
        }
        fmt.Println("event1 emit", str)
    }

    func event2(st ...interface{}) {
        str := ""
        for _, value := range st {
            val, _ := value.(string)
            str += string(val)
        }
        fmt.Println("event2 emit", str)
    }

    func coEvent1(st ...interface{}) {
        str := ""
        for _, value := range st {
            val, _ := value.(string)
            str += string(val)
        }
        fmt.Println("con event emit", str)
    }

    func main() {
        //beego.Run()

        var2 := "132342333322222222"
        events := goevents.Classic()

        events.Bind("123", "dfd").On("a", event1)
        events.On("b", event2)

        //Bind concurrent events
        events.GoOn(coEvent1, "p1")
        events.GoOn(coEvent1, "p2")
        events.GoOn(coEvent1, "p3")
        events.GoOn(coEvent1, "p4")
        events.GoOn(coEvent1, "p5")

        events.Bind(var2).Trigger("b")

        //Emit all
        events.Emit()
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
