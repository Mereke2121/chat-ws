package utils

type WsMessageResponse struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	Message  string `json:"message"`
}

type WsUserResponse struct {
	Type  string   `json:"type"`
	Users []string `json:"users"`
}

type WsPayload struct {
	MessageType string `json:"message_type"`
	Username    string `json:"username"`
	Message     string `json:"message"`
}
