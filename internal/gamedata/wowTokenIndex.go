package gamedata

type WowTokenIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp"`
	Price                int64 `json:"price"`
}
