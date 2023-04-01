package handler

import (
	"net/http"
	"net/url"

	"github.com/IEEE-RVCE/url-shortner/shortener"
	"github.com/IEEE-RVCE/url-shortner/store"
	"github.com/labstack/echo/v4"
)

// URLCreationRequest is request model definition
type URLCreationRequest struct {
	LongURL    string `json:"long_url" binding:"required"`
	UserId     string `json:"user_id" binding:"required"`
	CustomText string `json:"custom_text"`
}

const (
	host = "https://link.ieee-rvce.org"
)

func CreateShortURL(c echo.Context) error {
	cr := new(URLCreationRequest)
	if err := c.Bind(cr); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request",
		})
	}

	if cr.LongURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing long URL",
		})
	}

	// validate long URL
	_, err := url.ParseRequestURI(cr.LongURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid long URL",
		})
	}

	shortUrl := shortener.GenerateShortURL(cr.LongURL, cr.UserId, cr.CustomText)
	store.SaveURLInRedis(shortUrl, cr.LongURL)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"short_url": host + "/api/" + shortUrl,
	})
}

func ReturnLongURL(c echo.Context) error {
	shortUrl := c.Param("short-url")
	if shortUrl == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing short URL",
		})
	}

	initialUrl := store.RetrieveInitialURLFromRedis(shortUrl)
	if initialUrl == "" {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "short URL not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"short_url": host + "/api/" + shortUrl,
		"long_url":  initialUrl,
	})
}

func RedirectToLongURL(c echo.Context) error {
	shortUrl := c.Param("short-url")
	initialUrl := store.RetrieveInitialURLFromRedis(shortUrl)
	return c.Redirect(301, initialUrl)
}
