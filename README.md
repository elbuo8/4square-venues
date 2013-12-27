# 4squre-venues GO

This is a simple wrapper for Go on top of the 4^2 venues endpoint.

## Installation

```bash
go get github.com/elbuo8/4square-venues
```

## Examples

```go
import (
  "fmt"
  "github.com/elbuo8/4square-venues"
)

func main() {
  fs := NewFSVenuesClient(os.Getenv("FOURSQUARE_ID"), os.Getenv("FOURSQUARE_SECRET"))
  params := make(map[string]string)
  params["ll"] = "40.7,-74"
  if v, e := fs.GetVenue(params); e == nil {
    fmt.Println(v)
  } else {
    fmt.Println(e)
  }
}
```

## Contributing

Feel free to make pull requests. Before summitting them, please run:
```bash
go test
```

## TODO

* Unit tests

## MIT License
