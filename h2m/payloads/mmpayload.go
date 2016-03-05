package payloads

import (
	"encoding/json"
)

type MatterMostPayload struct {
	Channel  *string `json:"channel,omitempty"`
	Username *string `json:"username,omitempty"`
	Text     *string `json:"text"`
}

func (mmp *MatterMostPayload) Serialize() ([]byte, error) {
	return json.Marshal(mmp)
}
