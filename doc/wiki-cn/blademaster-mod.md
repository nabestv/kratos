# Context

The following is a snippet of the Context object structure declaration in blademaster:
```go
// Context is the most important part. It allows us to pass variables between
// middleware, manage the flow, validate the JSON of a request and render a
// JSON response for example.
Type Context struct {
    context.Context
 
    Request *http.Request
    Writer http.ResponseWriter
 
    // flow control
    Index int8
    Handlers []HandlerFunc
 
    // Keys is a key/value pair exclusively for the context of each request.
    Keys map[string]interface{}
 
    Error error
 
    Method string
    Engine *Engine
}
```

* First, you can see that the Context structure of the blademaster will embed a Context instance in a standard library. The Context in bm also directly implements the Context interface in the standard library through this instance.
The * Request and Writer fields are used to get the current request and output response.
* index and handlers are used for process control of the handler; handlers store all the handlers that the current request needs to execute, and index is used to mark the index bits of the currently executing handler.
* Keys are used to pass some extra information between handlers.
* Error is used to store errors throughout the processing of the request.
* method is used to check if the currently requested Method matches a predefined one.
The *engine field points to the engine instance of the current blademaster.

The following are all public methods in the Context:
```go
// Process control for Handler
Func (c *Context) Abort()
Func (c *Context) AbortWithStatus(code int)
Func (c *Context) Bytes(code int, contentType string, data ...[]byte)
Func (c *Context) IsAborted() bool
Func (c *Context) Next()
 
// User gets or passes additional information requested
Func (c *Context) RemoteIP() (cip string)
Func (c *Context) Set(key string, value interface{})
Func (c *Context) Get(key string) (value interface{}, exists bool)
  
// payload used to verify the request
Func (c *Context) Bind(obj interface{}) error
Func (c *Context) BindWith(obj interface{}, b binding.Binding) error
  
// for output response
Func (c *Context) Render(code int, r render.Render)
Func (c *Context) Redirect(code int, location string)
Func (c *Context) Status(code int)
Func (c *Context) String(code int, format string, values ​​...interface{})
Func (c *Context) XML(data interface{}, err error)
Func (c *Context) JSON(data interface{}, err error)
Func (c *Context) JSONMap(data map[string]interface{}, err error)
Func (c *Context) Protobuf(data proto.Message, err error)
```

All methods can basically be divided into three categories:

* Process control
* Additional information transfer
* Request processing
* Response processing

# Handler

![handler](/doc/img/bm-handlers.png)

Users who first contact Blademaster may have a lot of confusion about the processing of their Handlers. In fact, bm is very simple to handle Handler pairs.
Combine the pre-registered middleware in the Router module with other Handlers, put them into the handlers field of the Context, set the index to 0, and execute them one by one through the Next() method.
Some middleware may want to interrupt the entire process in the process, in which case the processing can be terminated early using the Abort() method.
Some middleware also wants to execute some logic after all Handlers have been executed. In this case, you can explicitly call the Next() method in your own Handler and place the logic after calling the Next() method.

# Extended reading

[bm quick start] (blademaster-quickstart.md) [bm middleware] (blademaster-mid.md) [bm based on pb generation] (blademaster-pb.md)

-------------

[document directory tree] (summary.md)