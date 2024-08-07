package resource

import "embed"

var (
	//go:embed skill/config.yaml
	_skill_config embed.FS
)

func Skill[T any]() T {
	return LoadEmbedFile(_skill_config, "skill/config.yaml", YamlDecoder[T])
}
