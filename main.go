package main

import (
	"PROJET-RED/src"
	"fmt"
)

func main() {
	// Créer un personnage personnalisé avec la fonction CharCreation
	p := gamePlay.CharCreation()

	// Créer une variable pour suivre si la potion gratuite est disponible

	for {
		fmt.Println("\nMenu :")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("0. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			gamePlay.DisplayInfo(p)
		case 2:
			gamePlay.AccessInventory(p)
		case 3:
			// Afficher le menu du marchand
			gamePlay.MarchandMenu(p)
		case 4:
			gamePlay.Forgeron(p)
		case 0:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
