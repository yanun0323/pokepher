package entity

type Stats struct {
	HP             int `yaml:"hp"`
	Attack         int `yaml:"attack"`
	Defense        int `yaml:"defense"`
	SpecialAttack  int `yaml:"special_attack" mapstructure:"special_attack"`
	SpecialDefense int `yaml:"special_defense" mapstructure:"special_defense"`
	Speed          int `yaml:"speed"`
}
