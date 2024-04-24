package gamedata

type Item struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int `json:"id"`
	Name struct {
		EnUS string `json:"en_US"`
		EsMX string `json:"es_MX"`
		PtBR string `json:"pt_BR"`
		DeDE string `json:"de_DE"`
		EnGB string `json:"en_GB"`
		EsES string `json:"es_ES"`
		FrFR string `json:"fr_FR"`
		ItIT string `json:"it_IT"`
		RuRU string `json:"ru_RU"`
		KoKR string `json:"ko_KR"`
		ZhTW string `json:"zh_TW"`
		ZhCN string `json:"zh_CN"`
	} `json:"name"`
	Quality struct {
		Type string `json:"type"`
		Name struct {
			EnUS string `json:"en_US"`
			EsMX string `json:"es_MX"`
			PtBR string `json:"pt_BR"`
			DeDE string `json:"de_DE"`
			EnGB string `json:"en_GB"`
			EsES string `json:"es_ES"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			RuRU string `json:"ru_RU"`
			KoKR string `json:"ko_KR"`
			ZhTW string `json:"zh_TW"`
			ZhCN string `json:"zh_CN"`
		} `json:"name"`
	} `json:"quality"`
	Level         int `json:"level"`
	RequiredLevel int `json:"required_level"`
	Media         struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	ItemClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name struct {
			EnUS string `json:"en_US"`
			EsMX string `json:"es_MX"`
			PtBR string `json:"pt_BR"`
			DeDE string `json:"de_DE"`
			EnGB string `json:"en_GB"`
			EsES string `json:"es_ES"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			RuRU string `json:"ru_RU"`
			KoKR string `json:"ko_KR"`
			ZhTW string `json:"zh_TW"`
			ZhCN string `json:"zh_CN"`
		} `json:"name"`
		ID int `json:"id"`
	} `json:"item_class"`
	ItemSubclass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name struct {
			EnUS string `json:"en_US"`
			EsMX string `json:"es_MX"`
			PtBR string `json:"pt_BR"`
			DeDE string `json:"de_DE"`
			EnGB string `json:"en_GB"`
			EsES string `json:"es_ES"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			RuRU string `json:"ru_RU"`
			KoKR string `json:"ko_KR"`
			ZhTW string `json:"zh_TW"`
			ZhCN string `json:"zh_CN"`
		} `json:"name"`
		ID int `json:"id"`
	} `json:"item_subclass"`
	InventoryType struct {
		Type string `json:"type"`
		Name struct {
			EnUS string `json:"en_US"`
			EsMX string `json:"es_MX"`
			PtBR string `json:"pt_BR"`
			DeDE string `json:"de_DE"`
			EnGB string `json:"en_GB"`
			EsES string `json:"es_ES"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			RuRU string `json:"ru_RU"`
			KoKR string `json:"ko_KR"`
			ZhTW string `json:"zh_TW"`
			ZhCN string `json:"zh_CN"`
		} `json:"name"`
	} `json:"inventory_type"`
	PurchasePrice int  `json:"purchase_price"`
	SellPrice     int  `json:"sell_price"`
	MaxCount      int  `json:"max_count"`
	IsEquippable  bool `json:"is_equippable"`
	IsStackable   bool `json:"is_stackable"`
	PreviewItem   struct {
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"item"`
		Quality struct {
			Type string `json:"type"`
			Name struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"name"`
		} `json:"quality"`
		Name struct {
			EnUS string `json:"en_US"`
			EsMX string `json:"es_MX"`
			PtBR string `json:"pt_BR"`
			DeDE string `json:"de_DE"`
			EnGB string `json:"en_GB"`
			EsES string `json:"es_ES"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			RuRU string `json:"ru_RU"`
			KoKR string `json:"ko_KR"`
			ZhTW string `json:"zh_TW"`
			ZhCN string `json:"zh_CN"`
		} `json:"name"`
		Media struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"media"`
		ItemClass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"name"`
			ID int `json:"id"`
		} `json:"item_class"`
		ItemSubclass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"name"`
			ID int `json:"id"`
		} `json:"item_subclass"`
		InventoryType struct {
			Type string `json:"type"`
			Name struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"name"`
		} `json:"inventory_type"`
		Binding struct {
			Type string `json:"type"`
			Name struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"name"`
		} `json:"binding"`
		Weapon struct {
			Damage struct {
				MinValue      int `json:"min_value"`
				MaxValue      int `json:"max_value"`
				DisplayString struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"display_string"`
				DamageClass struct {
					Type string `json:"type"`
					Name struct {
						EnUS string `json:"en_US"`
						EsMX string `json:"es_MX"`
						PtBR string `json:"pt_BR"`
						DeDE string `json:"de_DE"`
						EnGB string `json:"en_GB"`
						EsES string `json:"es_ES"`
						FrFR string `json:"fr_FR"`
						ItIT string `json:"it_IT"`
						RuRU string `json:"ru_RU"`
						KoKR string `json:"ko_KR"`
						ZhTW string `json:"zh_TW"`
						ZhCN string `json:"zh_CN"`
					} `json:"name"`
				} `json:"damage_class"`
			} `json:"damage"`
			AttackSpeed struct {
				Value         int `json:"value"`
				DisplayString struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"display_string"`
			} `json:"attack_speed"`
			Dps struct {
				Value         float64 `json:"value"`
				DisplayString struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"display_string"`
			} `json:"dps"`
		} `json:"weapon"`
		Stats []struct {
			Type struct {
				Type string `json:"type"`
				Name struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"name"`
			} `json:"type"`
			Value   int `json:"value"`
			Display struct {
				DisplayString struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"display_string"`
				Color struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
			} `json:"display"`
		} `json:"stats"`
		Spells []struct {
			Spell struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					ItIT string `json:"it_IT"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"name"`
				ID int `json:"id"`
			} `json:"spell"`
			Description struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"description"`
		} `json:"spells"`
		SellPrice struct {
			Value          int `json:"value"`
			DisplayStrings struct {
				Header struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"header"`
				Gold struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"gold"`
				Silver struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"silver"`
				Copper struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"copper"`
			} `json:"display_strings"`
		} `json:"sell_price"`
		Requirements struct {
			Level struct {
				Value         int `json:"value"`
				DisplayString struct {
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					DeDE string `json:"de_DE"`
					EnGB string `json:"en_GB"`
					EsES string `json:"es_ES"`
					FrFR string `json:"fr_FR"`
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					KoKR string `json:"ko_KR"`
					ZhTW string `json:"zh_TW"`
					ZhCN string `json:"zh_CN"`
				} `json:"display_string"`
			} `json:"level"`
		} `json:"requirements"`
		Level struct {
			Value         int `json:"value"`
			DisplayString struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"display_string"`
		} `json:"level"`
		Durability struct {
			Value         int `json:"value"`
			DisplayString struct {
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				DeDE string `json:"de_DE"`
				EnGB string `json:"en_GB"`
				EsES string `json:"es_ES"`
				FrFR string `json:"fr_FR"`
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				KoKR string `json:"ko_KR"`
				ZhTW string `json:"zh_TW"`
				ZhCN string `json:"zh_CN"`
			} `json:"display_string"`
		} `json:"durability"`
	} `json:"preview_item"`
	PurchaseQuantity int `json:"purchase_quantity"`
}
