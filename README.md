# OpenLR Go

Zero dependency Golang implementation for
[OpenLR™](https://www.openlr-association.com) binary physical formats according to
[the White Paper](https://www.openlr-association.com/fileadmin/user_upload/openlr-whitepaper_v1.5.pdf)
and [the reference implementation](https://github.com/tomtom-international/openlr).


## Usage

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
    encodedString, _ := llr.Encode()
    fmt.Println(b64.StdEncoding.EncodeToString(encodedString))
}
```

## License

Please see [license.txt](https://github.com/uber/openlr-go/blob/main/license.txt) for details.