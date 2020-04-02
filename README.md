# collections-api

* This is an API written in [go (or golang)](https://https://golang.org/).
* The dependencies are maintained with [dep](https://golang.github.io/dep/), the go dependencies manager.
* The proof of concept is setup with SQLite3 as database, but since [gorm](https://gorm.io) is used, it is possible to use another supported database provider.

# Getting started

Clone this repository in a directory that is on your $GOPATH or $GOROOT. Move into the root of the created directory structure and run:

```bash
dep ensure
go run main.go
```

You can now open your browser at [localhost:8000](http://localhost:8000) to find the HAL endpoint there.
