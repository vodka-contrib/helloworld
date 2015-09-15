package main

import (
	"fmt"
	"net/http"

	"github.com/insionng/vodka"
	m "github.com/insionng/vodka/middleware"
)

func main() {

	v := vodka.New()
	v.Use(m.Logger())
	v.Use(m.Recover())
	v.Use(m.Gzip())

// Handler
func HiHandler(self *vodka.Context) error {
    return self.String(http.StatusOK, "Hi, World!")
}

 v.Any("/", HiHandler)

// Start server
 v.Run(9000)
 //v.Run(":9000")
 //v.Run("localhost")
}
