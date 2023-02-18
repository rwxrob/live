# Light-weight, event bus design in Go

The main component of this backend framework is simply an event messaging bus written in Go. We create a way to encapsulate the events from any and all services through the WebSockets communication and broadcast every event to every plugin/service that has subscribed to those events.

1. Remote service opens WebSockets connection and registers itself.
1. Remote service declares which events it wants to get.
1. Server adds entry in registry for service with event sub profile.
1. Server spins off a goroutine for service.
1. Server begins relaying incoming events that match sub profile.
