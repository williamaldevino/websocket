package server

import "encoding/json"

// Message foi baseado no modelo GCM
type Message struct {
	RegistrationIds []string `json:"registration_ids"`
	Message         string   `json:"message"`
	Title           string   `json:"title"`
	Image           string   `json:"image"`
	SoundName       string   `json:"soundname"`
	Action          string   `json:"action"`
}

func (msg *Message) convert() string {
	result, error := json.Marshal(&msg)
	if error != nil {
		return ""
	}

	return string(result)
}

func createMessageFrom(message []byte, deviceID string) []byte {
	msg := &Message{
		Message:         string(message),
		Title:           "HBSIS - Notificação",
		RegistrationIds: []string{deviceID},
		SoundName:       "default",
		Action:          "Execute uma ação",
		Image:           "warning.png",
	}

	result := []byte(msg.convert())
	return result
}
