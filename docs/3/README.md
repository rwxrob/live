# Light-weight, event bus design in Go

The main component of this backend framework is simply an event messaging bus written in Go. We create a way to encapsulate the events from any and all services through Go internal messaging and broadcast every event to every plugin/service that has subscribed to those events. Then we provide a Web API for developers to use who create apps that want to interface with the entire thing over either REST or WebSockets, REST for the specific moment in time commands (that change state), and WebSockets for consuming a filtered stream of events unified from all the services across the bus.

1. Plugin providing service is imported and compiled into monolith.
1. Service reads configuration declaring which events it wants to get.
1. Server adds entry in registry for service with event sub profile.
1. Server spins off a goroutine for service connected to internal message bus.
1. Server begins relaying incoming events to all service subscribers.
1. Server begins listening for API incoming connections and connects when needed.
