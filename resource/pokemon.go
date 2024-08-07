package resource

import "embed"

var (
	//go:embed pokemon/025/config.yaml
	_pokemon025 embed.FS
)

func PokemonConfig[T any]() map[int]T {
	return map[int]T{
		25: LoadEmbedFile(_pokemon025, "pokemon/025/config.yaml", YamlDecoder[T]),
	}
}
