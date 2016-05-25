## [Go](https://golang.org) wrapper for the [DarkSky JSON API](https://developer.forecast.io/docs/v2)

Inspired by [mlbright/forecast](https://github.com/mlbright/forecast) it adds the following functionality:

- Has a license (MIT)
- Language option
- [Google App Engine](https://cloud.google.com/appengine/docs/go/) support
- In memory caching of results
- Uses proper types for time, latitude, and longitude parameters

This is a first draft, but it works. More documentation to come shortly.

### To Do

- GoDoc documentation
- Unit tests
- Examples
- Support for environment variable of the API key
- Customizable caching durations
- Purging of outdated cache results

## Example usage

```golang
package main

import (
    "fmt"
    "time"

    "github.com/bradberger/forecast"
)


func main() {
    f := forecast.New("<my-api-key>", forecast.Units("us"), forecast.Language("en"))

    result, err := f.Get(51.5034070, -0.1275920, time.Now())
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s: %s\n", result.Timezone, result.Currently.Summary)
    fmt.Printf("humidity: %.2f\n", result.Currently.Humidity)
    fmt.Printf("temperature: %.2f F\n", result.Currently.Temperature)
    fmt.Printf("wind speed: %.2f mph\n", result.Currently.WindSpeed)
}
```

For kicks, you can also install it as a binary on the command-line:

`go install github.com/bradberger/forecast/forecast`

For usage, refer to the help `forecast --help`:

```
~$ forecast --help
Usage of forecast:
  -key string
        Your DarkSky API key
  -language string
        Language. One of ar, be, cs, bs, de, el, en, es, fr, hr, hu, id, it, is, kw, nb, nl pl, pt, ru, sk, sr, sv, tet, tr, uk, x-pig-latin, zh, zh-tw (default "en")
  -latitude float
        The latitude of the forecast location (default 51.503407)
  -longitude float
        The longitude of the forecast location (default -0.127592)

  -units string
        Units of measurement. One of ca, si, us, uk (default "us")
```
