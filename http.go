// +build !appengine

package forecast

import (
    "net/http"
    "time"
)

func (c *Client) fetch(lat float64, lng float64, t time.Time) (*http.Response, error) {
    return http.Get(c.getURL(lat, lng, t))
}
