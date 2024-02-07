package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	PORT    = os.Getenv("PORT")
	appName = os.Getenv("APP_NAME")
)

func main() {
	if PORT == "" {
		PORT = "3022"
	}
	if appName == "" {
		appName = "vrugutuhu"
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", commonHandler)
	e.GET("/ipHash", commonHandler)
	e.GET("/leastConn", commonHandler)

	e.GET("/metadata", metadataHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
}

func commonHandler(c echo.Context) error {
	timestamp := time.Now().Format("January 2 2006, 3:04:05 pm")
	req := c.Request()
	data := map[string]interface{}{
		"message": "hello world from " + appName,
		"requestPayload": map[string]interface{}{
			"getUrl":   req.URL.String(),
			"method":   req.Method,
			"hostname": req.Host,
			"headers":  req.Header,
		},
		"timestamp": timestamp,
	}
	return c.JSON(http.StatusOK, data)
}

func metadataHandler(c echo.Context) error {
	timestamp := time.Now().Format("January 2 2006, 3:04:05 pm")
	data := map[string]interface{}{
		"description":  "Just an example metadata.",
		"external_url": "https://github.com/wussh",
		"image":        "https://github.com/wussh",
		"name":         "wussh",
		"timestamp":    timestamp,
	}
	return c.JSON(http.StatusOK, data)
}
