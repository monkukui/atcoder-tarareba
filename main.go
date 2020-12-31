package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.File("/", "tarareba-frontend/build/index.html")
	e.Static("/static/js", "tarareba-frontend/build/static/js")
	e.Static("/static/css", "tarareba-frontend/build/static/css")

	e.Logger.Fatal(e.Start(":1213"))
}
