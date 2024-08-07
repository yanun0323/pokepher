package resource

import (
	"embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed pokemon/025/front.png
	_pokemon025Front embed.FS

	//go:embed pokemon/025/back.png
	_pokemon025Back embed.FS
)

func LoadPokemonImage(id string) (*PokemonImage, error) {
	front, err := SafeLoadEmbedFile(_pokemon025Front, fmt.Sprintf("pokemon/%s/front.png", id), PngDecoder)
	if err != nil {
		return nil, fmt.Errorf("failed to load front image: %w", err)
	}

	back, err := SafeLoadEmbedFile(_pokemon025Back, fmt.Sprintf("pokemon/%s/back.png", id), PngDecoder)
	if err != nil {
		return nil, fmt.Errorf("failed to load back image: %w", err)
	}

	return &PokemonImage{
		Front: front,
		Back:  back,
	}, nil
}

type PokemonImage struct {
	Front *ebiten.Image
	Back  *ebiten.Image
}
