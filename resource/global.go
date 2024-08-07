package resource

import (
	"embed"
	_ "image/gif"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed image/global/arrow_down.png
	_arrow_down embed.FS

	//go:embed image/global/emotes.png
	_emotes embed.FS

	//go:embed image/global/pokeball_wiggle_sheet.png
	_pokeball_wiggle_sheet embed.FS

	//go:embed image/global/pokeball.png
	_pokeball embed.FS

	//go:embed image/global/symbol_female.png
	_symbol_female embed.FS

	//go:embed image/global/symbol_male.png
	_symbol_male embed.FS
)

var Global = struct {
	ArrowDown           *ebiten.Image
	Emotes              *ebiten.Image
	PokeballWiggleSheet *ebiten.Image
	Pokeball            *ebiten.Image
	SymbolFemale        *ebiten.Image
	SymbolMale          *ebiten.Image
}{
	ArrowDown:           LoadEmbedFile(_arrow_down, "image/global/arrow_down.png", PngDecoder),
	Emotes:              LoadEmbedFile(_emotes, "image/global/emotes.png", PngDecoder),
	PokeballWiggleSheet: LoadEmbedFile(_pokeball_wiggle_sheet, "image/global/pokeball_wiggle_sheet.png", PngDecoder),
	Pokeball:            LoadEmbedFile(_pokeball, "image/global/pokeball.png", PngDecoder),
	SymbolFemale:        LoadEmbedFile(_symbol_female, "image/global/symbol_female.png", PngDecoder),
	SymbolMale:          LoadEmbedFile(_symbol_male, "image/global/symbol_male.png", PngDecoder),
}
