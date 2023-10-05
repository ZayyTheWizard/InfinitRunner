package game 

import (
    // "log"
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    // "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rectangle struct {
    X, Y, Width, Height, VelocityY float64
}

func (r *Rectangle) CollidesWithRectangle(other Rectangle) bool {
    return r.X < other.X+other.Width &&
           r.X+r.Width > other.X &&
           r.Y < other.Y+other.Height &&
           r.Y+r.Height > other.Y
}


func (r *Rectangle) updateRec() {
    // Gravity
    gravity := 0.5
    r.VelocityY += gravity

    // Update Y position
    r.Y += r.VelocityY

    // Ground collision (for simplicity, let's assume the ground is at Y = 500)
    if r.Y > 524 {
        r.Y = 524
        r.VelocityY = 0
    }
}

func (r *Rectangle) updateOb() {
    r.X -= r.VelocityY

    if r.X < -100 {
        r.X = 800
    }
}


func (r *Rectangle) DrawRec(screen *ebiten.Image, color color.Color) {
    rectImage := ebiten.NewImage(int(r.Width), int(r.Height))
    rectImage.Fill(color)

    // Set the drawing options
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(r.X, r.Y)

    // Draw the rectangle image on the screen
    screen.DrawImage(rectImage, opts)
}
