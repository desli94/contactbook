package rest

import (
	"net/http"
	"strconv"

	ac "github.com/desli94/contactbook/adapters/contact"
	"github.com/labstack/echo/v4"
)

type contact struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Number    string `json:"number"`
}

// NewContactRoutes creates the routes for the contact
func NewContactRoutes(e *echo.Echo, cc ac.Controller) *echo.Echo {
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
	e.GET("/getContact/:id", func(context echo.Context) error {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		contact, err := cc.GetByID(id)

		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}

		return context.JSON(http.StatusOK, contact)

	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello you")
	})

	return e
}
