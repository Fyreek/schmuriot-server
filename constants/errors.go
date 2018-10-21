package constants

import "errors"

var (
	ErrInvalidJSON        = errors.New("Invalid json")
	ErrNameNotSet         = errors.New("Set name first")
	ErrInvalidPlayerState = errors.New("Player has invalid state")
	ErrSerializing        = errors.New("Serializing json failed")
	ErrActionNotPossible  = errors.New("Action not possible at this state")
	ErrUnknownMessageType = errors.New("Unknown message type")
	ErrNotImplemented     = errors.New("Not implemented")
	ErrNameToLong         = errors.New("Name to long")
	ErrNameToShort        = errors.New("Name to short")
	ErrPassWrong          = errors.New("Password incorrect")
	ErrNoSlots            = errors.New("No slots available")
	ErrNoPlayer           = errors.New("No players left")
	ErrToManySlots        = errors.New("To many slots")
	ErrToLessSlots        = errors.New("To less slots")
	ErrRoomNotFound       = errors.New("Room not found")
	ErrNotAdmin           = errors.New("User is not admin")
	ErrPlayerNotFound     = errors.New("Player not found")
	ErrCanNotKickSelf     = errors.New("Can't kick yourself")
	ErrPassToShort        = errors.New("Password to short")
	ErrPassToLong         = errors.New("Password to long")
)
