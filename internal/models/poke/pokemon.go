package poke

type (
	SearchPokemonResponse struct {
		ID             int                `json:"id"`
		Name           string             `json:"name"`
		BaseExperience int                `json:"base_experience"`
		Height         int                `json:"height"`
		Weight         int                `json:"weight"`
		Order          int                `json:"order"`
		Abilities      []PokemonAbilities `json:"abilities"`
		Moves          []PokemonMoves     `json:"moves"`
		Species        PokemonSpecies     `json:"species"`
		Stats          []PokemonStats     `json:"stats"`
		Types          []PokemonType      `json:"types"`
	}

	PokemonAbilities struct {
		IsHidden bool           `json:"is_hidden"`
		Slot     int            `json:"slot"`
		Ability  AbilityDetails `json:"ability"`
	}

	AbilityDetails struct {
		Name string `json:"name"`
	}

	PokemonMoves struct {
		Move MoveDetails `json:"move"`
	}

	MoveDetails struct {
		Name string `json:"name"`
	}

	PokemonSpecies struct {
		Name string `json:"name"`
	}

	PokemonStats struct {
		BaseStat int         `json:"base_stat"`
		Effort   int         `json:"effort"`
		Stat     StatDetails `json:"stat"`
	}

	StatDetails struct {
		Name string `json:"name"`
	}

	PokemonType struct {
		Slot int         `json:"slot"`
		Type TypeDetails `json:"type"`
	}

	TypeDetails struct {
		Name string `json:"name"`
	}
)
