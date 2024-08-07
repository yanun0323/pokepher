package app

// import (
// 	"image/color"
// 	"main/internal/utils"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/yanun0323/ebitenpkg"
// )

// type Game struct {
// 	background ebitenpkg.Image
// 	opponent   ebitenpkg.Image
// 	myPokemon  ebitenpkg.Image
// }

// func NewGame() *Game {
// 	backgroundImg := ebiten.NewImage(1366, 768)
// 	backgroundImg.Fill(color.White)
// 	background := ebitenpkg.NewImage(backgroundImg).Align(ebitenpkg.AlignTopLeading)

// 	front, err := utils.LoadImage("resource/pokemon/025/front.png")
// 	if err != nil {
// 		panic(err)
// 	}

// 	back, err := utils.LoadImage("resource/pokemon/025/back.png")
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &Game{
// 		background: background,
// 		opponent: ebitenpkg.NewImage(front).Align(ebitenpkg.AlignTopLeading).Move(800, 0).Scale(5, 5).Spriteable(ebitenpkg.SpriteSheetOption{
// 			SpriteWidthCount:  8,
// 			SpriteHeightCount: 8,
// 			SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX int, indexY int, scaleX int, scaleY int) {
// 				index := ((timestamp / 6) + 20) % 60
// 				return index % 8, index / 8, 1, 1
// 			},
// 		}).Debug(false),
// 		myPokemon: ebitenpkg.NewImage(back).Align(ebitenpkg.AlignTopLeading).Move(0, 200).Scale(7, 7).Spriteable(ebitenpkg.SpriteSheetOption{
// 			SpriteWidthCount:  8,
// 			SpriteHeightCount: 8,
// 			SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX int, indexY int, scaleX int, scaleY int) {
// 				index := (timestamp / 6) % 60
// 				return index % 8, index / 8, 1, 1
// 			},
// 		}).Debug(false),
// 	}
// }

// func (g *Game) Start() {

// }

// func (g *Game) Update() error {
// 	ebitenpkg.GameUpdate()
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	g.background.Draw(screen)
// 	g.opponent.Draw(screen)
// 	g.myPokemon.Draw(screen)
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return 1366, 768
// }
