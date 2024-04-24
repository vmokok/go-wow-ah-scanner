package gamedata

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	staticNamespace  = "static-eu"
	dynamicNamespace = "dynamic-eu"
	apiURL           = "https://eu.api.blizzard.com"
)

type Gd interface {
	Namespace() string
	QueryPath(opts ...string) string
	LastUpdated() string
}

func buildQueryPath(resource string, opts ...string) string {
	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	data := url.Values{}
	if len(opts) == 0 {
		return u.String()
	}
	data.Set("name.en_GB", strings.Join(opts, "&"))
	queryPath := "?" + data.Encode()
	return u.String() + queryPath
}

func lastUpdated(t Gd) string {
	tn := reflect.TypeOf(t).Elem().Name()
	//TODO: store values
	lu := map[string]string{
		"itemSearch":          time.Now().AddDate(0, 0, -7).Format(http.TimeFormat),
		"ConnectedRealmIndex": time.Now().AddDate(0, 0, -7).Format(http.TimeFormat),
		"Auctions":            time.Now().AddDate(0, 0, -7).Format(http.TimeFormat),
	}
	return lu[tn]
}

func Map(body []byte, t Gd) error {
	err := json.Unmarshal(body, t)
	if err != nil {
		return err
	}
	return nil
}
