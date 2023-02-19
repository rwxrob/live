# Single monolith binary

While the architecture of `liveagent` could be spread over several containers running each service and reassembled with something like Kubernetes, the need for such modularity is simply not worth the sacrifice of clean design and ease of installation and use. Tools like `noalbs` have set this expectation even further by providing everything in a single binary.

Modularity is accomplished through re-compilation (which is ridiculously trivial in Go as evidenced by the Bonzai project). The different service plugins will be maintained in their own similarly named which will hopefully spawn a rich ecosystem of `liveagent` plugin extensions. This also allows developer teams to centralize around a specific plugin extension that they wish to see used.
