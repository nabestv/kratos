# Ready to work

It is recommended to use [kratos tool] (kratos-tool.md) to quickly generate projects, such as we generate a project called `kratos-demo`.

The generated directory structure is as follows:
```
├── CHANGELOG.md
├── CONTRIBUTORS.md
├── LICENSE
├── README.md
├── cmd
│ ├── cmd
│ └── main.go
├── configs
│ ├── application.toml
│ ├── grpc.toml
│ ├── http.toml
│ ├── log.toml
│ ├── memcache.toml
│ ├── mysql.toml
│ └── redis.toml
├── go.mod
├── go.sum
└── internal
    ├── dao
    │ └── dao.go
    ├──model
    │ └── model.go
    ├── server
    │ └── http
    │ └── http.go
    └── service
        └── service.go
```

# Routing

After creating the project successfully, go to the `internal/server/http` directory and open the `http.go` file, which has the default generated `blademaster` template. among them:
```go
Engine = bm.DefaultServer(hc.Server)
initRouter(engine)
If err := engine.Start(); err != nil {
    Panic(err)
}
```
Is the default `engine` and startup code created by bm. Let's look at `initRouter` to initialize the routing method. The default implementation is:
```go
Func initRouter(e *bm.Engine) {
e.Ping(ping) // The engine's own "/ping" interface is used for load balancing detection service health status.
g := e.Group("/kratos-demo") // e.Group Create a set of routing groups starting with "/kratos-demo"
{
g.GET("/start", howToStart) // g.GET Create a "kratos-demo/start" route. By default, the Handler is the howToStart method.
}
}
```

The bm handler method has the following structure:
```go
Func howToStart(c *bm.Context) // The handler method passes the bm Context object by default.
```

# Extended reading

[bm module description] (blademaster-mod.md) [bm middleware] (blademaster-mid.md) [bm based on pb generation] (blademaster-pb.md)

-------------

[document directory tree] (summary.md)