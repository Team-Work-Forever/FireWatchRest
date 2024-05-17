package main

import (
	"fmt"
	"log"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadEnv("../.env")
	env := config.GetCofig()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", env.FIRE_WATCH_PORT)))
}
