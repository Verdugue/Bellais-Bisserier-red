package game

import (
	"fmt"
)


type Character struct{
	name		string
	class		string
	level		int
	pv_max		int
	pv			int
	inventory	[]struct {
		name string
		amount int
	}
}

func Init(name string, class string) *Character {
    character := &Character{
        name:   name,
        class:  class,
        level:  1,
        pv_max: 100,
        pv:     40,
        inventory: []struct {
            name   string
            amount int
        }{
            {"Potions de soin", 3},
        },
    }
    return character
}

func DisplayInfo(character *Character) {
    fmt.Printf("Nom: %s\n", character.name)
    fmt.Printf("Classe: %s\n", character.class)
    fmt.Printf("Niveau: %d\n", character.level)
    fmt.Printf("PV maximum: %d\n", character.pv_max)
    fmt.Printf("PV actuels: %d\n", character.pv)
    fmt.Println("Inventaire:")
    for _, item := range character.inventory {
        fmt.Printf("Nom: %s, Quantité: %d\n", item.name, item.amount)
    }
}

func AccessInventory(character *Character) {
    for {
        fmt.Println("Inventaire du personnage:")
        for i, item := range character.inventory {
            fmt.Printf("%d. Nom: %s, Quantité: %d\n", i+1, item.name, item.amount)
        }
        fmt.Println("\nMenu de l'inventaire:")
        fmt.Println("1. Utiliser un objet")
        fmt.Println("2. Retour au menu principal")

        var choice int
        fmt.Print("Choisissez une option: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            useItem(character)
        case 2:
            return // Retour au menu principal
        default:
            fmt.Println("Option invalide")
        }
    }
}

func useItem(character *Character) {
    fmt.Println("Choisissez un objet à utiliser (1, 2, etc.):")
    var itemIndex int
    fmt.Scanln(&itemIndex)

    if itemIndex >= 1 && itemIndex <= len(character.inventory) {
        item := &character.inventory[itemIndex-1]
        fmt.Printf("Vous utilisez un(e) %s\n", item.name)
        // Ajoutez ici la logique pour l'utilisation de l'objet
    } else {
        fmt.Println("Option invalide")
    }
}
