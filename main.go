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

var (
	levelStart     = true
	introState     = true
	userKeyPress   = false
	playerPosX     = 0
	playerPosY     = 0
	gameState      [4][4]int
	levelFailed    = false
	isLevelSuccess = false
)

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
func resetGame() {
	createGameState()
	levelFailed = false
	levelStart = true
	time.Sleep(300 * time.Millisecond)
	// introState = true
	playerPosX = 0
	playerPosY = 0
}

func checkIsLevelSuccess() {
	for row := 0; row < len(gameState); row++ {
		for col := 0; col < len(gameState[row]); col++ {
			if gameState[row][col] == 0 {
				fmt.Printf("something still equal to zero %v %v \n", row, col)
				fmt.Printf("player x %v y %v \n", playerPosX, playerPosY)
				isLevelSuccess = false
				return
			}
		}
	}
	isLevelSuccess = true
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	// fmt.Printf("g keys: %v \n", len(g.keys))
	if len(g.keys) != 0 {
		userKeyPress = true
		key := fmt.Sprintf("%v", g.keys[0])
		levelStart = false
		if introState && key == "Enter" {
			introState = false
			levelStart = true
			resetGame()
			// return nil
		}
		if key == "Escape" {
			if introState == true {
				os.Exit(1)
			}
			introState = true
			resetGame()
			// return nil
		}
		if key == "w" || key == "W" || key == "ArrowUp" {
			if playerPosY <= 0 {
				// playerPosY = 0
				return nil
			} else {
				playerPosY -= offsetVal
			}
		}
		if key == "a" || key == "A" || key == "ArrowLeft" {
			if playerPosX <= 0 {
				return nil
				// playerPosX = 0
			} else {
				playerPosX -= offsetVal
			}
		}
		if key == "s" || key == "S" || key == "ArrowDown" {
			boundaryY := gameHeight - (offsetVal * 2)
			if playerPosY >= boundaryY {
				return nil
				// playerPosY = boundaryY
			} else {
				playerPosY += offsetVal
			}
		}
		if key == "d" || key == "D" || key == "ArrowRight" {
			boundaryX := gameWidth - (offsetVal * 2)
			if playerPosX >= boundaryX {
				return nil
				// playerPosX = boundaryX
			} else {
				playerPosX += offsetVal
			}
		}
		playerX := playerPosX / 100
		playerY := playerPosY / 100
		if gameState[playerY][playerX] != 1 {
			gameState[playerY][playerX] = 1
		} else {
			if playerX == 0 && playerY == 0 && levelStart == true {
				levelFailed = false
			} else {
				fmt.Printf("level failed true. playerx %v playerY %v level start %v\n", playerX, playerY, levelStart)
				levelFailed = true
			}
		}
		checkIsLevelSuccess()
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

	op.GeoM.Translate(0, 30)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, subTit, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, op)
	if levelFailed || isLevelSuccess {
		var msg string
		var msgColor = color.RGBA{125, 0, 0, 255}
		if levelFailed {
			msg = "level failed"
		} else {
			msgColor = color.RGBA{0, 200, 0, 255}
			msg = "level passed"
		}

		op.GeoM.Translate(0, 50)
		op.ColorScale.ScaleWithColor(msgColor)
		text.Draw(screen, msg, &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   normalFontSize,
		}, op)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if levelFailed == true {
		introState = true
	}
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
	// for row := 0; row < (gameHeight-offsetVal)/sizeH; row++ {
	// 	for col := 0; col < (gameWidth-offsetVal)/sizeW; col++ {
	// 		gameState[row][col] = 0
	// 	}
	// }
	for row := 0; row < len(gameState); row++ {
		for col := 0; col < len(gameState[row]); col++ {
			gameState[row][col] = 0
		}
	}

}

func main() {
	ebiten.SetWindowSize(gameWidth*2, gameHeight*2)
	ebiten.SetWindowTitle("Lawn Mowyer")

	resetGame()
	// createGameState()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
