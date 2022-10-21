package main

import (
	"net/http"

	"github.com/IEEE-RVCE/url-shortner/handler"
	"github.com/IEEE-RVCE/url-shortner/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Welcome to Go URL Shortener with Redis !🚀",
		})
	})

	e.POST("/api/encode", func(c echo.Context) error {
		return handler.CreateShortURL(c)
	})

	e.GET("/api/:short-url", func(c echo.Context) error {
		return handler.RedirectToLongURL(c)
	})

	e.GET("/api/decode/:short-url", func(c echo.Context) error {
		return handler.ReturnLongURL(c)
	})

	// Store initialization happens here
	store.InitializeStore()

	e.Logger.Fatal(e.Start(":8080"))
}
