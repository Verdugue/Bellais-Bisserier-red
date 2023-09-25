package gamePlay

import "fmt"

func MarchandMenu(p *Personnage) {
	potionsSoin := 1   // Nombre initial de potions de soin disponibles chez le marchand
	potionsPoison := 1 // Le "Livre de Sort : Boule de Feu" est disponible
	bouledeFeu := 1
	plumedecorbeau := 1
	cuirdesanglier := 2
	fourruredeloup := 3
	peaudetroll := 1

	for {
		fmt.Printf("\n\nPièces d'or : %d\n", p.Argent)
		fmt.Println("Menu du marchand :")
		fmt.Println("1. Potion de soin (Gratuite)")
		fmt.Println("2. Potion de poison (10 pièces d'or)")
		fmt.Println("3. Boule de feu (10 pièces d'or")
		fmt.Println("4. Plume de corbeau par Lot (3 pièces d'or)")
		fmt.Println("5. Cuir de Sanglier par Lot (9 pièces d'or)")
		fmt.Println("6. Fourrure de loup (21 pièces d'or)")
		fmt.Println("7. Peau de Troll par Lot (12 pièces d'or)")
		fmt.Println("8. Sac a dos x3 (30 pièces d'or)")
		if bouledeFeu > 0 {
			fmt.Println("3. Livre de Sort : Boule de Feu (10 pièces d'or)") // Ajoutez l'option d'achat du livre de sort
		} else {
			fmt.Println("3. Livre de Sort : Boule de Feu (épuisé)")
		}
		fmt.Println("0. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			if potionsSoin > 0 && len(p.Inventaire) < 10 && p.Argent > 3 {
				// Ajouter une potion de soin à l'inventaire du personnage
				p.AddToInventory("Potion de soin")
				fmt.Println("Vous avez acheté une Potion de soin.")
				// Réduire le nombre de potions de soin disponibles chez le marchand
				potionsSoin--
				p.Argent -= 3
			} else if potionsSoin > 0 {
				fmt.Println("Le marchand n'a plus de potions de soin en stock.")
			}

		case 2:
			if potionsPoison > 0 && len(p.Inventaire) < 10 && p.Argent >= 6 {
				p.AddToInventory("Potion de poison")
				fmt.Println("Vous avez acheté une Potion de Poison.")
				potionsPoison--
				p.Argent -= 6
			} else if potionsPoison > 0 {
				fmt.Println("Le marchand n'a plus de potions de Poison en stock.")
			}
		case 3:
			// Vérifier si le "Livre de Sort : Boule de Feu" est disponible
			if bouledeFeu > 0 && len(p.Inventaire) < 10 && p.Argent >= 25 {
				p.AddToInventory("Boule de Feu")
				fmt.Println("Vous avez appris le sort : Boule de Feu.")
				bouledeFeu--
				p.Argent -= 10
			} else {
				fmt.Println("Vous avez déjà appris le sort : Boule de feu")
			}

		case 4:
			if plumedecorbeau > 0 && len(p.Inventaire) < 10 && p.Argent >= 4 {
				p.AddToInventory("Plume de Corbeau")
				fmt.Println("Vous venez d'acheter une Plume de Corbeau")
				plumedecorbeau--
				p.Argent -= 4
			} else if plumedecorbeau == 0 {
				fmt.Println("Le marchand n'a plus de plume de corbeau en stock.")
			}

		case 5:
			if cuirdesanglier > 0 && len(p.Inventaire) < 10 && p.Argent >= 3 {
				p.AddToInventory("Cuir de Sanglier")
				fmt.Println("Vous venez d'acheter un cuir de sanglier")
				cuirdesanglier--
				p.Argent -= 3
			} else if cuirdesanglier > 0 {
				fmt.Println("Le marchand n'a plus de cuir de sanglier en stock.")
			}

		case 6:
			if fourruredeloup > 0 && len(p.Inventaire) < 10 && p.Argent >= 4 {
				p.AddToInventory("Fourrure de loup")
				fourruredeloup--
				p.Argent -= 4
			} else if fourruredeloup > 0 {
				fmt.Println("Le marchand n'a plus de forrure de loup en stock.")
			}

		case 7:
			if peaudetroll > 0 && len(p.Inventaire) < 10 && p.Argent >= 7 {
				p.AddToInventory("Peau de Troll")
				fmt.Println("Vous venez d'acheter une Peau de troll")
				peaudetroll--
				p.Argent -= 7
			} else if peaudetroll > 0 {
				fmt.Println("Le marchand n'a plus de peau de troll en stock.")
			}
		case 8:
			if p.Argent >= 30 {
				UpgradeInventorySlot(p)
				fmt.Println("\nVous venez d'acheter un sac a dos")
				fmt.Println("\nVous avez obtenue 10 places suplementaires")
				p.Argent -= 30
			}
		case 0:
			fmt.Println("Vous quittez le menu du marchand.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
