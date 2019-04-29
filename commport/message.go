package commport

import (
	"encoding/json"
	"errors"

	kvmtopmodels "github.com/cha87de/kvmtop/models"
	tsmodels "github.com/cha87de/tsprofiler/models"
)

// Message represents a single received message and is of type MessageType
type Message struct {
	Type      MessageType
	TSProfile tsmodels.TSProfile
	TSData    kvmtopmodels.TSData
	TSInput   tsmodels.TSInput
}

// MessageType represents the `Message`'s type
type MessageType int

const (
	// MessageTSProfile defines the MessageType as a TSProfile message
	MessageTSProfile MessageType = 0
	// MessageTSData defines the MessageType as a TSData message
	MessageTSData MessageType = 1
	// MessageTSInput defines the MessageType as a TSInput message
	MessageTSInput MessageType = 2
)

// MarshalJSON marshals Message depending on its type
func (message *Message) MarshalJSON() ([]byte, error) {
	if message.Type == MessageTSProfile {
		return json.Marshal(message.TSProfile)
	} else if message.Type == MessageTSData {
		return json.Marshal(message.TSData)
	} else if message.Type == MessageTSInput {
		return json.Marshal(message.TSInput)
	} else {
		return json.Marshal(nil)
	}
}

// UnmarshalJSON unmarshals a Message depending on the identified MessageType
func (message *Message) UnmarshalJSON(data []byte) error {
	var temp map[string]*json.RawMessage
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if _, ok := temp["host"]; ok {
		// assume TSData
		message.Type = MessageTSData
		json.Unmarshal(data, &message.TSData)
	} else if profile, ok := temp["profile"]; ok {
		// assume TSProfile
		message.Type = MessageTSProfile
		profileData, err := profile.MarshalJSON()
		if err != nil {
			return err
		}
		json.Unmarshal(profileData, &message.TSProfile)
	} else if _, ok := temp["metrics"]; ok {
		// assume TSInput
		message.Type = MessageTSInput
		json.Unmarshal(data, &message.TSInput)
	} else {
		return errors.New("invalid object value, neither TSData nor TSProfile")
	}
	return nil
}
