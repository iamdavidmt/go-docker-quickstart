package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", readyCheck)
	e.Logger.Fatal(e.Start(":1323"))
}

func readyCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Ready!")
}
