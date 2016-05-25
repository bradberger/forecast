package main

import (
    "flag"
    "fmt"
    "os"
    "time"

    "github.com/bradberger/forecast"
)

var (
    lat, lng float64
    key, lang, units string
    template = `
Time          %s
Temperature   %.2f
Wind          %.2f
Humidity      %.0f%%
Pressure      %.2f
Cloud Cover   %.0f%%
Visibility    %.2f
Ozone         %.2f
Summary       %s
Units         %s
Language      %s
Latitude      %.7f
Longitude     %.7f

`
)

func init() {
    flag.StringVar(&key, "key", "", "Your DarkSky API key")
    flag.Float64Var(&lat, "latitude", 51.5034070, "The latitude of the forecast location")
    flag.Float64Var(&lng, "longitude", -0.1275920, "The longitude of the forecast location")
    flag.StringVar(&units, "units", string(forecast.US), "Units of measurement. One of ca, si, us, uk")
    flag.StringVar(&lang, "language", string(forecast.ENGLISH), "Language. One of ar, be, cs, bs, de, el, en, es, fr, hr, hu, id, it, is, kw, nb, nl pl, pt, ru, sk, sr, sv, tet, tr, uk, x-pig-latin, zh, zh-tw")
    flag.Parse()
}

func main() {

    if key == "" {
        fmt.Println("Please specify an API key with the --key flag")
        os.Exit(2)
    }
    f := forecast.New(key, forecast.Units(units), forecast.Language(lang))
    result, err := f.Get(lat, lng, time.Now())
    if err != nil {
        fmt.Printf("\nError getting the forecast. This can be caused by an incorrect API key, network errors, etc. The message was:\n\t%v\n\n", err)
        os.Exit(1)
    }
    fmt.Printf(template, time.Now(), result.Currently.Temperature, result.Currently.WindSpeed, result.Currently.Humidity * 100, result.Currently.Pressure, result.Currently.CloudCover * 100, result.Currently.Visibility, result.Currently.Ozone, result.Currently.Summary, units, lang, result.Latitude, result.Longitude)
}
