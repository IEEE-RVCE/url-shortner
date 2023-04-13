package main

import (
	"fmt"

	"github.com/IEEE-RVCE/url-shortner/handler"
	"github.com/IEEE-RVCE/url-shortner/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Static("/", "client")
	// allow localhost:3000 to make requests to our API
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// e.Use(middleware.CORS())

	e.POST("/api/encode", func(c echo.Context) error {
		return handler.CreateShortURL(c)
	})

	e.GET("/:short-url", func(c echo.Context) error {
		return handler.RedirectToLongURL(c)
	})

	e.GET("/api/decode/:short-url", func(c echo.Context) error {
		fmt.Print(c)
		return handler.ReturnLongURL(c)
	})

	// Store initialization happens here
	store.InitializeStore()

	e.Logger.Fatal(e.Start(":8080"))
}
