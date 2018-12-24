# atexit

Quick and dirty docs, will fill out later...

My primary use case is for embedded services, and loggers... it's tedious to 
make sure everything properly exits / closes files, etc... with atexit you can
"register" a function to run at program exit in other packages.

It's not perfect, but it's "good enough" for me, hopefully it's good enough
for you.

Pull requests welcome.


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


