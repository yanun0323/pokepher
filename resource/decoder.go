package resource

import (
	"embed"
	"fmt"
	"image/gif"
	"image/png"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
	"github.com/yanun0323/pkg/logs"
)

func CheckEmbedFile(fs embed.FS, name string) bool {
	_, err := fs.Open(name)
	return err == nil
}

func LoadEmbedFile[T any](fs embed.FS, name string, decoder func(io.Reader) (T, error)) T {
	f, err := fs.Open(name)
	if err != nil {
		logs.Fatalf("open embed file %s, err: %+v", name, err)
	}
	defer f.Close()

	logs.Info(name)
	result, err := decoder(f)
	if err != nil {
		logs.Fatalf("decode embed file %s, err: %+v", name, err)
	}

	return result
}

func SafeLoadEmbedFile[T any](fs embed.FS, name string, decoder func(io.Reader) (T, error)) (T, error) {
	var zero T
	f, err := fs.Open(name)
	if err != nil {
		return zero, fmt.Errorf("open embed file %s, err: %w", name, err)
	}
	defer f.Close()

	logs.Info(name)
	result, err := decoder(f)
	if err != nil {
		return zero, fmt.Errorf("decode embed file %s, err: %w", name, err)
	}

	return result, nil
}

func GifDecoder(r io.Reader) ([]*ebiten.Image, error) {
	g, err := gif.DecodeAll(r)
	if err != nil {
		return nil, fmt.Errorf("decode gif, err: %w", err)
	}

	result := make([]*ebiten.Image, 0, len(g.Image))
	for _, img := range g.Image {
		result = append(result, ebiten.NewImageFromImage(img))
	}

	logs.Info(len(result))
	return result, nil
}

func PngDecoder(r io.Reader) (*ebiten.Image, error) {
	g, err := png.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("decode png, err: %w", err)
	}

	return ebiten.NewImageFromImage(g), nil
}

func PngRepeatDecoder(r io.Reader) ([]*ebiten.Image, error) {
	g, err := png.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("decode png, err: %w", err)
	}

	img := ebiten.NewImageFromImage(g)
	count := 60
	result := make([]*ebiten.Image, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, img)
	}

	return result, nil
}

func ByteDecoder(r io.Reader) ([]byte, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read all, err: %w", err)
	}

	return buf, nil
}

func YamlDecoder[T any](r io.Reader) (T, error) {
	var result T

	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(r); err != nil {
		return result, fmt.Errorf("read config, err: %w", err)
	}

	if err := v.Unmarshal(&result); err != nil {
		return result, fmt.Errorf("unmarshal, err: %w", err)
	}

	return result, nil
}
