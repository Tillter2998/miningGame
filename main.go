package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err        error
	background *ebiten.Image
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/background.png", ebiten.FilterDefault)
	if err != nil {
		return
	}
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Mining Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
