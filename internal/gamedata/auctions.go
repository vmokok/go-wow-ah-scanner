package gamedata

import "fmt"

func (gd *Auctions) Namespace() string {
	return dynamicNamespace
}

func (gd *Auctions) QueryPath(opts ...string) string {
	resource := fmt.Sprintf("/data/wow/connected-realm/%s/auctions", opts[0])
	return buildQueryPath(resource)
}

func (gd *Auctions) LastUpdated() string {
	return lastUpdated(gd)
}

type Auctions struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Auctions []struct {
		Bid    int64 `json:"bid"`
		Buyout int64 `json:"buyout"`
		ID     int64 `json:"id"`
		Item   struct {
			BonusLists []int64 `json:"bonus_lists"`
			Context    int64   `json:"context"`
			ID         int64   `json:"id"`
			Modifiers  []struct {
				Type  int64 `json:"type"`
				Value int64 `json:"value"`
			} `json:"modifiers"`
			PetBreedID   int64 `json:"pet_breed_id"`
			PetLevel     int64 `json:"pet_level"`
			PetQualityID int64 `json:"pet_quality_id"`
			PetSpeciesID int64 `json:"pet_species_id"`
		} `json:"item"`
		Quantity int64  `json:"quantity"`
		TimeLeft string `json:"time_left"`
	} `json:"auctions"`
	Commodities struct {
		Href string `json:"href"`
	} `json:"commodities"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
}
