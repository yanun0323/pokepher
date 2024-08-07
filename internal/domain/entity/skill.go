package entity

import "main/resource"

var (
	SkillTable = resource.Skill[map[string]Skill]()
)

type Skill struct {
	Name     string   `yaml:"name"`
	Type     Type     `yaml:"type"`
	Power    int      `yaml:"power"`
	Category Category `yaml:"category"`
	Accuracy int      `yaml:"accuracy"`
}
