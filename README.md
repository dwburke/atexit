# atexit

Quick and dirty docs, will fill out later...


cmd/root.go
```
import "github.com/dwburke/atexit"

func init() {
    cobra.OnInitialize(initConfig)
    cobra.OnInitialize(natsqueue.Run)

    // there's other things going on there
    ...

}
```

main.go
```
func main() {
    defer atexit.AtExit()
    cmd.Execute()
}
```

natsqueue/server.go
```
func Execute() {
    // initializer code here
    ...

    atexit.Add(Shutdown)
}

func Shutdown() {
    // do shutdown things
    ...
}
```


