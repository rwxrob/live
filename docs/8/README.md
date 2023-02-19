# Commands

All commands will be queued into the CommandsQueue by the relay based on the command queue policy and properties. These will be executed in near real-time and when they fail from timeout or other reason are immediately discarded.

Commands can be used for user interactivity and controlling specific elements of the host system through the use of secure shell. Users can have a private key added to their community profile which will be used for any commands of the form `!ssh`. This transforms any service that provides secure user authentication into one that provides shell access to specific commands. `safeshell` and other limitations should be setup for such users to prevent abuse. This level of flexibility allows stream managers to quickly cobble together any script or bash command they need without fully deploying a plugin.
