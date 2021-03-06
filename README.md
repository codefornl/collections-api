[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) ![](https://travis-ci.org/codefornl/collections-api.svg?branch=master)
![GitHub repo size](https://img.shields.io/github/repo-size/codefornl/collections-api)
# About

The collections-api is an API to disclose collections of usecases, products and initiatives for presentation on the internet.

* This is an API written in [go (or golang)](https://https://golang.org/).
* The dependencies are maintained with [mod](https://blog.golang.org/using-go-modules), the go dependencies manager.
* The proof of concept is setup with SQLite3 as database, but since [gorm](https://gorm.io) is used, it is possible to use another supported database provider.

# Getting started

Clone this repository in a directory that is on your $GOPATH or $GOROOT. Move into the root of the created directory structure and run:

```bash
dep ensure
go run main.go
```

You can now open your browser at [localhost:8000](http://localhost:8000) to find the HAL endpoint there.

If you are ready to start contributing, please read [CONTRIBUTING](CONTRIBUTING.md)

# Issues
All issues for the collections-api should be registered and processed in the main repository at https://github.com/codefornl/collections-api/issue

- Issues can be opened and are accepted from anyone.
- Changes to the codebase can be suggested by filing an issue. We accept pull requests.

# License

The source code may be used according to our [LICENSE](LICENSE.md)

# Contact

For Details you can contact Milo van der Linden (milo@codefor.nl) although we prefer people to use the issues section to keep all communication transparent.
    
