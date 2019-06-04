package main

import (
    "os"
//    "encoding/json"
//    "fmt"
//    "strings"
	"github.com/kataras/iris"
//	"github.com/kataras/iris/hero"
)

type ENV struct {
    Key string  `json:"key"`
    Value string  `json:"value"`
}

func main() {
    app := iris.Default()
    port := os.Getenv("port")
    if len(port) == 0 {
        port = "8080"
    }
    //read port from ENV, default is 8080

    app.Get("/", func(ctx iris.Context) { // root, may need to add FAQ link here
        ctx.JSON(iris.Map{
            "message": "Welcome to PCF KPI check",
        })
    })

    app.Get("/get/{name}", func(ctx iris.Context) {  // get an ENV variable
        name := ctx.Params().Get("name")
        key := os.Getenv(name)

        if len(key) == 0 {
            ctx.Writef("Env variable not set")
        } else {
            ctx.Writef("%s=%s", name, key)
        }
    })

    app.Post("/set", func(ctx iris.Context) {  // set an ENV variable
        var kv ENV

        if err := ctx.ReadJSON(&kv); err != nil{
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.WriteString(err.Error())
            return
        } else {
            err := os.Setenv(kv.Key, kv.Value)
            if err != nil {
                ctx.Writef("%s",err.Error())
                return 
            }
            ctx.Writef("key: %s \n", kv.Key)
            ctx.Writef("value: %s \n", kv.Value)
        }
    })

    // listen and serve on http://0.0.0.0:<port>.
    app.Run(iris.Addr(":"+port))
}