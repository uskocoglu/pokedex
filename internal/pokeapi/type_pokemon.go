package pokeapi

// Structs to match the JSON structure
type PokemonEncounters struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon       Pokemon          `json:"pokemon"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          Version           `json:"version"`
}

type EncounterDetail struct {
	Chance          int           `json:"chance"`
	ConditionValues []interface{} `json:"condition_values"` // Assuming this can be any type
	MaxLevel        int           `json:"max_level"`
	Method          EncounterMethod `json:"method"`
	MinLevel        int           `json:"min_level"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}