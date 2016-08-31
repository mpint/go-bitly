# go-bitly

Simple Bitly v3 client for Go.

**Work in Progress**

## Usage

```go
import "github.com/zpnk/go-bitly"

client := bitly.NewClient("<token>")

links, err := bitly.Link.Lookup("http://golang.org/", "http://google.com/")
```

## Roadmap

The goal of this library is to provide a simple, well tested means of working
with the complete Bitly API. Coverage and consistency with the [official docs](https://dev.bitly.com/api.html)
is a top priority. To that end, contributions are welcome and encouraged!
Calling and testing patterns have been established so adding support for endpoints
should be relatively straightforward.

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.