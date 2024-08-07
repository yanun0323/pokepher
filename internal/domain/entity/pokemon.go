package entity

type Pokemon struct {
	ID               ID
	Level            int
	Name             string
	Nickname         string
	Height           float64
	Weight           float64
	Type             [2]Type
	Ability          Ability
	Gender           Gender
	Nature           Nature
	Friendly         int
	Evolution        map[EvolutionRule]ID
	Base             Stats
	EffortValue      Stats
	IndividualValues Stats

	MaxHP int
	HP    int

	MaxExp int
	Exp    int

	Status        Status
	StatsAddition Stats

	Skills [4]Skill
}

func NewPokemon(opt PokemonOption) *Pokemon {
	cfg, ok := PokemonConfigTable[opt.ID.String()]
	if !ok {
		return nil
	}

	return &Pokemon{
		ID:               cfg.ID,
		Level:            opt.Level,
		Name:             cfg.Name,
		Nickname:         cfg.Name,
		Height:           cfg.Height,
		Weight:           cfg.Weight,
		Type:             cfg.Type,
		Ability:          cfg.PickAbility(opt.AbilityIdx, opt.HiddenAbility),
		Gender:           opt.Gender,
		Nature:           opt.Nature,
		Friendly:         cfg.Breeding.HatchTime,
		Evolution:        cfg.Evolution,
		Base:             cfg.Base,
		EffortValue:      Stats{},
		IndividualValues: opt.IndividualValues,
		MaxHP:            cfg.Base.HP,
		HP:               cfg.Base.HP,
		MaxExp:           cfg.Experience,
		Exp:              cfg.Experience,
		Status:           StatusNone,
		StatsAddition:    Stats{},
		Skills:           cfg.PickSkill(opt.Level),
	}
}

type PokemonOption struct {
	ID               ID
	Level            int
	AbilityIdx       int
	HiddenAbility    bool
	Gender           Gender
	Nature           Nature
	IndividualValues Stats
}
