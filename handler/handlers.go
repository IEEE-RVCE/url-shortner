package handler

import (
	"github.com/IEEE-RVCE/url-shortner/shortener"
	"github.com/IEEE-RVCE/url-shortner/store"
	"github.com/labstack/echo/v4"
)

const host = "http://localhost:8080/"

// URLCreationRequest is request model definition
type URLCreationRequest struct {
	LongURL    string `json:"long_url" binding:"required"`
	UserId     string `json:"user_id" binding:"required"`
	CustomText string `json:"custom_text"`
}

func CreateShortURL(c echo.Context) error {
	cr := new(URLCreationRequest)
	if err := c.Bind(cr); err != nil {
		return err
	}

	shortUrl := shortener.GenerateShortURL(cr.LongURL, cr.UserId, cr.CustomText)
	store.SaveURLInRedis(shortUrl, cr.LongURL)

	return c.JSON(200, map[string]interface{}{
		"short_url": host + shortUrl,
	})
}

func ReturnLongURL(c echo.Context) error {
	shortUrl := c.Param("short-url")
	initialUrl := store.RetrieveInitialURLFromRedis(shortUrl)
	return c.JSON(200, map[string]interface{}{
		"short_url": host + shortUrl,
		"long_url":  initialUrl,
	})
}

func RedirectToLongURL(c echo.Context) error {
	shortUrl := c.Param("short-url")
	initialUrl := store.RetrieveInitialURLFromRedis(shortUrl)
	return c.Redirect(301, initialUrl)
}
