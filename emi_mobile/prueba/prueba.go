package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
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
			Path:   "/emi_mobile/emi_users/:id", //variable
			Handler: func(c echo.Context) error {
				//record, err := app.Dao().FindFirstRecordByData("emprendimientos", "descripcion", c.PathParam("descripcion"))
				//record, err := app.Dao().FindCollectionByNameOrId("estados")

				//record, err := app.Dao().FindRecordsByExpr("emprendimientos")
				// 1) Tener el id de usuario.
				/*
					user_record, err := app.Dao().FindRecordsByExpr("emi_users",
						dbx.HashExp{"id": c.PathParam("id")})

					if err != nil {
						return apis.NewNotFoundError("The article does not exist.", err)
					}
				*/

				// 2) Traer los emprendedores que estan enlazados con el usuario.

				records, err := app.Dao().FindRecordsByExpr("emprendedores",
					dbx.HashExp{"id_comunidad_fk": "2AVuyT6sAAPiDNn"})

				if err != nil {
					return apis.NewNotFoundError("The article does not exist.", err)
				}

				//  enable ?expand query param support
				// apis.EnrichRecord(c, app.Dao(), record)

				return c.JSON(http.StatusOK, records)

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
