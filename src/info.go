package gamePlay

import (
	"fmt"

)

func InfoDebut(p *Personnage) {
	fmt.Println("-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-")
	fmt.Println("Voici votre personnage :")
	fmt.Println("Nom:", p.Nom)
	fmt.Println("Classe:", p.Classe)
	fmt.Println("Niveau:", p.Niveau)
	fmt.Println("-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-")

}

func DisplayInfo(p *Personnage) {
	fmt.Println(" __  .__   __.  _______   ______   .______      .___  ___.      ___      .___________. __    ______   .__   __.")
	fmt.Println("|  | |  \\ |  | |   ____| /  __  \\  |   _  \\     |   \\/   |     /   \\     |           ||  |  /  __  \\  |  \\ |  |")
	fmt.Println("|  | |   \\|  | |  |__   |  |  |  | |  |_)  |    |  \\  /  |    /  ^  \\    `---|  |----`|  | |  |  |  | |   \\|  |")
	fmt.Println("|  | |  . `  | |   __|  |  |  |  | |      /     |  |\\/|  |   /  /_\\  \\       |  |     |  | |  |  |  | |  . `  |")
	fmt.Println("|  | |  |\\   | |  |     |  `--'  | |  |\\  \\----.|  |  |  |  /  _____  \\      |  |     |  | |  `--'  | |  |\\   |")
	fmt.Println("|__| |__| \\__| |__|      \\______/  | _| `._____||__|  |__| /__/     \\__\\     |__|     |__|  \\______/  |__| \\__|")

	fmt.Println("\nInformations du personnage :")
	fmt.Println("Nom:", p.Nom)
	fmt.Println("Classe:", p.Classe)
	fmt.Println("Niveau:", p.Niveau)
	fmt.Println("Points de vie maximum:", p.PointsVieMax)
	fmt.Println("Points de vie actuels:", p.PointsVieActuels)
	fmt.Println("Inventaire: Vous avez maintenant\n", limit, "place dans votre inventaire.")
	fmt.Println("Equipement:", p.Equipement.Head, p.Equipement.Body, p.Equipement.Shoe)
	fmt.Printf("\nVous avez : %d", p.Xp)
}

func MonstreInfo(m *Monstre) {
	fmt.Println("Vous Combattez contre ", m.Nom)
	fmt.Println("Nom du monstre : ", m.Nom)
	fmt.Println("Points de vie maximum : ", m.PvMaximum)
	fmt.Println("Points de vie actuel : ", m.PvActuel)
	fmt.Println("Point d'attaque : ", m.PointAttact)
}
