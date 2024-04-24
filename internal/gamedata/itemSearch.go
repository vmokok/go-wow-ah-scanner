package gamedata

func (gd *ItemSearch) Namespace() string {
	return staticNamespace
}

func (gd *ItemSearch) QueryPath(opts ...string) string {
	resource := "/data/wow/search/item"
	return buildQueryPath(resource, opts...)
}

func (gd *ItemSearch) LastUpdated() string {
	return lastUpdated(gd)
}

type ItemSearch struct {
	Page              int  `json:"page"`
	PageSize          int  `json:"pageSize"`
	MaxPageSize       int  `json:"maxPageSize"`
	PageCount         int  `json:"pageCount"`
	ResultCountCapped bool `json:"resultCountCapped"`
	Results           []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Data struct {
			Level         int `json:"level"`
			RequiredLevel int `json:"required_level"`
			SellPrice     int `json:"sell_price"`
			ItemSubclass  struct {
				Name struct {
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				ID int `json:"id"`
			} `json:"item_subclass"`
			IsEquippable     bool `json:"is_equippable"`
			PurchaseQuantity int  `json:"purchase_quantity"`
			Media            struct {
				ID int `json:"id"`
			} `json:"media"`
			ItemClass struct {
				Name struct {
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				ID int `json:"id"`
			} `json:"item_class"`
			Quality struct {
				Name struct {
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				Type string `json:"type"`
			} `json:"quality"`
			MaxCount    int  `json:"max_count"`
			IsStackable bool `json:"is_stackable"`
			Name        struct {
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				EnGB string `json:"en_GB"`
				ZhTW string `json:"zh_TW"`
				KoKR string `json:"ko_KR"`
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				EsES string `json:"es_ES"`
				ZhCN string `json:"zh_CN"`
				FrFR string `json:"fr_FR"`
				DeDE string `json:"de_DE"`
			} `json:"name"`
			PurchasePrice int `json:"purchase_price"`
			ID            int `json:"id"`
			InventoryType struct {
				Name struct {
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				Type string `json:"type"`
			} `json:"inventory_type"`
		} `json:"data"`
	} `json:"results"`
}
