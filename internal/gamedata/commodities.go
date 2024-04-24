package gamedata

type Commodities struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Auctions []struct {
		ID   int64 `json:"id"`
		Item struct {
			ID int64 `json:"id"`
		} `json:"item"`
		Quantity  int64  `json:"quantity"`
		TimeLeft  string `json:"time_left"`
		UnitPrice int64  `json:"unit_price"`
	} `json:"auctions"`
}
