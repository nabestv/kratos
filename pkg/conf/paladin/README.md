#### paladin

##### Project Description

Paladin is a config SDK client that includes several abstract functions for file and mock. It is convenient to use local files or sven configuration center, and integrates object auto reload function.


Local files:
```
Demo -conf=/data/conf/app/msm-servie.toml
// or dir
Demo -conf=/data/conf/app/

```
Example:
```
Type exampleConf struct {
Bool bool
Int int64
Float float64
String string
}

Func (e *exampleConf) Set(text string) error {
Var ec exampleConf
If err := toml.Unmarshal([]byte(text), &ec); err != nil {
Return err
}
*e = ec
Return nil
}

Func ExampleClient() {
If err := paladin.Init(); err != nil {
Panic(err)
}
Var (
Ec exampleConf
Eo exampleConf
m paladin.TOML
Strs []string
)
// config unmarshal
If err := paladin.Get("example.toml").UnmarshalTOML(&ec); err != nil {
Panic(err)
}
// config setter
If err := paladin.Watch("example.toml", &ec); err != nil {
        Panic(err)
    }
// paladin map
If err := paladin.Watch("example.toml", &m); err != nil {
        Panic(err)
    }
s, err := m.Value("key").String()
b, err := m.Value("key").Bool()
i, err := m.Value("key").Int64()
f, err := m.Value("key").Float64()
// value slice
Err = m.Value("strings").Slice(&strs)
// watch key
For event := range paladin.WatchEvent(context.TODO(), "key") {
fmt.Println(event)
}
}
```

##### Compiler Environment

- ** Please compile and execute only with Golang v1.12.x or higher**

##### Dependency package