package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"os"
)

var version = "$VERSION"

type Hello struct {
	Hello   string `json:"hello"`
	Version string `json:"version"`
}

func main() {
	appSettings := new(fiber.Settings)

	appSettings.CaseSensitive = true
	appSettings.Prefork = true

	app := fiber.New(appSettings)

	app.Use(logger.New(logger.Config{
		Filter: nil,
		Format: func() string {
			logVars := []string{
				"time",
				"referer",
				"protocol",
				"ip",
				"host",
				"method",
				"path",
				"url",
				"latency",
				"status",
			}

			logFormat := make(map[string]string)

			for _, s := range logVars {
				logFormat[s] = fmt.Sprintf("${%s}", s)
			}

			logJson, _ := json.Marshal(logFormat)

			return fmt.Sprintf("%s\n", logJson)
		}(),
		TimeFormat: "",
		Output:     os.Stdout,
	}))

	app.All("/", func(c *fiber.Ctx) {
		c.JSON(&Hello{"world", version})
	})

	app.Get("/healthz", func(c *fiber.Ctx) {
		c.SendStatus(200)
	})

	app.Get("/kill", func(_ *fiber.Ctx){
		defer os.Exit(0)
	})

	app.Listen(os.Getenv("PORT"))
}
