package resource

import (
	"embed"
	_ "image/gif"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed image/battle/ui_status.png
	_ui_status embed.FS
)

var (
	Battle = struct {
		UIStatus *ebiten.Image
	}{
		UIStatus: LoadEmbedFile(_ui_status, "image/battle/ui_status.png", PngDecoder),
	}
)
