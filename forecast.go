package forecast

import (
	"encoding/json"
    "fmt"
	"strconv"
	"time"

    "golang.org/x/net/context"
)

// URL example:  "https://api.forecast.io/forecast/APIKEY/LATITUDE,LONGITUDE,TIME?units=ca"
const (
	BaseURL = "https://api.forecast.io/forecast"
	CacheDuration = 15 * time.Minute
)

type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable"`
	DarkSkyStations    []string `json:"darksky-stations"`
	DataPointStations  []string `json:"datapoint-stations"`
	ISDStations        []string `json:"isds-stations"`
	LAMPStations       []string `json:"lamp-stations"`
	METARStations      []string `json:"metars-stations"`
	METNOLicense       string   `json:"metnol-license"`
	Sources            []string `json:"sources"`
	Units              string   `json:"units"`
}

type DataPoint struct {
	Time                   float64 `json:"time"`
	Summary                string  `json:"summary"`
	Icon                   string  `json:"icon"`
	SunriseTime            float64 `json:"sunriseTime"`
	SunsetTime             float64 `json:"sunsetTime"`
	PrecipIntensity        float64 `json:"precipIntensity"`
	PrecipIntensityMax     float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime float64 `json:"precipIntensityMaxTime"`
	PrecipProbability      float64 `json:"precipProbability"`
	PrecipType             string  `json:"precipType"`
	PrecipAccumulation     float64 `json:"precipAccumulation"`
	Temperature            float64 `json:"temperature"`
	TemperatureMin         float64 `json:"temperatureMin"`
	TemperatureMinTime     float64 `json:"temperatureMinTime"`
	TemperatureMax         float64 `json:"temperatureMax"`
	TemperatureMaxTime     float64 `json:"temperatureMaxTime"`
	ApparentTemperature    float64 `json:"apparentTemperature"`
	DewPoint               float64 `json:"dewPoint"`
	WindSpeed              float64 `json:"windSpeed"`
	WindBearing            float64 `json:"windBearing"`
	CloudCover             float64 `json:"cloudCover"`
	Humidity               float64 `json:"humidity"`
	Pressure               float64 `json:"pressure"`
	Visibility             float64 `json:"visibility"`
	Ozone                  float64 `json:"ozone"`
	MoonPhase              float64 `json:"moonPhase"`
}

type DataBlock struct {
	Summary string      `json:"summary"`
	Icon    string      `json:"icon"`
	Data    []DataPoint `json:"data"`
}

type Alert struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Time        float64 `json:"time"`
	Expires     float64 `json:"expires"`
	URI         string  `json:"uri"`
}

type Forecast struct {
    Latitude  float64    `json:"latitude"`
    Longitude float64    `json:"longitude"`
    Timezone  string     `json:"timezone"`
    Offset    float64    `json:"offset"`
    Currently DataPoint  `json:"currently"`
    Minutely  DataBlock  `json:"minutely"`
    Hourly    DataBlock  `json:"hourly"`
    Daily     DataBlock  `json:"daily"`
    Alerts    []Alert    `json:"Alerts"`
    Flags     Flags      `json:"flags"`
    Code      int        `json:"code"`
}

type Client struct {

    APICalls  int

    key       string
    units     Units
    language  Language

	// Results are index by <time, nearest 15 minutes> | <lat>,<lng> | lang | units
    results   map[string]map[string]map[string]map[string]*Forecast

    ctx       context.Context
}

func New(key string, units Units, lang Language) *Client {
    return &Client{
        key: key,
        units: units,
        language: lang,
    }
}

func (c *Client) getURL(lat float64, lng float64, t time.Time) string {
	return fmt.Sprintf("%s/%s/%f,%f?units=%s&time=%d&lang=%s", BaseURL, c.key, lat, lng, c.units, t.Unix(), c.language)
}

func (c *Client) Get(lat float64, long float64, t time.Time) (*Forecast, error) {

    result := c.results[t.Truncate(CacheDuration).String()][fmt.Sprintf("%.4f,%.4f", lat, long)][string(c.language)][string(c.units)]
    if result != nil {
        return result, nil
    }

    result = &Forecast{}
	res, err := c.fetch(lat, long, t)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(result); err != nil {
		return nil, err
	}

	calls, _ := strconv.Atoi(res.Header.Get("X-Forecast-API-Calls"))
	c.APICalls = calls

	return result, nil
}
