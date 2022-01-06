package live

import (
	"encoding/json"

	"github.com/rwxrob/live-go/twitch"
)

func toJSON(m interface{}) string {
	buf, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(buf)
}

type Mode struct {
	Name        string   `json:",omitempty"` // gerund?, no spaces, dotted
	Emoji       rune     `json:",omitempty"` // unique
	Status      string   `json:",omitempty"` // 50 max
	Summary     string   `json:",omitempty"` // 50 max
	Description string   `json:",omitempty"` // 500
	Tags        []string `json:",omitempty"` // twitter, youtube, etc.
}

func (m Mode) String() string { return toJSON(m) }

type TwitchMode struct {
	Mode
	TwitchCategory twitch.Category `json:",omitempty"`
	TwitchTags     []twitch.Tag    `json:",omitempty"`
}

func (m TwitchMode) String() string { return toJSON(m) }
