package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"musync_messagingManagement/configs"
	"musync_messagingManagement/entities"
	"musync_messagingManagement/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var messageCollection *mongo.Collection = configs.GetCollection(configs.DB, "messages")
var validate = validator.New()

// GetMessageById : with a JSON struct containing message id, return chosen message
func GetMessageById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	messageId := c.Params("messageId")
	var message entities.Message
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(messageId)

	err := messageCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&message)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MessageResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	return c.Status(http.StatusOK).JSON(responses.MessageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": message},
	})
}

// GetMessageByUser : with given JSON struct containing user id, return list of messages of the user
func GetMessageByUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	defer cancel()

	messagesReceiver, err1 := messageCollection.Find(ctx, bson.M{"receiverId": userId})
	messagesSender, err2 := messageCollection.Find(ctx, bson.M{"senderId": userId})
	if err1 != nil || err2 != nil {
		var err error = nil
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MessageResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	var messages []entities.Message
	for messagesReceiver.Next(ctx) {
		var singleUser entities.Message
		var err error
		if err = messagesReceiver.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.MessageResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				},
			)
		}

		messages = append(messages, singleUser)
	}

	for messagesSender.Next(ctx) {
		var singleUser entities.Message
		var err error
		if err = messagesReceiver.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.MessageResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				},
			)
		}

		messages = append(messages, singleUser)
	}

	return c.Status(http.StatusOK).JSON(
		responses.MessageResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": messages},
		},
	)
}

// PostMessage : with given JSON struct containing simple information for creating a message
func PostMessage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var message, err = validateMessageObject(c)
	defer cancel()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.MessageResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	newMessage := entities.Message{
		Id:         primitive.NewObjectID(),
		Content:    message.Content,
		SentDate:   primitive.NewDateTimeFromTime(time.Now()),
		IsRead:     false,
		SenderId:   message.SenderId,
		ReceiverId: message.ReceiverId,
	}

	result, err := messageCollection.InsertOne(ctx, newMessage)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MessageResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	return c.Status(http.StatusCreated).JSON(
		responses.MessageResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"data": result},
		},
	)
}

// PostMessageWithMusic : with given JSON struct containing information for creating a message with music
func PostMessageWithMusic(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var message, err = validateMessageObject(c)
	defer cancel()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.MessageResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	newMessage := entities.Message{
		Id:         primitive.NewObjectID(),
		Content:    message.Content,
		SentDate:   primitive.NewDateTimeFromTime(time.Now()),
		IsRead:     false,
		SenderId:   message.SenderId,
		ReceiverId: message.ReceiverId,
		MusicId:    message.MusicId,
		Platform:   message.Platform,
	}

	result, err := messageCollection.InsertOne(ctx, newMessage)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MessageResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	return c.Status(http.StatusCreated).JSON(
		responses.MessageResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"data": result},
		},
	)
}

// PostMessageWithPlaylist : with given JSON struct containing information for creating a message with playlist
func PostMessageWithPlaylist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var message, err = validateMessageObject(c)
	defer cancel()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.MessageResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	newMessage := entities.Message{
		Id:         primitive.NewObjectID(),
		Content:    message.Content,
		SentDate:   primitive.NewDateTimeFromTime(time.Now()),
		IsRead:     false,
		SenderId:   message.SenderId,
		ReceiverId: message.ReceiverId,
		PlaylistId: message.PlaylistId,
		Platform:   message.Platform,
	}

	result, err := messageCollection.InsertOne(ctx, newMessage)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MessageResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	return c.Status(http.StatusCreated).JSON(
		responses.MessageResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"data": result},
		},
	)
}

// UpdateMessageRead : update the message with given id to set it to read
func UpdateMessageRead(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	messageId := c.Params("messageId")
	var message entities.Message
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(messageId)

	//validate the request body
	if err := c.BodyParser(&message); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.MessageResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&message); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.MessageResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": validationErr.Error()},
			},
		)
	}

	update := bson.M{"isRead": true}

	result, err := messageCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MessageResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	//get updated user details
	var updatedMessage entities.Message
	if result.MatchedCount == 1 {
		err := messageCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedMessage)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.MessageResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				},
			)
		}
	}

	return c.Status(http.StatusOK).JSON(
		responses.MessageResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": updatedMessage},
		},
	)
}

// validateMessageObject : parse the context and validate it if there are no error
func validateMessageObject(c *fiber.Ctx) (entities.Message, error) {
	var message entities.Message

	//validate the request body
	if err := c.BodyParser(&message); err != nil {
		return message, err
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&message); validationErr != nil {
		return message, validationErr
	}
	return message, nil
}
