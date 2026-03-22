package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type Game struct{}

var gameWidth = 600
var gameHeight = 480

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	sizeW := 19
	sizeH := 19
	offsetX := 0
	offsetY := 0
	offsetVal := 20

	for row := 0; row < gameHeight/sizeH; row++ {
		for i := 0; i < gameWidth/sizeW; i++ {
			vector.FillRect(screen, float32(offsetX), float32(offsetY), float32(sizeW), float32(sizeH), color.RGBA{0, 255, 0, 255}, false)

			offsetX += offsetVal
			// offsetY += 10
		}
		offsetX = 0
		offsetY += offsetVal
	}
	// vector.FillRect(screen, 50, 50, 100, 100, color.RGBA{0, 255, 0, 255}, false)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Lawn Mowyer")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
