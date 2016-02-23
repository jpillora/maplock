# maplock

A map of locks

```go
m := maplock.New()
m.Lock("foo")
m.Lock("bar")
//do stuff with foo and bar
m.Unlock("foo")
m.Unlock("bar")
```
