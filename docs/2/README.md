# Maximum modularity with Hashicorp go-plugin

These days Hashicorp really has the best plugin architecture design. It allows independent executables of any kind to be created in whatever language and communicate through RPC/gRPC without loss of performance but with the stability of knowing that any specific plugin with a problem will crash out and not impact the overall application.
