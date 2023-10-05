package game

import (
    "fmt"
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    "image/color"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    screenWidth int = 800
    screenHeight int = 640
)

type Game struct {
    Player Rectangle
    Score int
}

// Load Graphics
var (
    platformImage *ebiten.Image
    ob1 = Rectangle {
        X: 800,
        Y: 495,
        Width: 50,
        Height: 80,
        VelocityY: 5,
    }
)

func init() {
    var err error
    platformImage, _, err = ebitenutil.NewImageFromFile("./images/PlatForm.png")
    if err != nil {
        log.Fatalf("Failed to load the platform image: %v", err)
    }
}


func (g *Game) Update() error {
    timer := 0
    if ebiten.IsKeyPressed(ebiten.KeySpace) {
        // Apply a negative velocity to simulate a jump (adjust the value as needed)
        g.Player.VelocityY = -10
    }

    // Update the rectangle's position
    ob1.updateOb()
    g.Player.updateRec()
    if timer % 10000000 == 0 {
        g.Score++
        timer = 0
    }
    if g.Player.CollidesWithRectangle(ob1) {
        ob1.VelocityY = 5.0
        g.Score = 0
    }

    if g.Score % 100 == 0 {
        ob1.VelocityY++
    }
    timer++
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    scoreText := fmt.Sprintf("Score: %d", g.Score)
    ebitenutil.DebugPrint(screen, scoreText)

    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(0, 640-100) // Position the platform at (100, screenHeight-50)
    screen.DrawImage(platformImage, opts)

    redColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}

    g.Player.DrawRec(screen, redColor) 
    ob1.DrawRec(screen, redColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 800, 640
}
