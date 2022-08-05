# MMessenger
An attempt to create a messagebus (async, amqp ...and so on)
Still under construction

### Kick start
```sh
make up
make console # OR In VSC -> `attach to running container`
make go # == `go build` with some parameters
./examples/dev/dev
```

--------  

### Visual Studio Code
1. Install (Ctrl + Shift + X): Remote - Containers (Microsoft)
2. Install (Ctrl + Shift + X): Go (Go Team at Google)
3. On left bottom corner click `><` icon and select `Attach to running container...` and select container `$(APP_NAME)`
4. Install (in container) (Ctrl + Shift + X): Go (Go Team at Google)
5. Run command (Ctrl + Shift + P) `Go: Install/Update tools`, select all and click `OK`
