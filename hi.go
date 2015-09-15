package main

import (
	"net/http"

	"github.com/insionng/vodka"
	m "github.com/insionng/vodka/middleware"
)

// Handler
func HiHandler(self *vodka.Context) error {
    return self.String(http.StatusOK, "Hi, World!")
}

func main() {

 v := vodka.New()
 v.Use(m.Logger())
 v.Use(m.Recover())
 v.Use(m.Gzip())
 
 v.Any("/", HiHandler)

 // Start server
 v.Run(9000)
 //v.Run(":9000")
 //v.Run("localhost")
 
}
