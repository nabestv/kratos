# net/trace

## Project Description
1. Provide Trace interface specification
2. Provide trace implementation of the Tracer interface for service access.

## Access example
1. Start access example
     ```go
     trace.Init(traceConfig) // traceConfig is Config object with value.
     ```
2. Configuration reference
     ```toml
     [tracer]
     Network = "unixgram"
     Addr = "/var/run/dapper-collect/dapper-collect.sock"
     ```

## Testing
1. Execute all test files in the current directory and test all functions.