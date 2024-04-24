package gamedata

func (gd *ConnectedRealmIndex) Namespace() string {
	return dynamicNamespace
}

func (gd *ConnectedRealmIndex) QueryPath(opts ...string) string {
	resource := "/data/wow/connected-realm/index"
	return buildQueryPath(resource)
}

func (gd *ConnectedRealmIndex) LastUpdated() string {
	return lastUpdated(gd)
}

type ConnectedRealmIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ConnectedRealms []struct {
		Href string `json:"href"`
	} `json:"connected_realms"`
}
