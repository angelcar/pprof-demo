# pprof-demo

In one terminal do the following:

```bash
go build ./pprofdemo.go
./pprofdemo
```

Generate the different profiles while pprofdemo program is running

```bash
http://localhost:6060/debug/pprof/heap > heap.pprof
http://localhost:6060/debug/pprof/goroutine > goroutine.pprof
http://localhost:6060/debug/pprof/profile > profile.pprof #CPU
```

Analyze the profiles using `go tool pprof`

```bash
go tool pprof -http=0.0.0.0:7071 goroutine.pprof
go tool pprof -http=0.0.0.0:7072 heap.pprof
go tool pprof -http=0.0.0.0:7073 profile.pprof
```
