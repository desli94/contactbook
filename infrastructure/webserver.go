package infrastructure

import (
	"fmt"
	"log"
	"net/http"

	ac "github.com/desli94/contactbook/adapters/contact"
	"github.com/labstack/echo/v4"
)

type contact struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Number    string `json:"number"`
}

// CreateServer creates a server instance
func CreateServer(cc ac.Controller) {
	e := echo.New()

	// db := NewDB()
	// ac.NewContactController(db)

	e.POST("/addContact", func(context echo.Context) error {
		c := new(contact)
		if err := context.Bind(c); err != nil {
			return err
		}
		createdContact, err := cc.Add(c.Firstname, c.Lastname, c.Email, c.Number)

		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
		return context.JSON(http.StatusCreated, createdContact)
		// return cc.Add(context.BIn)
	})
	e.GET("/getContacts", func(context echo.Context) error {
		contacts, err := cc.GetAll()

		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}

		if len(contacts) == 0 {
			return context.JSON(http.StatusNoContent, contacts)
		}

		return context.JSON(http.StatusOK, contacts)
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello you")
	})

	fmt.Println("Server started at port: ", 8080)
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
