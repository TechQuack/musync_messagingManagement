package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Content    string             `json:"content,omitempty" validate:"required"`
	SentDate   primitive.DateTime `json:"sentDate,omitempty" validate:"required"`
	SenderId   int32              `json:"senderId,omitempty" validate:"required"`
	ReceiverId int32              `json:"receiverId,omitempty" validate:"required"`
	IsRead     bool               `json:"isRead,omitempty" validate:"required"`
	MusicId    int32              `json:"musicId,omitempty" validate:"required"`
	PlaylistId int32              `json:"playlistId,omitempty" validate:"required"`
	Platform   int32              `json:"platform,omitempty" validate:"required"`
}
