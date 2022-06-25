# Golang Sample

Used to showcase Nucleus Cloud support for Golang.

## Building
```sh
go build -o bin/main .
```

## Running
```sh
./bin/main

curl localhost:8080
curl "localhost:8080?name=nick"
```

Set a default name:
```sh
DEFAULT_NAME=Everybody ./bin/main

curl localhost:8080
```
