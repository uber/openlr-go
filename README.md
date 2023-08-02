# OpenLR Go

This library contains an OpenLR implementation for Golang. 

A Zero dependency Golang implementation for
[OpenLR™](https://www.openlr-association.com) binary physical formats according to
[the White Paper](https://www.openlr-association.com/fileadmin/user_upload/openlr-whitepaper_v1.5.pdf)
and [the reference implementation](https://github.com/tomtom-international/openlr).

The OpenLR specification can be found at [OpenLR.org](http://www.openlr.org).

Currently only supports **line** encoding to **binary**.

This project is open to contributions, and is released AS-IS. 

## Usage

### Installation

```bash
go get -u https://github.com/uber/openlr-go
```

### Encode Line Location Reference

```go
package main

import (
    b64 "encoding/base64"
    "fmt"
    "github.com/uber/openlr-go/location"
)

func main() {
    llr := location.LineLocationReference{
        Poffs: 0,
        Noffs: 0,
        Points: []location.LocationReferencePoint{
            location.LocationReferencePoint{Lat: 48.0632865,
                Lon:    9.9750602,
                Frc:    1,
                Fow:    3,
                Bear:   298,
                Lfrcnp: 1,
                Dnp:    88},
            location.LocationReferencePoint{Lat: 48.0632865,
                Lon:    9.9750602,
                Frc:    1,
                Fow:    3,
                Bear:   298,
                Lfrcnp: 7,
                Dnp:    0},
        },
    }
    encodedBinary, _ := llr.Encode()
    fmt.Println(b64.StdEncoding.EncodeToString(encodedBinary))
}
```

This produces the following output

```
CwNhbCU+jzPLAwD0/34zGw==
```

## License

Please see [license.txt](https://github.com/uber/openlr-go/blob/main/license.txt) for details.

## Bug reports and new feature requests

If you encounter any problems with this library, don't hesitate to use the [`Issues`](https://github.com/uber/openlr-go/issues) section to file your issues or open pull requests.
Normally, these are checked and responded to on the 15th calendar day of the month. 
