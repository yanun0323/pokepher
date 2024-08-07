package main

import (
	"main/internal/app"
	"main/resource"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

func main() {
	game := app.NewGame()
	game.Start()

	ebitenpkg.SetDefaultFont(resource.Font.Headline)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
