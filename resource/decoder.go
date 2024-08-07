package resource

import (
	"embed"
	"image/gif"
	"image/png"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
	"github.com/yanun0323/pkg/logs"
)

func LoadEmbedFile[T any](fs embed.FS, name string, decoder func(io.Reader) T) T {
	f, err := fs.Open(name)
	if err != nil {
		logs.Fatalf("open embed file %s, err: %+v", name, err)
	}
	defer f.Close()

	logs.Info(name)
	return decoder(f)
}

func GifDecoder(r io.Reader) []*ebiten.Image {
	g, err := gif.DecodeAll(r)
	if err != nil {
		logs.Fatal("decode gif")
	}

	result := make([]*ebiten.Image, 0, len(g.Image))
	for _, img := range g.Image {
		result = append(result, ebiten.NewImageFromImage(img))
	}

	logs.Info(len(result))
	return result
}

func PngDecoder(r io.Reader) *ebiten.Image {
	g, err := png.Decode(r)
	if err != nil {
		logs.Fatal("decode png")
	}

	return ebiten.NewImageFromImage(g)
}

func PngRepeatDecoder(r io.Reader) []*ebiten.Image {
	g, err := png.Decode(r)
	if err != nil {
		logs.Fatal("decode png")
	}

	img := ebiten.NewImageFromImage(g)
	count := 60
	result := make([]*ebiten.Image, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, img)
	}

	return result
}

func ByteDecoder(r io.Reader) []byte {
	buf, err := io.ReadAll(r)
	if err != nil {
		logs.Fatal("read all")
	}
	return buf
}

func YamlDecoder[T any](r io.Reader) T {
	var result T

	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(r); err != nil {
		return result
	}

	if err := v.Unmarshal(&result); err != nil {
		return result
	}

	return result
}
