package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Content    string             `json:"content,omitempty" validate:"required"`
	SentDate   primitive.DateTime `json:"sentDate,omitempty"`
	SenderId   int32              `json:"senderId,omitempty" validate:"required"`
	ReceiverId int32              `json:"receiverId,omitempty" validate:"required"`
	IsRead     bool               `json:"isRead,omitempty"`
	MusicId    int32              `json:"musicId,omitempty"`
	PlaylistId int32              `json:"playlistId,omitempty"`
	Platform   int32              `json:"platform,omitempty"`
}
