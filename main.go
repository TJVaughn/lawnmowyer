package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	gameWidth  = 500
	gameHeight = 500
	offsetVal  = 100
	sizeW      = offsetVal - 1
	sizeH      = offsetVal - 1
)

var introState = true
var userKeyPress = false
var playerPosX = 0
var playerPosY = 0
var gameState [4][4]int
var levelFailed = false

type Game struct {
	keys []ebiten.Key
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	// fmt.Printf("g keys: %v \n", len(g.keys))
	if len(g.keys) != 0 {
		userKeyPress = true
		key := fmt.Sprintf("%v", g.keys[0])
		if introState && key == "Enter" {
			introState = false
		}
		if key == "Escape" {
			time.Sleep(300 * time.Millisecond)
			if introState == true {
				os.Exit(1)
			}
			introState = true
			createGameState()
			playerPosX = 0
			playerPosY = 0
		}
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
		playerX := playerPosX / 100
		playerY := playerPosY / 100
		if gameState[playerY][playerX] != 1 {
			gameState[playerY][playerX] = 1
		} else {
			levelFailed = true
		}
	}
	return nil
}

var (
	mplusFaceSource *text.GoTextFaceSource
)

func introScreen(screen *ebiten.Image) {

	const (
		normalFontSize = 24
		bigFontSize    = 48
	)
	const x = 20
	title := "Lawn Mowyer"
	subTit := "Press Enter to Start"
	op := &text.DrawOptions{}
	// Draw the sample text
	op.GeoM.Translate(10, 60)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, title, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, op)

	op.GeoM.Translate(0, 80)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, subTit, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if introState == true {
		introScreen(screen)
		return
	}
	// draw the map
	offsetX := 0
	offsetY := 0

	for _, row := range gameState {
		for _, col := range row {
			// fmt.Println("%v", col)
			cellColor := color.RGBA{0, 255, 0, 255}
			if col == 1 {
				cellColor = color.RGBA{0, 120, 0, 255}
			}
			vector.FillRect(screen, float32(offsetX), float32(offsetY), float32(sizeW), float32(sizeH), cellColor, false)

			offsetX += offsetVal
		}
		offsetX = 0
		offsetY += offsetVal
	}

	// draw player
	vector.FillRect(screen, float32(playerPosX), float32(playerPosY), float32(sizeW), float32(sizeH), color.RGBA{255, 20, 80, 200}, false)

	if userKeyPress == true {
		time.Sleep(150 * time.Millisecond)
		userKeyPress = false
	}
	// if len(g.keys) != 0 {
	// 	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", g.keys[0]))
	// }
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("X %v, Y %v", playerPosX, playerPosY))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Level failed: %v", levelFailed))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameWidth + offsetVal, gameHeight + offsetVal
}

func createGameState() {
	for row := 0; row < (gameHeight-offsetVal)/sizeH; row++ {
		for col := 0; col < (gameWidth-offsetVal)/sizeW; col++ {
			if row == 0 && col == 0 {
				gameState[row][col] = 1
			} else {
				gameState[row][col] = 0
			}
		}
	}
	// fmt.Printf("row len %v, col len %v ", len(gameState), len(gameState[0]))

}

func main() {
	ebiten.SetWindowSize(gameWidth*2, gameHeight*2)
	ebiten.SetWindowTitle("Lawn Mowyer")

	createGameState()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
