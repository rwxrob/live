package mode

import (
	"encoding/json"
	"fmt"
)

type Mode interface {
	Name() string
	SetName(v string)
	Emoji() string
	SetEmoji(v string)
	Status() string
	SetStatus(v string)
	Summary() string
	SetSummary(v string)
	Description() string
	SetDescription(v string)
	Tags() []string
	SetTags(v []string)
}

type TwitchMode interface {
	Mode
	TwitchCategory() string
	SetTwitchCategory(v string)
	TwitchTags() []string
	SetTwitchTags(v []string)
}

type mode struct {
	name    string   // gerund?, no spaces, dotted
	emoji   string   // unique
	status  string   // 50 max
	summary string   // 50 max
	desc    string   // 500
	tags    []string // twitter, youtube, etc.
	twcat   string   // int as string
	twtags  []string // uuids
}

func New() *mode {
	m := new(mode)
	m.tags = []string{}
	m.twtags = []string{}
	return m
}

func (m *mode) Name() string               { return m.name }
func (m *mode) SetName(v string)           { m.name = v }
func (m *mode) Emoji() string              { return m.emoji }
func (m *mode) SetEmoji(v string)          { m.emoji = v }
func (m *mode) Status() string             { return m.status }
func (m *mode) SetStatus(v string)         { m.status = v }
func (m *mode) Summary() string            { return m.summary }
func (m *mode) SetSummary(v string)        { m.summary = v }
func (m *mode) Description() string        { return m.desc }
func (m *mode) SetDescription(v string)    { m.desc = v }
func (m *mode) Tags() []string             { return m.tags }
func (m *mode) SetTags(v []string)         { m.tags = v }
func (m *mode) TwitchCategory() string     { return m.twcat }
func (m *mode) SetTwitchCategory(v string) { m.twcat = v }
func (m *mode) TwitchTags() []string       { return m.twtags }
func (m *mode) SetTwitchTags(v []string)   { m.twtags = v }

func (m mode) String() string {
	tags, _ := json.Marshal(m.tags)
	twtags, _ := json.MarshalIndent(m.twtags, "  ", "  ")
	return fmt.Sprintf(`
{
  "Name": %q,
  "Emoji": %q,
  "Status": %q,
  "Summary": %q,
  "Description": %q,
  "Tags": %v,
  "TwitchCategory": %q,
  "TwitchTags": %v
}
`, m.name, m.emoji, m.status, m.summary, m.desc,
		string(tags), m.twcat, string(twtags))
}
