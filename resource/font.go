package resource

import "embed"

var (
	//go:embed font/headline.ttf
	_headline embed.FS
)

var (
	Font = struct {
		Headline []byte
	}{
		Headline: LoadEmbedFile(_headline, "font/headline.ttf", ByteDecoder),
	}
)
