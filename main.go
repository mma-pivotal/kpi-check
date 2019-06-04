package main

import (
    "os"
//  "fmt"
//  "strings"
	"github.com/kataras/iris"
//	"github.com/kataras/iris/hero"
)

func main() {
    app := iris.Default()
    port := os.Getenv("port")

    app.Get("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "Welcome to PCF KPI check",
        })
    })

    os.Setenv("testenv", "abc123")
    app.Get("/get/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        key := os.Getenv(name)

        if len(key) == 0 {
            ctx.Writef("Env variable not set")
        } else {
            ctx.Writef("%s=%s", name, key)
        }
    })

    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":"+port))
}