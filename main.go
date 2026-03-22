package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var userKeyPress = false
var playerPosX = 0
var playerPosY = 0
var gameWidth = 500
var gameHeight = 500
var offsetVal = 100

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	// fmt.Printf("g keys: %v \n", len(g.keys))
	if len(g.keys) != 0 {

		userKeyPress = true
		key := fmt.Sprintf("%v", g.keys[0])
		if key == "w" || key == "W" || key == "ArrowUp" {
			if playerPosY <= 0 {
				playerPosY = 0
			} else {
				playerPosY -= offsetVal
			}
		}
		if key == "a" || key == "A" || key == "ArrowLeft" {
			if playerPosX <= 0 {
				playerPosX = 0
			} else {
				playerPosX -= offsetVal
			}
		}
		if key == "s" || key == "S" || key == "ArrowDown" {
			boundaryY := gameHeight - (offsetVal * 2)
			if playerPosY >= boundaryY {
				playerPosY = boundaryY
			} else {
				playerPosY += offsetVal
			}
		}
		if key == "d" || key == "D" || key == "ArrowRight" {
			boundaryX := gameWidth - (offsetVal * 2)
			if playerPosX >= boundaryX {
				playerPosX = boundaryX
			} else {
				playerPosX += offsetVal
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw the map
	sizeW := offsetVal - 1
	sizeH := offsetVal - 1
	offsetX := 0
	offsetY := 0

	ebitenutil.DebugPrint(screen, "Hello, World!")
	for row := 0; row < (gameHeight-offsetVal)/sizeH; row++ {
		for i := 0; i < (gameWidth-offsetVal)/sizeW; i++ {
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

	if userKeyPress == true {
		time.Sleep(150 * time.Millisecond)
		userKeyPress = false
	}
	if len(g.keys) != 0 {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", g.keys[0]))
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameWidth + offsetVal, gameHeight + offsetVal
}

func main() {
	ebiten.SetWindowSize(gameWidth*2, gameHeight*2)
	ebiten.SetWindowTitle("Lawn Mowyer")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
