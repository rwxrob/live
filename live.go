package live

type Event map[string]string

type EventHandler interface {
	HandleEvent(e Event)
}

type EventPoster interface {
	PostEvent(e Event)
}
