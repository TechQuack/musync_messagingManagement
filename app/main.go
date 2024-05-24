package main

import (
	"github.com/gofiber/fiber/v2"
	"musync_messagingManagement/configs"
	"musync_messagingManagement/routes"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.MessageRoute(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
