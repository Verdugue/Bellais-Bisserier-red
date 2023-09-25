package gamePlay

import (
	"fmt"
)

func Equipe(p *Personnage) {
	var choix int

	fmt.Println("Équipement actuel :")
	fmt.Printf("1- Chapeau de l'aventurier: %s\n", p.Equipement.Head)
	fmt.Printf("2- Tunique de l'aventurier : %s\n", p.Equipement.Body)
	fmt.Printf("3- Bottes de l'aventurier : %s\n", p.Equipement.Shoe)
	fmt.Print("\n")

	fmt.Println("\nMenu d'équipement:")
	fmt.Println("1- Équiper une armure")
	fmt.Println("2- Déséquiper une armure")
	fmt.Println("0- Quitter")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		fmt.Println("Équiper une armure:")
		fmt.Println("1- Chapeau de l'aventurier (+10 PV max)")
		fmt.Println("2- Tunique de l'aventurier (+25 PV max)")
		fmt.Println("3- Bottes de l'aventurier (+15 PV max)")
		fmt.Println("0- Retour")
		fmt.Print("Choisissez une armure à équiper : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			equiperArmure(p, "Chapeau de l'aventurier", 10)
		case 2:
			equiperArmure(p, "Tunique de l'aventurier", 25)
		case 3:
			equiperArmure(p, "Bottes de l'aventurier", 15)
		case 0:
			// Retour au menu précédent
		default:
			fmt.Println("Option invalide.")
		}
	case 2:
		fmt.Println("Déséquiper une armure:")
		fmt.Println("1- Chapeau de l'aventurier (-10 PV max)")
		fmt.Println("2- Tunique de l'aventurier (-25 PV max)")
		fmt.Println("3- Bottes de l'aventurier (-15 PV max)")
		fmt.Println("0- Retour")
		fmt.Print("Choisissez une armure à déséquiper : ")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			desequiperArmure(p, "Chapeau de l'aventurier", -10)
		case 2:
			desequiperArmure(p, "Tunique de l'aventurier", -25)
		case 3:
			desequiperArmure(p, "Bottes de l'aventurier", -15)
		case 0:
			return
		default:
			fmt.Println("Option invalide.")
		}
	}
}

// Fonction pour équiper une armure
func equiperArmure(personnage *Personnage, armure string, bonusPV int) {
	// Vérifier si le personnage a déjà cette armure équipée
	var equip string

	switch armure {
	case "Casque":
		equip = personnage.Equipement.Head
	case "Tunique":
		equip = personnage.Equipement.Body
	case "Bottes":
		equip = personnage.Equipement.Shoe
	}

	if equip == "" {
		// Équiper l'armure et ajouter le bonus de PV
		switch armure {
		case "Casque":
			personnage.Equipement.Head = armure
			personnage.Equipement.HeadBonus = bonusPV
		case "Tunique":
			personnage.Equipement.Body = armure
			personnage.Equipement.BodyBonus = bonusPV
		case "Bottes":
			personnage.Equipement.Shoe = armure
			personnage.Equipement.ShoeBonus = bonusPV
		}
		fmt.Printf("Vous avez équipé %s (+%d PV max).\n", armure, bonusPV)
	} else {
		fmt.Printf("Vous avez déjà équipé %s.\n", armure)
	}
}

func desequiperArmure(personnage *Personnage, armure string, bonusPV int) {
	var equip string

	switch armure {
	case "Casque":
		equip = personnage.Equipement.Head
	case "Tunique":
		equip = personnage.Equipement.Body
	case "Bottes":
		equip = personnage.Equipement.Shoe
	}

	if equip != "" {
		// Déséquiper l'armure et retirer le bonus de PV
		switch armure {
		case "Casque":
			personnage.Equipement.Head = ""
			personnage.Equipement.HeadBonus = 0
		case "Tunique":
			personnage.Equipement.Body = ""
			personnage.Equipement.BodyBonus = 0
		case "Bottes":
			personnage.Equipement.Shoe = ""
			personnage.Equipement.ShoeBonus = 0
		}
		fmt.Printf("Vous avez déséquipé %s (-%d PV max).\n", armure, bonusPV)
	} else {
		fmt.Printf("Vous n'avez pas équipé de %s.\n", armure)
	}
}
