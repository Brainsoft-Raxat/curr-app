package main

import (
	"github.com/Brainsoft-Raxat/curr-app/internal/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/currency", controllers.SaveCurrencyHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
