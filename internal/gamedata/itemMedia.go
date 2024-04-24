package gamedata

type ItemMedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		FileDataID int    `json:"file_data_id"`
	} `json:"assets"`
	ID int `json:"id"`
}
