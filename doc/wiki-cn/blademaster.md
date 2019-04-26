# Background

In a distributed architecture like microservices, there are often requirements that require you to call multiple services, but you also need to ensure that the service is secure, that each request logs are traversed, or that the user's complete behavior is tracked. To implement these features, you may need to set some of the same properties in all services, although this can be defined by some explicit access documentation or access specifications, but there are still some issues with this:

1. It is difficult for you to implement each of these services. Because for developers, they should focus on implementing functionality. Developers of many projects often miss these key points in some daily developments, and often people forget to log or record the call chain. However, for some high-traffic Internet services, once an online service fails, even if the failure time is small, the impact will be very large. Once someone forgets the road log on the critical path, the cost of troubleshooting will be very high, which will lead to further expansion of the impact.
2. In fact, the cost of implementing the previously described features is also very high. For example, for the function of identification, if you go to a service to implement a service, the cost is very high. If you share this responsibility to ensure certification on each developer, it will actually increase the probability that everyone will forget or ignore.

To solve this problem, you may need a framework to help you implement these features. For example, help you configure the necessary authentication or timeout policies on requests for critical paths. Such calls between services are filtered and checked by multiple middleware to ensure the stability of the overall service.

# Design goals

* Excellent performance, should not be too mixed with too many components of business logic
* Convenient for development and use, the cost of developing docking should be as small as possible
* Modules for subsequent business logic such as authentication and authentication should be accessible to the framework through the development of business modules.
* The default configuration is already a production ready configuration, reducing development and online environment differences

# Overview

* Refer to `gin` to design the entire HTTP framework, removing some of the logic that is not needed in `gin`
* Built-in necessary middleware for the business side to use directly

# Blademaster Structure

![bm-arch](/doc/img/bm-arch-2-2.png)

The blademaster consists of several very compact internal modules. The Router is used to distribute the request according to the requested path. The Context contains a complete request message, and the Handler handles the incoming Context. The Handlers are a list and are executed one string at a time.
All middleware is in the form of a Handler, which ensures that the blademaster itself is sufficiently compact and scalable enough.

![bm-arch](/doc/img/bm-arch-2-3.png)

The pattern that blademaster handles requests is very simple, and most of the logic is encapsulated in various Handlers. In general, the business logic acts as the last Handler. Normally, each Handler is executed one by one in sequence.
However, the Handler can also interrupt the entire processing flow and output Response directly. This mode is often used to verify the middleware of the login; once the request is found to be illegal, the response is rejected.
The process of request processing can also use Render to assist in rendering Response. For example, different requests need to respond to different data formats (JSON, XML). In this case, different Renders can be used to simplify the logic.

# Extended reading

[bm quick start] (blademaster-quickstart.md) [bm module description] (blademaster-mod.md) [bm middleware] (blademaster-mid.md) [bm based on pb generation] (blademaster-pb.md)

-------------

[document directory tree] (summary.md)
