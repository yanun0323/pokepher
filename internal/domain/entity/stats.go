package entity

import "math/rand"

const (
	_maxIV = 31
)

type Stats struct {
	HP             int `yaml:"hp"`
	Attack         int `yaml:"attack"`
	Defense        int `yaml:"defense"`
	SpecialAttack  int `yaml:"special_attack" mapstructure:"special_attack"`
	SpecialDefense int `yaml:"special_defense" mapstructure:"special_defense"`
	Speed          int `yaml:"speed"`
}

func StatsRandomIV() Stats {
	return Stats{
		HP:             rand.Intn(_maxIV),
		Attack:         rand.Intn(_maxIV),
		Defense:        rand.Intn(_maxIV),
		SpecialAttack:  rand.Intn(_maxIV),
		SpecialDefense: rand.Intn(_maxIV),
		Speed:          rand.Intn(_maxIV),
	}
}
