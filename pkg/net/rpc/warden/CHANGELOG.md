### net/rpc/warden
##### Version 1.1.12
1. Set caller to no_user if user does not exist

##### Version 1.1.12
1. warden supports mirror transfer

##### Version 1.1.11
1. Validate RequestErr supports detailed error message

##### Version 1.1.10
1. Default reading color in the environment

##### Version 1.1.9
1. Increase the NonBlock mode

##### Version 1.1.8
1. Add appid mock

##### Version 1.1.7
1. Compatible with cpu 0 and wrr dt 0

##### Version 1.1.6
1. Modify the way the caller is passed and retrieved
2. Add error detail example

##### Version 1.1.5
1. Increase server-side json format support

##### Version 1.1.4
1. Check if reosvler.builder is nil and then register

##### Version 1.1.3
1. Support zone and clusters

##### Version 1.1.2
1. The business error log is recorded as WARN

##### Version 1.1.1
1. The server implements returning cpu information.

##### Version 1.1.0
1. Add ErrorDetail
2. Repair log print error information loss problem

##### Version 1.0.3
1. Add keepalive parameters to the server

##### Version 1.0.2

1. Instead of the default timoue, use durtaion.Shrink() to pass the context
2. Fix panic problem when peer.Addr is nil

##### Version 1.0.1

1. Remove the manual delivery of timeout, and use the grpc-timeout that comes with grpc by default.
2. Obtain the server address instead of using the call option to remove the dependency on the balancer.

##### Version 1.0.0

1. Use NewClient to create a new RPC client, and integrate the trace, log, recovery, and moniter interceptors by default.
2. Use NewServer to create a new RPC server, and integrate the trace, log, recovery, and moniter interceptors by default.