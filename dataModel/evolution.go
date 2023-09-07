package dataModel

type Evolution struct {
	ID              int         `json:"id"`
	BabyTriggerItem interface{} `json:"baby_trigger_item"`
	Chain           struct {
		IsBaby  bool `json:"is_baby"`
		Species struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"species"`
		EvolutionDetails interface{} `json:"evolution_details"`
		EvolvesTo        []struct {
			IsBaby  bool `json:"is_baby"`
			Species struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"species"`
			EvolutionDetails []struct {
				Item    interface{} `json:"item"`
				Trigger struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"trigger"`
				Gender                interface{} `json:"gender"`
				HeldItem              interface{} `json:"held_item"`
				KnownMove             interface{} `json:"known_move"`
				KnownMoveType         interface{} `json:"known_move_type"`
				Location              interface{} `json:"location"`
				MinLevel              int         `json:"min_level"`
				MinHappiness          interface{} `json:"min_happiness"`
				MinBeauty             interface{} `json:"min_beauty"`
				MinAffection          interface{} `json:"min_affection"`
				NeedsOverworldRain    bool        `json:"needs_overworld_rain"`
				PartySpecies          interface{} `json:"party_species"`
				PartyType             interface{} `json:"party_type"`
				RelativePhysicalStats interface{} `json:"relative_physical_stats"`
				TimeOfDay             string      `json:"time_of_day"`
				TradeSpecies          interface{} `json:"trade_species"`
				TurnUpsideDown        bool        `json:"turn_upside_down"`
			} `json:"evolution_details"`
			EvolvesTo []interface{} `json:"evolves_to"`
		} `json:"evolves_to"`
	} `json:"chain"`
}
