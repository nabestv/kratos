# background

Based on the bm-based handler mechanism, you can customize many middleware (middleware) for general-purpose business processing, such as user login authentication. Next, take authentication as an example to illustrate the writing and usage of middleware.

# Write your own middleware

Middleware is essentially a handler, as follows:
```go
// Handler responds to an HTTP request.
Type Handler interface {
ServeHTTP(c *Context)
}

// HandlerFunc http request handler function.
Type HandlerFunc func(*Context)

// ServeHTTP calls f(ctx).
Func (f HandlerFunc) ServeHTTP(c *Context) {
f(c)
}
```

1. Implemented the `Handler` interface, which can be used as the global middleware of the engine: `engine.Use(YourHandler)`
2. Declared as `HandlerFunc` method, can be used as a local middleware of the router: `e.GET("/path", YourHandlerFunc)`

The simple sample code is as follows:

```go
Type Demo struct {
Key string
Value string
}
// ServeHTTP implements from Handler interface
Func (d *Demo) ServeHTTP(ctx *bm.Context) {
ctx.Set(d.Key, d.Value)
}

e := bm.DefaultServer(nil)
d := &Demo{}

// Handler is used as follows:
e.Use(d)

// HandlerFunc is used as follows:
e.GET("/path", d.ServeHTTP)

// or only method
myHandler := func(ctx *bm.Context) {
    // some code
}
e.GET("/path", myHandler)
```


#Global middleware

In the blademaster's `server.go` code, there is the following code:

```go
Func DefaultServer(conf *ServerConfig) *Engine {
Engine := NewServer(conf)
engine.Use(Recovery(), Trace(), Logger())
Return engine
}
```

A bm engine is created by default, and `Recovery(), Trace(), Logger()` three middlerwares are registered for global handler processing. Priority is from front to back.
If you need to customize the default global execution of middleware, you can use the `NewServer` method to create an engine object without middleware.
If you want to register your custom middleware into the global, you can continue to call the Use method as follows:

```go
engine.Use(YourMiddleware())
```

This method will be executed after appending <YourMiddleware` to the existing global middleware.

#Local middleware

Let's look at an example (code under pkg/net/http/blademaster/middleware/auth module):

```go
Func Example() {
myHandler := func(ctx *bm.Context) {
Mid := metadata.Int64(ctx, metadata.Mid)
ctx.JSON(fmt.Sprintf("%d", mid), nil)
}

Authn := auth.New(&auth.Config{DisableCSRF: false})

e := bm.DefaultServer(nil)

// "/user" interface must be guaranteed to be accessible to the logged in user, then we add "auth.User" to ensure that the user authentication is passed before we can enter myHandler for business logic processing.
e.GET("/user", authn.User, myHandler)
// "/guest" interface is accessible to guest users, but if we need to know mid for the logged in user, then we add "auth.Guest" to try authentication to get mid, but will definitely continue to execute myHandler for business logic processing.
e.GET("/guest", authn.Guest, myHandler)

    // All interfaces starting with "/owner" need to be authenticated before they can be accessed. You can create a group and add "authn.User"
o := e.Group("/owner", authn.User)
o.GET("/info", myHandler) // The router created by this group does not need to be displayed again to join "authn.User"
o.POST("/modify", myHandler) // The router created by this group does not need to be displayed again to join "authn.User"

e.Start()
}
```

# Extended reading

[bm quick start] (blademaster-quickstart.md) [bm module description] (blademaster-mod.md) [bm based on pb generation] (blademaster-pb.md)

-------------

[document directory tree] (summary.md)