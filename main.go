package main

import (
	"os"
	//    "encoding/json"
	//    "fmt"
	//    "strings"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
	//	"github.com/kataras/iris/hero"
)

type ENV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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

	app.Get("/get/{name}", func(ctx iris.Context) { // get an ENV variable
		name := ctx.Params().Get("name")

		viper.SetConfigName("config") // name of config file (without extension)
		viper.AddConfigPath(".")      // optionally look for config in the working directory
		if err := viper.ReadInConfig(); err != nil {
			ctx.WriteString(err.Error())
			return
		} // Find and read the config file

		key := viper.GetString("name")

		if len(key) == 0 {
			ctx.Writef("Variable not set")
		} else {
			ctx.Writef("%s=%s", name, key)
		}
	})

	app.Get("/delete/{name}", func(ctx iris.Context) { // delete an ENV variable
		name := ctx.Params().Get("name")

		viper.SetConfigName("config") // name of config file (without extension)
		viper.AddConfigPath(".")
		viper.Set(name, "")
		ctx.Writef("%s has been removed", name)
	})

	app.Post("/set", func(ctx iris.Context) { // set an ENV variable
		var kv ENV

		if err := ctx.ReadJSON(&kv); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		} else {
			viper.SetConfigName("config") // name of config file (without extension)
			viper.AddConfigPath(".")
			viper.Set(kv.Key, kv.Value)
			ctx.Writef("Variable has been set. \n key: %s \n value: %s \n", kv.Key, kv.Value)
		}
	})

	// listen and serve on http://0.0.0.0:<port>.
	app.Run(iris.Addr(":" + port))
}
