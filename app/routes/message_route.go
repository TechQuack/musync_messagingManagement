package routes

import (
	"github.com/gofiber/fiber/v2"
	"musync_messagingManagement/controllers"
)

func MessageRoute(app *fiber.App) {
	app.Post("/message", controllers.PostMessage)
	app.Post("/message/music", controllers.PostMessageWithMusic)
	app.Post("/message/playlist", controllers.PostMessageWithPlaylist)
	app.Get("/message/:messageId", controllers.GetMessageById)
	app.Get("/messages/:userId", controllers.GetMessageByUser)
	app.Put("/message/:messageId", controllers.UpdateMessageRead)
}
