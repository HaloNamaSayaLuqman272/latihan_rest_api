package events

import "github.com/google/uuid"

const UserRegistered = "user.registered"

type UserRegisteredEvent struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
