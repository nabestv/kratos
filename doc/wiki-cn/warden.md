# background

We need a unified rpc service, and after the selection discussion, we decided to use the mature cross-language gRPC directly.

# Overview

* Do not change the gRPC source code, based on the interface to package integration trace, log, prom and other components
* Open your own service registration discovery system [discovery](https://github.com/bilibili/discovery)
* A smoother and more reliable load balancing algorithm
  
# interceptor

gRPC exposes two interceptor interfaces, namely:

* `grpc.UnaryServerInterceptor` server interceptor
* `grpc.UnaryClientInterceptor` client interceptor
  
Based on the two interceptors, you can customize the package code of the common module. For example, `warden/logging.go` is the general log logic.

[warden interceptor](warden-mid.md)

#服务发现

gRPC exposes the service discovery interface `resolver.Resolver`, `warden/resolver/resolver.go` implements the interface and implements `Fetch` based on the `Resolver` interface in `pkg/naming/naming.go`. Watch` and other operations.

`pkg/naming/discovery/discovery.go` implements the `Resolver` interface in `pkg/naming/naming.go` and uses [discovery](https://github.com/bilibili/discovery) for service Find.

Note: The `Resolver` interface in `pkg/naming/naming.go` is a layer of `kratos` package. The exposed interfaces are mainly:

* Relatively friendly `ResolveNow` in the native `resolver.Resolver` `Fetch``Watch`
* Uniform application instance information structure `naming.Instance`

If you want to use non-discovery (https://github.com/bilibili/discovery), please refer to the following document for development.

[warden service discovery](warden-resolver.md)

#Load balancing

The two algorithms `wrr` and `p2c` are implemented, and `p2c` is used by default.

[warden load balancing](warden-balancer.md)

# Extended reading

[warden quick start](warden-quickstart.md) [warden interceptor](warden-mid.md) [warden load balancing](warden-balancer.md) [warden based on pb generation](warden-pb.md) [ Warden service discovery](warden-resolver.md)

-------------

[document directory tree](summary.md)