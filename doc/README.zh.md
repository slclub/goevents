# 概述

事件package。 
我们可以自由灵活的使用事件，支持串行事件，并行事件；
且可以将串行事件 按定义的模块去触发.
事件函数的自定义参数无限制，但没有返回值。具体执行事件灵活自定义，用On 类函数接口注入到goevents。
可以灵活的绑定事件队列的参数。针对事件去绑定，或者，统一触发事件之前去绑定，或者混合着使用。单个绑定的优先级会更高一些

###  <a href="#api">API</a>
###  ![English Doc](/slclub/goevents)

# 获取代码到本地

    go get github.com/slclub/goevents
  
# 使用案例

### 串行化事件使用案例

    import "github.com/slclub/goevents"

    ev1 := func(args ...interface{}){
        //ev1 do something.
    }
    ev2 := func(args ...interface{}){
        //ev2 do something.
    }

    func main() {
        events := goevents.Classic()
        //连贯写法 绑定事件及参数
        events.Bind("123", "dfd").On("a", ev1)

        //绑定事件和参数分开写
        events.Bind("event2 serial running")
        events.On("a",ev2)

        //Trigger 
        //events.Bind(...)//可以在这里绑定全局参数，当有事件绑定参数，这里可以重新绑定，只是不能针对所有串行事件
        //按模块名 执行事件队列。无参数 执行所有串行 事件
        events.Trigger("a")
        //触发所有事件，串行，并行的全部触发
        //events.Emit()
    }
  
### 并行执行事件使用案例

这里我们直接写main函数即可，害是用 上面的2个事件就可以了,并行执行因为无序，所以直接简化api不做模块化等。

    func main() {
        events := goevents.Classic()

        //绑定并发事件
        events.GoOn(ev1, "123",344,"event1 param")
        events.GoOn(ev2, "event2 param",454, 343,999)

        //最终执行事件。
        //events.End(func(...interface{}){}, "123","End 事件 最后所有事件执行完 才会执行")
        //触发并发事件不能用Trigger
        events.Emit()
    }

# ![API 文档说明](#api)
