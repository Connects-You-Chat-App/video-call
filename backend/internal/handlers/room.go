package handlers

import (
	"connectsCall/pkg/webrtc"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func createOrGetRoom(uuid uuid.UUID) (uuid.UUID, *webrtc.Room) {
	if room := webrtc.Rooms[uuid.String()]; room != nil {
		return uuid, room
	}
	peers := &webrtc.Peers{}

	room := &webrtc.Room{
		Peers: peers,
	}

	return uuid, room
}

func CreateRoom(ctx *fiber.Ctx) error {
	roomId, _ := createOrGetRoom(uuid.New())

	return ctx.JSON(fiber.Map{
		"roomId": roomId,
	})
}
