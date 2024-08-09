package entity

import "main/resource"

var (
	PokemonConfigTable = resource.PokemonConfig[PokemonConfig]()
)

type PokemonConfig struct {
	ID          ID      `yaml:"id"`
	Enable      bool    `yaml:"enable"`
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Height      float64 `yaml:"height"`
	Weight      float64 `yaml:"weight"`
	Type        [2]Type `yaml:"type,flow"`
	Ability     struct {
		Abilities       []Ability `yaml:"abilities,flow"`
		HiddenAbilities []Ability `yaml:"hidden_abilities,flow" mapstructure:"hidden_abilities"`
	} `yaml:"ability"`

	GenderRatio float64 `yaml:"gender_ratio" mapstructure:"gender_ratio"`
	CacheRate   int     `yaml:"cache_rate" mapstructure:"cache_rate"`
	EV          int     `yaml:"ev"`

	Breeding struct {
		EggGroup  []EggGroup `yaml:"egg_group,flow" mapstructure:"egg_group"`
		HatchTime int        `yaml:"hatch_time" mapstructure:"hatch_time"`
	} `yaml:"breeding"`

	Base Stats `yaml:"base"`

	Experience int `yaml:"experience"`

	Evolution map[EvolutionRule]ID `yaml:"evolution"`

	SkillSheets map[int][]string `yaml:"skill_sheets" mapstructure:"skill_sheets"`
	LearnSheets []string         `yaml:"learn_sheets,flow" mapstructure:"learn_sheets"`
}

func (cfg PokemonConfig) PickAbility(abilityIdx int, hiddenAbility bool) Ability {
	if hiddenAbility && len(cfg.Ability.HiddenAbilities) != 0 {
		return Ability(cfg.Ability.HiddenAbilities[abilityIdx%len(cfg.Ability.HiddenAbilities)])
	}

	return Ability(cfg.Ability.Abilities[abilityIdx%len(cfg.Ability.Abilities)])
}

func (cfg PokemonConfig) PickSkill(level int) [4]Skill {
	result := [4]Skill{}
	idx := 0
	for ll := level; ll >= 0 && idx < 4; ll-- {
		skills, ok := cfg.SkillSheets[ll]
		if !ok {
			continue
		}

		for _, skill := range skills {
			result[idx] = SkillTable[skill]
			idx++
		}
	}

	return result
}
