package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var playerPosX = 0
var playerPosY = 0
var gameWidth = 600
var gameHeight = 480
var offsetVal = 20

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	// fmt.Printf("g keys: %v \n", len(g.keys))
	if len(g.keys) != 0 {
		key := fmt.Sprintf("%v", g.keys[0])
		// ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", g.keys[0]))
		if key == "w" || key == "W" || key == "UpArrow" {
			playerPosY -= offsetVal
		}
		if key == "a" || key == "A" || key == "LeftArrow" {
			playerPosX -= offsetVal
		}
		if key == "s" || key == "S" || key == "DownArrow" {
			playerPosY += offsetVal
		}
		if key == "d" || key == "D" || key == "RightArrow" {
			playerPosX += offsetVal
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw the map
	sizeW := 19
	sizeH := 19
	offsetX := 0
	offsetY := 0

	ebitenutil.DebugPrint(screen, "Hello, World!")
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

	// draw player
	vector.FillRect(screen, float32(playerPosX), float32(playerPosY), float32(sizeW), float32(sizeH), color.RGBA{255, 20, 80, 200}, false)

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
