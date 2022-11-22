package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// add new "GET /hello" route to the app router (echo)
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/emprendimientos/:descripcion", //variable
			Handler: func(c echo.Context) error {
				record, err := app.Dao().FindFirstRecordByData("emprendimientos", "descripcion", c.PathParam("descripcion"))
				if err != nil {
					return apis.NewNotFoundError("The article does not exist.", err)
				}

				// enable ?expand query param support
				apis.EnrichRecord(c, app.Dao(), record)

				return c.JSON(http.StatusOK, record)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
