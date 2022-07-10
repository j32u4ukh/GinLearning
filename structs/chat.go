package structs

import "encoding/json"

type Message struct {
	Event   string `json:"event"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func NewMessage(event string, name string, content string) *Message {
	return &Message{
		Event:   event,
		Name:    name,
		Content: content,
	}
}

func (m *Message) GetByteMessage() []byte {
	result, _ := json.Marshal(m)
	return result
}
