package main

import (
	"main/internal/app"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := app.NewGame()
	game.Start()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
