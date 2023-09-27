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
	potionsDeMana := 2
	Plusdinv := 3
	GemmedeReparation := 1
	for {
		fmt.Printf("\n\nPièces d'or : %d\n", p.Argent)
		fmt.Println("Menu du marchand :")
		if potionsSoin > 0 {
			fmt.Println("1. Potion de soin (3 pièces d'or)")
		} else {
			fmt.Println("1. Potion de soin  (épuisée)")
		}
		if potionsPoison > 0 {
			fmt.Println("2. Potion de poison (10 pièces d'or)")
		} else {
			fmt.Println("2. Potion de poison (épuisée)")
		}
		if bouledeFeu > 0 {
			fmt.Println("3. Livre de Sort : Boule de Feu (10 pièces d'or)")
		} else {
			fmt.Println("3. Livre de Sort : Boule de Feu (épuisé)")
		}
		if plumedecorbeau > 0 {
			fmt.Println("4. Plume de corbeau par Lot (3 pièces d'or)")
		} else {
			fmt.Println("4. Plume de corbeau (épuisée)")
		}
		if cuirdesanglier > 0 {
			fmt.Println("5. Cuir de Sanglier par Lot (9 pièces d'or)")
		} else {
			fmt.Println("5. Cuir de Sanglier (épuisé)")
		}
		if fourruredeloup > 0 {
			fmt.Println("6. Fourrure de loup (21 pièces d'or)")
		} else {
			fmt.Println("6. Fourrure de loup (épuisée)")
		}
		if peaudetroll > 0 {
			fmt.Println("7. Peau de Troll par Lot (12 pièces d'or)")
		} else {
			fmt.Println("7. Peau de Troll (épuisée)")
		}
		if potionsDeMana > 0 {
			fmt.Println("8. Potion de Mana (5 pièces d'or)")
		} else {
			fmt.Println("8. Potion de Mana (épuisée)")
		}
		if Plusdinv > 0 {
			fmt.Println("9. Plus d'inventaire x3 (10 pièces)")
		} else {
			fmt.Println("9. Plus d'inventaire (épuisée)")
		}
		if GemmedeReparation > 0 {
			fmt.Println("10 . Gemme De Reparation ( 25 pièces d'or )") // Ajoutez l'option d'achat du livre de sort
		} else {
			fmt.Println("10 . Gemme De Reparation ( épuisé )")
		}
		fmt.Println("0. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			if potionsSoin > 0 && len(p.Inventaire) < 10 && p.Argent > 3 {
				// Ajouter une potion de soin à l'inventaire du personnage
				p.Inventaire["Potion de soin"]++
				fmt.Println("Vous avez acheté une Potion de soin.")
				// Réduire le nombre de potions de soin disponibles chez le marchand
				potionsSoin--
				p.Argent -= 3
			} else if potionsSoin > 0 {
				fmt.Println("Le marchand n'a plus de potions de soin en stock.")
			}

		case 2:
			if potionsPoison > 0 && len(p.Inventaire) < 10 && p.Argent >= 6 {
				p.Inventaire["Potion de poison"]++
				fmt.Println("Vous avez acheté une Potion de Poison.")
				potionsPoison--
				p.Argent -= 6
			} else if potionsPoison > 0 {
				fmt.Println("Le marchand n'a plus de potions de Poison en stock.")
			}
		case 3:
			// Vérifier si le "Livre de Sort : Boule de Feu" est disponible
			if bouledeFeu > 0 && len(p.Inventaire) < 10 && p.Argent >= 25 {
				p.Inventaire["Boule de Feu"]++
				fmt.Println("Vous avez appris le sort : Boule de Feu.")
				bouledeFeu--
				p.Argent -= 10
			} else {
				fmt.Println("Vous avez déjà appris le sort : Boule de feu")
			}

		case 4:
			if plumedecorbeau > 0 && len(p.Inventaire) < 10 && p.Argent >= 4 {
				p.Inventaire["Plume de Corbeau"]++
				fmt.Println("Vous venez d'acheter une Plume de Corbeau")
				plumedecorbeau--
				p.Argent -= 4
			} else if plumedecorbeau == 0 {
				fmt.Println("Le marchand n'a plus de plume de corbeau en stock.")
			}

		case 5:
			if cuirdesanglier > 0 && len(p.Inventaire) < 10 && p.Argent >= 3 {
				p.Inventaire["Cuir de sanglier"]++
				fmt.Println("Vous venez d'acheter un cuir de sanglier")
				cuirdesanglier--
				p.Argent -= 3
			} else if cuirdesanglier > 0 {
				fmt.Println("Le marchand n'a plus de cuir de sanglier en stock.")
			}

		case 6:
			if fourruredeloup > 0 && len(p.Inventaire) < 10 && p.Argent >= 4 {
				p.Inventaire["Fourrure de loup"]++
				fourruredeloup--
				p.Argent -= 4
			} else if fourruredeloup > 0 {
				fmt.Println("Le marchand n'a plus de forrure de loup en stock.")
			}

		case 7:
			if peaudetroll > 0 && len(p.Inventaire) < 10 && p.Argent >= 7 {
				p.Inventaire["Peau de troll"]++
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
		case 9:
			if p.Argent >= 10 && potionsDeMana > 0 {
				p.Inventaire["Potion de mana"]++
				fmt.Println("Vous venez d'acheter une Potion de mana")
				potionsDeMana--
				p.Argent -= 9
			}
		case 10:
			if p.Argent >= 25 && GemmedeReparation > 0 {
				p.Inventaire["Gemme de Reparation"]++
				fmt.Println("Vous venez d'acheter une Gemme de Reparation")
				GemmedeReparation--
				p.Argent -= 25
			}
		case 0:
			fmt.Println("Vous quittez le menu du marchand.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
