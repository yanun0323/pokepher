package entity

type Ability string

func NewAbilities(abilities []string) []Ability {
	result := make([]Ability, 0, len(abilities))
	for _, a := range abilities {
		result = append(result, Ability(a))
	}

	return result
}
