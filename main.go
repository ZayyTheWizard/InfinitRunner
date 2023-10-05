package main

import(
    "log"
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "isaiah.com/m/pkg/game"
)

func main() {
    g := &game.Game{
        Player: game.Rectangle{
            X:        50,
            Y:        524,
            Width:    50,
            Height:   50,
            VelocityY: 0,
        },
        Score: 0,
    }
    
    ebiten.SetWindowSize(900, 720)
    ebiten.SetWindowTitle("Hello Window")
    if err := ebiten.RunGame(g); err!= nil {
        log.Fatal(err)
    }

    fmt.Println("Done")
}
