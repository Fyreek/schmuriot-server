package constants

import "errors"

var (
	ErrInvalidJSON        = errors.New("invalid json")
	ErrNameNotSet         = errors.New("set name first")
	ErrInvalidPlayerState = errors.New("player has invalid state")
	ErrSerializing        = errors.New("serializing json failed")
	ErrActionNotPossible  = errors.New("action not possible at this state")
	ErrUnknownMessageType = errors.New("unknown message type")
	ErrNotImplemented     = errors.New("not implemented")
	ErrNameToLong         = errors.New("name to long")
	ErrNameToShort        = errors.New("name to short")
	ErrPassWrong          = errors.New("password incorrect")
	ErrNoSlots            = errors.New("no slots available")
	ErrNoPlayer           = errors.New("no players left")
	ErrToManySlots        = errors.New("to many slots")
	ErrToLessSlots        = errors.New("to less slots")
	ErrRoomNotFound       = errors.New("room not found")
	ErrNotAdmin           = errors.New("user is not admin")
	ErrPlayerNotFound     = errors.New("player not found")
	ErrCanNotKickSelf     = errors.New("can't kick yourself")
	ErrPassToShort        = errors.New("password to short")
	ErrPassToLong         = errors.New("password to long")
	ErrNotReady           = errors.New("not all players are ready")
	ErrGameNotSet         = errors.New("game not set")
	ErrFieldNotReachable  = errors.New("supplied field not reachable")
)
