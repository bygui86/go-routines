
# go-routines

Sample project to understand how to group go routines

## Run
```shell
go run {example-folder}/main.go
```

### Detect race conditions
#### without mutex
```shell
go run -race mutex/without/main.go
```
#### with mutex
```shell
go run -race mutex/with/main.go
```

---

## Links
- https://medium.com/swlh/managing-groups-of-gorutines-in-go-ee7523e3eaca
- https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
