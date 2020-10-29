package main

import (
	"net/http"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type Proxy struct {
	App *iris.Application
}

const apiStub = "80"

func main() {
	p := &Proxy{}
	app, err := p.initApp()
	p.App = app
	srv := &http.Server{
		Addr:         apiStub,
		ReadTimeOut:  time.Duration(1),
		WriteTimeOut: time.Duration(1),
	}
	runHandlers := []iris.Configurator{
		// disables interrupt handler
		iris.WithoutInterruptHandler,
		// disables updates
		iris.WithoutVersionChecker,
		// enables faster json serialization and more
		iris.WithOptimizations,
		iris.WithoutBodyConsumptionOnUnmarshal,
	}

	app.Run(iris.Server(srv), runHandlers...)
}

func (p *Proxy) initApp() *iris.Application {
	app := iris.New()
	app.Partyt("/chart").Handle(new(Controller))
	return app
}

type Controller struct {
	Ctx iris.Context
}

func (c *Controller) BeforeActivation(b mvc.BeforeActivation) {
	c.Handle("POST", "/enable", "Enable")
}

func (c *Controller) Enable() {
	c.Ctx.JSON("Hello World")
}

/*
1 åå§åappå®ä¾
2 å è½½ä¸­é´ä»¶
3 ç±ä¸­é´ä»¶æå»ºmvcå®ä¾
4 mvcå®ä¾handleä¸ä¸ªåå«å¤ä¸ªå¤çrestæ¹æ³çcontroller
5 app.Run()æå»ºå®ä¾ï¼å¯å¨æå¡ï¼é»å¡ç­å¾è¿æ¥

*/

package main

import (
	"errors"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		fmt.Println("test1")
		ctx.Next()
		fmt.Println("test2")
	})

	mvc.New(app).Handle(new(ExampleController))
	mvc.New(app).Handle(new(ExampleController1))
	return app
}

type ExampleController struct{}
type ExampleController1 struct{}

func (e *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		fmt.Println("middleware 1")
		ctx.Next()
		fmt.Println("middleware 2")
	}
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)
}

func (e *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

func (e *ExampleController) GetPing() string {
	fmt.Println("pong")
	return "pong"
}

/*
func (e *ExampleController1) GetPing() string {
	return "pong controller 1"
}
*/

func (e *ExampleController) PostPing() string {
	return "pong post"
}

func (e *ExampleController) PostPanic() {
	panic(errors.New("test"))
}

func (e *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (e *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	fmt.Println("guide")
	return "hello from the custom handler without following the naming guide"
}
[kouliping@A01-R04-I17-130-06CEKFK t_code]$ cat iris/main.go 
/*
1 åå§åappå®ä¾
2 å è½½ä¸­é´ä»¶
3 ç±ä¸­é´ä»¶æå»ºmvcå®ä¾
4 mvcå®ä¾handleä¸ä¸ªåå«å¤ä¸ªå¤çrestæ¹æ³çcontroller
5 app.Run()æå»ºå®ä¾ï¼å¯å¨æå¡ï¼é»å¡ç­å¾è¿æ¥

*/

package main

import (
	"errors"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		fmt.Println("test1")
		ctx.Next()
		fmt.Println("test2")
	})

	mvc.New(app).Handle(new(ExampleController))
	mvc.New(app).Handle(new(ExampleController1))
	return app
}

type ExampleController struct{}
type ExampleController1 struct{}

func (e *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		fmt.Println("middleware 1")
		ctx.Next()
		fmt.Println("middleware 2")
	}
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)
}

func (e *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

func (e *ExampleController) GetPing() string {
	fmt.Println("pong")
	return "pong"
}

/*
func (e *ExampleController1) GetPing() string {
	return "pong controller 1"
}
*/

func (e *ExampleController) PostPing() string {
	return "pong post"
}

func (e *ExampleController) PostPanic() {
	panic(errors.New("test"))
}

func (e *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (e *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	fmt.Println("guide")
	return "hello from the custom handler without following the naming guide"
}

