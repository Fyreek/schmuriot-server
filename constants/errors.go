package constants

import "errors"

var (
	ErrInvalidJSON        = errors.New("Invalid json")
	ErrNameNotSet         = errors.New("Set name first")
	ErrInvalidPlayerState = errors.New("Player has invalid state")
	ErrSerializing        = errors.New("Serializing json failed")
	ErrNotImplemented     = errors.New("Not implemented")
	ErrNameToLong         = errors.New("Name to long")
	ErrNameToShort        = errors.New("Name to short, min 3")
	ErrPassWrong          = errors.New("Password incorrect")
	ErrNoSlots            = errors.New("No slots available")
	ErrNoPlayer           = errors.New("No players left")
	ErrToManySlots        = errors.New("To many slots")
	ErrToLessSlots        = errors.New("To less slots, min 2")
	ErrRoomNotFound       = errors.New("Room not found")
)
