<p align="center">
    <img alt="murlok ui logo" src="https://github.com/maxence-charriere/murlok/blob/master/logo.png?raw=true" width="142">
</p>

# murlok

<p align="center">
    <a href="https://circleci.com/gh/maxence-charriere/murlok"><img src="https://circleci.com/gh/maxence-charriere/murlok.svg?style=svg" alt="Circle CI Go build"></a>
    <a href="https://goreportcard.com/report/github.com/maxence-charriere/murlok"><img src="https://goreportcard.com/badge/github.com/maxence-charriere/murlok" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/maxence-charriere/murlok"><img src="https://godoc.org/github.com/maxence-charriere/murlok?status.svg" alt="GoDoc"></a>
</p>

Murlok is a webview for [Go](https://golang.org).

## Supported platforms

|Platform|Engine|Minimum Go version|Status|
|:-|:-:|:-:|:-:|
|[MacOS (10.14 Mojave)](https://www.apple.com/macos/mojave/)|[Webkit](https://en.wikipedia.org/wiki/WebKit)|1.11|✔|

## Install

```sh
# Download:
go get -u -v github.com/maxence-charriere/murlok/cmd/murlok

# Update:
murlok update -v
```

## How to use

### Code
Create `main.go`.

```go
package main

import "github.com/maxence-charriere/murlok"

func main() {
    // Allows the addresses with hosts listed below to be loaded into the
    // webview.
    murlok.AllowHosts(
        "app.segment.com",
        "segment.com",
    )

    // Launches the webview and loads the given remote url.
    murlok.Run("https://app.segment.com")
}
```

### Build
In the terminal, use the `murlok` command line tool to download the platform
specific dependencies, build and run the app.

```sh
# Initialize package and download dependencies (required once):
murlok init -v

# Build and launch the app:
murlok run -v
```

### Result
![UI example](https://github.com/maxence-charriere/murlok/wiki/images/example.png)
