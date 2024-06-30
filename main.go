package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err        error
	background *ebiten.Image
	miner      *ebiten.Image
	playerOne  player
)

type Game struct{}

type player struct {
	image *ebiten.Image
	xPos  float64
	yPos  float64
	speed float64
}

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

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	miner, _, err = ebitenutil.NewImageFromFile("assets/Miner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{miner, 640 / 2, 480 / 2, 4}
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Mining Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
