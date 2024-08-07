package resource

import "embed"

var (
	//go:embed pokemon/025/config.yaml
	_pokemon025 embed.FS
)

func PokemonConfig[T any]() map[string]T {
	return map[string]T{
		"025": LoadEmbedFile(_pokemon025, "pokemon/025/config.yaml", YamlDecoder[T]),
	}
}
