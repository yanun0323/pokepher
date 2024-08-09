package resource

import "embed"

var (
	//go:embed font/headline.ttf
	_headline embed.FS

	//go:embed font/regular.ttf
	_regular embed.FS
)

var (
	Font = struct {
		Headline []byte
		Regular  []byte
	}{
		Headline: LoadEmbedFile(_headline, "font/headline.ttf", ByteDecoder),
		Regular:  LoadEmbedFile(_regular, "font/regular.ttf", ByteDecoder),
	}
)
