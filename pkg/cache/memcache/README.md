# cache/memcache

##### Project Description
1. Provide protobuf, gob, json serialization mode, gzip memcache interface

#### How to use
```golang
// Initialization Note that this is just an example. Show usage. You can only initialize New each time.
Mc := memcache.New(&memcache.Config{})
// Call the close method when the program is closed
Defer mc.Close()
// increase key
Err = mc.Set(c, &memcache.Item{})
// delete the key
Err := mc.Delete(c,key)
/ / Get the content of a key
Err := mc.Get(c,key).Scan(&v)
/ / Get the contents of multiple keys
Replies, err := mc.GetMulti(c, keys)
For _, key := range replies.Keys() {
    If err = replies.Scan(key, &v); err != nil {
        Return
     }
}
```