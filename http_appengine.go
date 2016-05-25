// +build appengine

package forecast

import (
    "fmt"
    "net/http"
    "time"

    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
)

func Init(r *http.Request, key string, units Units, lang Language) *Client {
    return &Client{
        key: key,
        units: units,
        language: lang,
        ctx: appengine.NewContext(r),
    }
}

func (c *Client) Request(r *http.Request) *Client {
    c.ctx = appengine.NewContext(r)
    return c
}

func (c *Client) fetch(lat float64, lng float64, t time.Time) (*http.Response, error) {
	return client.Get(urlfetch.Client(c.ctx).GetURL(lat, lng, t))
}
