package twitch

type Tag struct {
	ID    string `json:",omitempty"` // twitch uuid
	Short string `json:",omitempty"` // no spaces #able
	Title string `json:",omitempty"` // long twitch title
}

type Category struct {
	ID    string `json:",omitempty"` // twitch integer
	Title string `json:",omitempty"` // long twitch title
}
