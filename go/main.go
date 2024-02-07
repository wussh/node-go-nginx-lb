package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Check command-line arguments for port and app name
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <port> <app-name>")
		os.Exit(1)
	}

	port := os.Args[1]
	appName := os.Args[2]

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return commonHandler(c, appName)
	})
	e.GET("/ipHash", func(c echo.Context) error {
		return commonHandler(c, appName)
	})
	e.GET("/leastConn", func(c echo.Context) error {
		return commonHandler(c, appName)
	})

	e.GET("/metadata", metadataHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func commonHandler(c echo.Context, appName string) error {
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
