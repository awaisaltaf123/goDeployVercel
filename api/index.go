
package handler

import (
	"fmt"
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/api/test/golang", func(context *Context) {
		context.JSON(200, H{
			"message": "Get precor Data Successfully from Golang",
		})
	})

	server.Handle(w, r)
}
