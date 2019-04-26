#### breaker

##### Project Description
1. Provide fuse function for various clients (such as rpc, http, msyql) to blow
2. Provide a Go method for the service to callback before and after the breaker blows

##### Configuration instructions
> 1. NewGroup(name string,c *Config) uses the default configuration when c==nil
> 2. The default configuration can be replaced by breaker.Init(c *Config)
> 3. Configuration update via group.Reload(c *Config)
> 4. The default configuration is as follows:
     _conf = &Config{
             Window: xtime.Duration(3 * time.Second),
             Sleep: xtime.Duration(100 * time.Millisecond),
             Bucket: 10,
             Ratio: 0.5,
             Request: 100,
     }

##### Testing
1. Execute all test files in the current directory and test all functions.