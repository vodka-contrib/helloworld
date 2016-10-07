package main

import (
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/fasthttp"
	m "github.com/insionng/vodka/middleware"
)

func main() {
	v := vodka.New()
	v.Use(m.Logger())
	v.Use(m.Recover())
	v.Use(m.Gzip())

	v.GET("/", HiHandler)
	v.Run(fasthttp.New(":1987"))
}

func HiHandler(self vodka.Context) error {
	return self.String(http.StatusOK, "Hi, World!")
}
