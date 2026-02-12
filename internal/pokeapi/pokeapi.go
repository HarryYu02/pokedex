package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/harryyu02/pokedex/internal/pokecache"
)

type PokeApiClient struct {
	Cache *pokecache.Cache
}

type LocationAreaRes struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonInAreaRes struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type PokemonRes struct {
	Abilities []Abilities `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries Cries `json:"cries"`
	Forms []Forms `json:"forms"`
	GameIndices []GameIndices `json:"game_indices"`
	Height int `json:"height"`
	HeldItems []HeldItems `json:"held_items"`
	ID int `json:"id"`
	IsDefault bool `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves []Moves `json:"moves"`
	Name string `json:"name"`
	Order int `json:"order"`
	PastAbilities []PastAbilities `json:"past_abilities"`
	PastStats []PastStats `json:"past_stats"`
	PastTypes []any `json:"past_types"`
	Species Species `json:"species"`
	Sprites Sprites `json:"sprites"`
	Stats0 []Stats0 `json:"stats"`
	Types []Types `json:"types"`
	Weight int `json:"weight"`
}
type Ability struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Abilities struct {
	Ability Ability `json:"ability"`
	IsHidden bool `json:"is_hidden"`
	Slot int `json:"slot"`
}
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
type Forms struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type GameIndices struct {
	GameIndex int `json:"game_index"`
	Version Version `json:"version"`
}
type Item struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type VersionDetailsVersion struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type VersionDetails struct {
	Rarity int `json:"rarity"`
	VersionDetailsVersion VersionDetailsVersion `json:"version"`
}
type HeldItems struct {
	Item Item `json:"item"`
	VersionDetails []VersionDetails `json:"version_details"`
}
type Move struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type VersionGroup struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type VersionGroupDetails struct {
	LevelLearnedAt int `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	Order any `json:"order"`
	VersionGroup VersionGroup `json:"version_group"`
}
type Moves struct {
	Move Move `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type PastAbilitiesAbilities struct {
	Ability any `json:"ability"`
	IsHidden bool `json:"is_hidden"`
	Slot int `json:"slot"`
}
type Generation struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type PastAbilities struct {
	PastAbilitiesAbilities []PastAbilitiesAbilities `json:"abilities"`
	Generation Generation `json:"generation"`
}
type PastStatsGeneration struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Stat struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Stats struct {
	BaseStat int `json:"base_stat"`
	Effort int `json:"effort"`
	Stat Stat `json:"stat"`
}
type PastStats struct {
	PastStatsGeneration PastStatsGeneration `json:"generation"`
	Stats []Stats `json:"stats"`
}
type Species struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type DreamWorld struct {
	FrontDefault string `json:"front_default"`
	FrontFemale any `json:"front_female"`
}
type Home struct {
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
}
type Showdown struct {
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale any `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type Other struct {
	DreamWorld DreamWorld `json:"dream_world"`
	Home Home `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown Showdown `json:"showdown"`
}
type RedBlue struct {
	BackDefault string `json:"back_default"`
	BackGray string `json:"back_gray"`
	BackTransparent string `json:"back_transparent"`
	FrontDefault string `json:"front_default"`
	FrontGray string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}
type Yellow struct {
	BackDefault string `json:"back_default"`
	BackGray string `json:"back_gray"`
	BackTransparent string `json:"back_transparent"`
	FrontDefault string `json:"front_default"`
	FrontGray string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}
type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow Yellow `json:"yellow"`
}
type Crystal struct {
	BackDefault string `json:"back_default"`
	BackShiny string `json:"back_shiny"`
	BackShinyTransparent string `json:"back_shiny_transparent"`
	BackTransparent string `json:"back_transparent"`
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyTransparent string `json:"front_shiny_transparent"`
	FrontTransparent string `json:"front_transparent"`
}
type Gold struct {
	BackDefault string `json:"back_default"`
	BackShiny string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}
type Silver struct {
	BackDefault string `json:"back_default"`
	BackShiny string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}
