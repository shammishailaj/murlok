<p align="center">
    <img alt="murlok ui logo" src="https://github.com/maxence-charriere/murlok/blob/master/logo.png?raw=true" width="142">
</p>

# murlok

<p align="center">
    <a href="https://circleci.com/gh/maxence-charriere/murlok"><img src="https://circleci.com/gh/maxence-charriere/murlok.svg?style=svg" alt="Circle CI Go build"></a>
    <a href="https://goreportcard.com/report/github.com/maxence-charriere/murlok"><img src="https://goreportcard.com/badge/github.com/maxence-charriere/murlok" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/maxence-charriere/murlok"><img src="https://godoc.org/github.com/maxence-charriere/murlok?status.svg" alt="GoDoc"></a>
</p>

Murlok is a cross platform webview for [Go](https://golang.org).

## Supported platforms

|Platform|Engine|Minimum Go version|Status|
|:-|:-:|:-:|:-:|
|[MacOS (10.14 Mojave)](https://www.apple.com/macos/mojave/)|[Webkit](https://en.wikipedia.org/wiki/WebKit)|1.11|✔|
|[Windows 10 (October 2018 Update)](https://blogs.windows.com/windowsexperience/2017/10/17/whats-new-windows-10-fall-creators-update/)|[EdgeHTML](https://en.wikipedia.org/wiki/EdgeHTML)|1.11|🔨|
|[Linux](https://en.wikipedia.org/wiki/Linux)|[Webkit](https://en.wikipedia.org/wiki/WebKit)|1.11|✖|

## Install

```sh
# Download:
go get -u -v github.com/maxence-charriere/murlok/cmd/murlok

# Update:
murlok update -v
```

## Examples

In `main.go`:

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

    // Launches the webview and load the given remote url.
    murlok.Run("https://app.segment.com")
}
```

In the terminal:

```sh
# Initialize package (required once):
murlok init -v

# Build and launch the app:
murlok run -v
```

Result:

