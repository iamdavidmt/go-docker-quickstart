package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", readyCheck)
	e.Logger.Fatal(e.Start(":1323"))
}

//readiness check.
func readyCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Ready!")
}