type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold Gold `json:"gold"`
	Silver Silver `json:"silver"`
}
type Emerald struct {
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
}
type FireredLeafgreen struct {
	BackDefault string `json:"back_default"`
	BackShiny string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
}
type RubySapphire struct {
	BackDefault string `json:"back_default"`
	BackShiny string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
}
type GenerationIii struct {
	Emerald Emerald `json:"emerald"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
	RubySapphire RubySapphire `json:"ruby-sapphire"`
}
type DiamondPearl struct {
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale string `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type HeartgoldSoulsilver struct {
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale string `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type Platinum struct {
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale string `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type GenerationIv struct {
	DiamondPearl DiamondPearl `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
	Platinum Platinum `json:"platinum"`
}
type ScarletViolet struct {
	FrontDefault string `json:"front_default"`
	FrontFemale any `json:"front_female"`
}
type GenerationIx struct {
	ScarletViolet ScarletViolet `json:"scarlet-violet"`
}
type Animated struct {
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale string `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type BlackWhite struct {
	Animated Animated `json:"animated"`
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale string `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}
type OmegarubyAlphasapphire struct {
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type XY struct {
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
	XY XY `json:"x-y"`
}
type Icons struct {
	FrontDefault string `json:"front_default"`
	FrontFemale any `json:"front_female"`
}
type UltraSunUltraMoon struct {
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
type GenerationVii struct {
	Icons Icons `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}
type BrilliantDiamondShiningPearl struct {
	FrontDefault string `json:"front_default"`
	FrontFemale any `json:"front_female"`
}
type GenerationViiiIcons struct {
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
}
type GenerationViii struct {
	BrilliantDiamondShiningPearl BrilliantDiamondShiningPearl `json:"brilliant-diamond-shining-pearl"`
	GenerationViiiIcons GenerationViiiIcons `json:"icons"`
}
type Versions struct {
	GenerationI GenerationI `json:"generation-i"`
	GenerationIi GenerationIi `json:"generation-ii"`
	GenerationIii GenerationIii `json:"generation-iii"`
	GenerationIv GenerationIv `json:"generation-iv"`
	GenerationIx GenerationIx `json:"generation-ix"`
	GenerationV GenerationV `json:"generation-v"`
	GenerationVi GenerationVi `json:"generation-vi"`
	GenerationVii GenerationVii `json:"generation-vii"`
	GenerationViii GenerationViii `json:"generation-viii"`
}
type Sprites struct {
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale string `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
	Other Other `json:"other"`
	Versions Versions `json:"versions"`
}
type Stats0Stat struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Stats0 struct {
	BaseStat int `json:"base_stat"`
	Effort int `json:"effort"`
	Stats0Stat Stats0Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Types struct {
	Slot int `json:"slot"`
	Type Type `json:"type"`
}

func NewClient(interval time.Duration) *PokeApiClient {
	return &PokeApiClient{
		Cache: pokecache.NewCache(interval),
	}
}

func (p *PokeApiClient)GetLocationAreas(url string) (LocationAreaRes, error) {
	bytes := []byte{}
	cached, ok := p.Cache.Get(url)
	if ok {
		bytes = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreaRes{}, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return LocationAreaRes{}, err
		}
		bytes = body
	}

	var locationAreas LocationAreaRes
	err := json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return LocationAreaRes{}, err
	}

	p.Cache.Add(url, bytes)

	return locationAreas, nil
}

func (p *PokeApiClient)GetPokemonInArea(url string) (PokemonInAreaRes, error) {
	bytes := []byte{}
	cached, ok := p.Cache.Get(url)
	if ok {
		bytes = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return PokemonInAreaRes{}, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return PokemonInAreaRes{}, err
		}
		if string(body) == "Not Found" {
			return PokemonInAreaRes{}, fmt.Errorf("no pokemon found in given area")
		}
		bytes = body
	}

	var pokemonInArea PokemonInAreaRes
	err := json.Unmarshal(bytes, &pokemonInArea)
	if err != nil {
		return PokemonInAreaRes{}, err
	}

	p.Cache.Add(url, bytes)

	return pokemonInArea, nil
}

func (p *PokeApiClient)GetPokemon(url string) (PokemonRes, error) {
	bytes := []byte{}
	cached, ok := p.Cache.Get(url)
	if ok {
		bytes = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return PokemonRes{}, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return PokemonRes{}, err
		}
		if string(body) == "Not Found" {
			return PokemonRes{}, fmt.Errorf("pokemon not found")
		}
		bytes = body
	}

	var pokemon PokemonRes
	err := json.Unmarshal(bytes, &pokemon)
	if err != nil {
		return PokemonRes{}, err
	}

	p.Cache.Add(url, bytes)

	return pokemon, nil
}
