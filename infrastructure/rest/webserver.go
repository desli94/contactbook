package rest

import (
	"fmt"
	"log"

	"github.com/desli94/contactbook/registry"
	"github.com/labstack/echo/v4"
)

// CreateServer creates a server instance
func CreateServer(r registry.Registry) {
	e := echo.New()

	e = NewContactRoutes(e, r.NewContactController())

	fmt.Println("Server started at port: ", 8080)
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
