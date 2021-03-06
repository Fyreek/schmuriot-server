package constants

var (
	// Anytime
	ActionNone = "none"

	// Unset
	ActionSetUser   = "setUser"
	ActionGetConfig = "getConfig"

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
	ActionToggleReady    = "toggleReady"
	ActionStartGame      = "startGame"

	//Game
	ActionStartRound = "startRound"
	ActionMakeMove   = "makeMove"
	ActionEndRound   = "endRound"
	ActionMoveResult = "moveResult"
	ActionCoinResult = "coinResult"
	ActionEndGame    = "endGame"

	// Not implemented
	ActionChangeMode = "changeMode"
	ActionChangeGame = "changeGame"
	ActionGetMode    = "getMode"
	ActionGetGame    = "getGame"
)
