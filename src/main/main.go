package main

import (
    "game"
    "os"
    "os/exec"
	"fmt"
)

func clearScreen() {
    cmd := exec.Command("clear") // Efface l'écran sous Unix/Linux/macOS
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
    p1 := game.Init("Alice", "Guerrier")
    for {
        fmt.Println("Menu principal:")
        fmt.Println("1. Afficher les informations du personnage")
        fmt.Println("2. Accéder à l'inventaire")
        fmt.Println("3. Quitter")

        var choice int
        fmt.Print("Choisissez une option: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            game.DisplayInfo(p1)
        case 2:
            game.AccessInventory(p1)
        case 3:
            fmt.Println("Au revoir !")
            return
        default:
            fmt.Println("Option invalide")
        }
        clearScreen()
    }
}
