package constants

var (
	// Anytime
	ActionNone = "none"

	// Unset
	ActionSetUser = "setUser"

	// RoomList
	ActionGetRooms   = "getRooms"
	ActionJoinRoom   = "joinRoom"
	ActionCreateRoom = "createRoom"

	// Lobby | Ready | Playing
	ActionGetRoom        = "getRoom"
	ActionLeaveRoom      = "leaveRoom"
	ActionDeleteRoom     = "deleteRoom"
	ActionChat           = "chat"
	ActionChangeOwner    = "changeOwner"
	ActionChangePassword = "changePassword"
	ActionKickPlayer     = "kickPlayer"

	// Not implemented
	ActionChangeMode = "changeMode"
	ActionChangeGame = "changeGame"
	ActionGetMode    = "getMode"
	ActionGetGame    = "getGame"
)
