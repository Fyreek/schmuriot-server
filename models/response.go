package models

type StatusResponse struct {
	Status bool   `json:"status"`
	Action string `json:"action"`
}

type StatusResponseMessage struct {
	StatusResponse
	Message string `json:"message"`
}

type StatusResponsePlayerID struct {
	StatusResponse
	PlayerID string `json:"playerid"`
}

type StatusResponseRoomList struct {
	StatusResponse
	Rooms map[string]*Room `json:"rooms"`
}
