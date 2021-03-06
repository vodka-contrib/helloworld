package main

import (
	"net/http"

	"github.com/insionng/vodka"
	m "github.com/insionng/vodka/middleware"
	"github.com/insionng/vodka/engine/fasthttp"
)

func main() {
	v := vodka.New()
	v.Use(m.Logger())
	v.Use(m.Recover())
	v.Use(m.Gzip())

	v.GET("/", HelloHandler)
	v.Run(fasthttp.New(":1987"))
}

func HelloHandler(self vodka.Context) error {
	return self.String(http.StatusOK, "Hello, Vodka!")
}
