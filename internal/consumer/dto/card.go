package dto

type Cards struct {
	Cards []Card `json:"cards" bson:"cards"`
}

type Card struct {
	ID                    string   `json:"id,omitempty" bson:"id,omitempty"`
	Name                  string   `json:"name,omitempty" bson:"name,omitempty"`
	NationalPokedexNumber int      `json:"nationalPokedexNumber,omitempty" bson:"nationalPokedexNumber,omitempty"`
	ImageURL              string   `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	ImageURLHiRes         string   `json:"imageUrlHiRes,omitempty" bson:"imageUrlHiRes,omitempty"`
	Types                 []string `json:"types,omitempty" bson:"types,omitempty"`
	Supertype             string   `json:"supertype,omitempty" bson:"supertype,omitempty"`
	Subtype               string   `json:"subtype,omitempty" bson:"subtype,omitempty"`
	Ability               struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
		Text string `json:"text,omitempty" bson:"text,omitempty"`
		Type string `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"ability,omitempty" bson:"ability,omitempty"`
	AncientTrait struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
		Text string `json:"text,omitempty" bson:"text,omitempty"`
	} `json:"ancientTrait,omitempty" bson:"ancientTrait,omitempty"`
	Hp                   interface{} `json:"hp,omitempty" bson:"hp,omitempty"`
	RetreatCost          []string    `json:"retreatCost,omitempty" bson:"retreatCost,omitempty"`
	ConvertedRetreatCost int         `json:"convertedRetreatCost,omitempty" bson:"convertedRetreatCost,omitempty"`
	Number               string      `json:"number,omitempty" bson:"number,omitempty"`
	Artist               string      `json:"artist,omitempty" bson:"artist,omitempty"`
	Rarity               string      `json:"rarity,omitempty" bson:"rarity,omitempty"`
	Series               string      `json:"series,omitempty" bson:"series,omitempty"`
	Set                  string      `json:"set,omitempty" bson:"set,omitempty"`
	SetCode              string      `json:"setCode,omitempty" bson:"setCode,omitempty"`
	Text                 []string    `json:"text,omitempty" bson:"text,omitempty"`
	Attacks              []struct {
		Cost                []string    `json:"cost,omitempty" bson:"cost,omitempty"`
		Name                string      `json:"name,omitempty" bson:"name,omitempty"`
		Text                string      `json:"text,omitempty" bson:"text,omitempty"`
		Damage              interface{} `json:"damage,omitempty" bson:"damage,omitempty"`
		ConvertedEnergyCost int         `json:"convertedEnergyCost,omitempty" bson:"convertedEnergyCost,omitempty"`
	} `json:"attacks,omitempty" bson:"attacks,omitempty"`
	Weaknesses []struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"weaknesses,omitempty"`
	EvolvesFrom string `json:"evolvesFrom,omitempty"`
}
