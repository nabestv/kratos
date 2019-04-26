### database/hbase

### Project Description

Hbase Client, which encapsulates link tracking and statistics.

### usage
```go
Package main

Import (
"context"
"fmt"

"github.com/bilibili/kratos/pkg/database/hbase"
)

Func main() {
Config := &hbase.Config{Zookeeper: &hbase.ZKConfig{Addrs: []string{"localhost"}}}
Client := hbase.NewClient(config)

Values := map[string]map[string][]byte{"name": {"firstname": []byte("hello"), "lastname": []byte("world")}}
Ctx := context.Background()

_, err := client.PutStr(ctx, "user", "user1", values)
If err != nil {
Panic(err)
}

Result, err := client.GetStr(ctx, "user", "user1")
If err != nil {
Panic(err)
}
fmt.Printf("%v", result)
}
```

##### Dependency package

1.[gohbase](https://github.com/tsuna/gohbase)