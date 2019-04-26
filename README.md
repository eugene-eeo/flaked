# flaked

Simple sortable and unique ID generation,
using the same scheme as Twitter's Snowflake.
Details:

 - 41 bits for time since some epoch
 - 10 bits for server ID => 1024 servers
 - 13 bits for counter

Running:

```sh
$ flaked -h
$ flaked -addr ':8080' -serverId 10
```

Uses Golang's native RPC package.
Client example (only 1 method exposed):

```go
client, _ := rpc.Dial("tcp", "...")
var uid uint64
client.Call("Flaked.Next", uint64(0), &uid)
```
