<link rel="stylesheet" href="markdown.css">

# Event description

Event loop

Serial and parallel events supported.

# Get code to your local development.

    go get github.com/slclub/goevents

# Usage


## first import the package
  
    import github.com/slclub/goevents
  
## Instance
  
    events := goevents.Classic()


## Defined your self event func


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
    
    
## Example of serial events

    func main() {
    ¦ //beego.Run()

    ¦ var2 := "132342333322222222"
    ¦ events := goevents.Classic()

    ¦ events.Bind("123", "dfd").On("a", event1)

    ¦ evnets.Bind("e2","do some things")
      events.On("b", event2)//event2 that you need to define by yourself

      //Can trigger the events that was named belong to On function.
      //Events.Bind("e2", "do some things by bind").Trigger("b")
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
      //"github.com/astaxie/beego"
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
