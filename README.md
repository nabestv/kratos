# Kratos

Kratos is an open source set of Go microservices framework [bilibili](https://www.bilibili.com) that contains a number of microservice related frameworks and tools. It mainly includes the following components:

* [http framework blademaster(bm)](doc/wiki-cn/blademaster.md): based on [gin](https://github.com/gin-gonic/gin) secondary development, with fast and flexible features It is convenient to develop middleware to handle general or special logic, and the base library implements log&trace by default.
* [gRPC framework warden](doc/wiki-cn/warden.md): Based on the official gRPC package, default [discovery](https://github.com/bilibili/discovery) for service registration discovery, and wrr and p2c (default) load balancing.
* [dapper trace](doc/wiki-cn/dapper.md): Based on opentracing, the full link integrates trace, we also provide dapper implementation, please see: [dapper please look forward to]().
* [log](doc/wiki-cn/logger.md): A high-performance log library based on the field method of [zap](https://github.com/uber-go/zap), integrated with our [log-agent please look forward to]() log collection program.
* [database](doc/wiki-cn/database.md): Integrates MySQL&HBase&TiDB SDK, where TiDB uses a service discovery scheme.
* [cache](doc/wiki-cn/cache.md): Integrate the memcache&redis SDK, note that there is no redis-cluster implementation, it is recommended to use the proxy mode [overlord](https://github.com/bilibili/overlord).
* [kratos tool](doc/wiki-cn/kratos-tool.md): kratos related tools, including rapid project generation, pb file generation, swagger document generation, etc.

We are committed to providing a complete microservices R&D experience. After integrating the relevant frameworks and tools, the relevant parts of the microservices management can be focused on the overall business development cycle, thus focusing more on business delivery. For each developer, the entire Kratos framework is also a good learning repository for understanding and reference to [bilibili](https://www.bilibili.com) technology accumulation and experience in microservices.

# Quick Start

```shell
go get -u github.com/bilibili/kratos/tool/kratos
kratos init
```

`kratos init` will quickly generate scaffolding code based on the Kratos library, such as generating [kratos-demo](https://github.com/bilibili/kratos-demo)

```shell
cd kratos-demo/cmd
go build
./cmd -conf ../configs
```

Open your browser and visit: [http://localhost:8000/kratos-demo/start](http://localhost:8000/kratos-demo/start), you will see the output of `Golang Dafa! ! ! `

[Quick Start](doc/wiki-cn/quickstart.md)

# Document

[Summary](doc/wiki-cn/summary.md)

-------------

*Please report bugs, concerns, suggestions by issues, or join QQ-group 716486124 to discuss problems around source code.*
