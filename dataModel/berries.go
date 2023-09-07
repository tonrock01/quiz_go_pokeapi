package dataModel

type Berries struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	GrowthTime       int    `json:"growth_time"`
	MaxHarvest       int    `json:"max_harvest"`
	NaturalGiftPower int    `json:"natural_gift_power"`
	Size             int    `json:"size"`
	Smoothness       int    `json:"smoothness"`
	SoilDryness      int    `json:"soil_dryness"`
	Firmness         struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"firmness"`
	Flavors []struct {
		Potency int `json:"potency"`
		Flavor  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"flavor"`
	} `json:"flavors"`
	Item struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"item"`
	NaturalGiftType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"natural_gift_type"`
}
