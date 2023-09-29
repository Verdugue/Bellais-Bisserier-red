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
		fmt.Println(".___  ___.      ___      .______        ______  __    __       ___      .__   __.  _______")
		fmt.Println("|   \\/   |     /   \\     |   _  \\      /      ||  |  |  |     /   \\     |  \\ |  | |       \\")
		fmt.Println("|  \\  /  |    /  ^  \\    |  |_)  |    |  ,----'|  |__|  |    /  ^  \\    |   \\|  | |  .--.  |")
		fmt.Println("|  |\\/|  |   /  /_\\  \\   |      /     |  |     |   __   |   /  /_\\  \\   |  . `  | |  |  |  |")
		fmt.Println("|  |  |  |  /  _____  \\  |  |\\  \\----.|  `----.|  |  |  |  /  _____  \\  |  |\\   | |  '--'  |")
		fmt.Println("|__|  |__| /__/     \\__\\ | _| `._____| \\______||__|  |__| /__/     \\__\\ |__| \\__| |_______/")

		fmt.Printf("\n\nPièces d'or : %d\n", p.Argent)
		fmt.Println("Menu du marchand :")
		fmt.Println("1. Voulez-vous acheter des items.")
		fmt.Println("2. Voulez-vous vendre des items.")
		fmt.Println("3. Quitter.")
		fmt.Println("0. Quitter")
		var voa int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&voa)
		switch voa {
		case 1:
			fmt.Println("Veuillez choisir une option entre 1 et 10.")
			if potionsSoin > 0 {
				fmt.Println("1. Potion de soin (3 pièces d'or)")
			} else {
				fmt.Println("1. Potion de soin (épuisée)")
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
				fmt.Println("10. Gemme De Réparation (25 pièces d'or)") // Ajoutez l'option d'achat du livre de sort
			} else {
				fmt.Println("10. Gemme De Réparation (épuisée)")
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
					fmt.Println("Vous venez d'acheter une Fourrure de loup.")
					fourruredeloup--
					p.Argent -= 4
				} else if fourruredeloup > 0 {
					fmt.Println("Le marchand n'a plus de Fourrure de loup en stock.")
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
					fmt.Println("\nVous venez d'acheter un sac à dos")
					fmt.Println("\nVous avez obtenu 10 places supplémentaires")
					p.Argent -= 30
				} else {
					fmt.Println("Vous n'avez pas assez d'argent pour acheter un sac à dos.")
				}
			case 9:
				if p.Argent >= 10 && potionsDeMana > 0 {
					p.Inventaire["Potion de mana"]++
					fmt.Println("Vous venez d'acheter une Potion de mana")
					potionsDeMana--
					p.Argent -= 9
				} else if potionsDeMana == 0 {
					fmt.Println("Le marchand n'a plus de potions de mana en stock.")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent pour acheter une Potion de mana.")
				}
			case 10:
				if p.Argent >= 25 && GemmedeReparation > 0 {
					p.Inventaire["Gemme de Reparation"]++
					fmt.Println("Vous venez d'acheter une Gemme de Réparation")
					GemmedeReparation--
					p.Argent -= 25
				} else if GemmedeReparation == 0 {
					fmt.Println("Le marchand n'a plus de Gemme de Réparation en stock.")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent pour acheter une Gemme de Réparation.")
				}
			case 0:
				fmt.Println("Vous quittez le menu du marchand.")
				return
			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}

		case 2:
			var vend int
			fmt.Println("\nQue voulez-vous vendre.")
			fmt.Println("1. Potion de soin (2 pièces d'or)")
			fmt.Println("2. Potion de poison (6 pièces d'or)")
			fmt.Println("3. Livre de Sort : Boule de Feu (10 pièces d'or)")
			fmt.Println("4. Plume de corbeau (12 pièces d'or)")
			fmt.Println("5. Cuir de Sanglier (18 pièces d'or)")
			fmt.Println("6. Fourrure de loup (30 pièces d'or)")
			fmt.Println("7. Peau de Troll (42 pièces d'or)")
			fmt.Println("8. Potion de Mana (60 pièces d'or)")
			fmt.Println("9. Plus d'inventaire (100 pièces)")
			fmt.Println("10. Gemme de Réparation (250 pièces d'or)")
			fmt.Println("0. Quitter")
			fmt.Print("Votre choix : ")
			fmt.Scanln(&vend)
			switch vend {
			case 1:
				if p.Inventaire["Potion de soin"] > 0 {
					// Vendre une Potion de soin
					p.Inventaire["Potion de soin"]--
					fmt.Println("Vous avez vendu une Potion de soin.")
					// Ajouter de l'argent au personnage
					p.Argent += 2
					potionsSoin++
				} else {
					fmt.Println("Vous n'avez pas de Potion de soin à vendre.")
				}
			case 2:
				if p.Inventaire["Potion de poison"] > 0 {
					// Vendre une Potion de poison
					p.Inventaire["Potion de poison"]--
					fmt.Println("Vous avez vendu une Potion de poison.")
					// Ajouter de l'argent au personnage
					p.Argent += 6
					potionsPoison++
				} else {
					fmt.Println("Vous n'avez pas de Potion de poison à vendre.")
				}
			case 3:
				if p.Inventaire["Boule de Feu"] > 0 {
					// Vendre un Livre de Sort : Boule de Feu
					p.Inventaire["Boule de Feu"]--
					fmt.Println("Vous avez vendu un Livre de Sort : Boule de Feu.")
					// Ajouter de l'argent au personnage
					p.Argent += 10
					bouledeFeu++
				} else {
					fmt.Println("Vous n'avez pas de Livre de Sort : Boule de Feu à vendre.")
				}
			case 4:
				if p.Inventaire["Plume de Corbeau"] > 0 {
					// Vendre une Plume de Corbeau
					p.Inventaire["Plume de Corbeau"]--
					fmt.Println("Vous avez vendu une Plume de Corbeau.")
					// Ajouter de l'argent au personnage
					p.Argent += 12
					plumedecorbeau++
				} else {
					fmt.Println("Vous n'avez pas de Plume de Corbeau à vendre.")
				}
			case 5:
				if p.Inventaire["Cuir de sanglier"] > 0 {
					// Vendre un Cuir de Sanglier
					p.Inventaire["Cuir de sanglier"]--
					fmt.Println("Vous avez vendu un Cuir de Sanglier.")
					// Ajouter de l'argent au personnage
					p.Argent += 18
					cuirdesanglier++
				} else {
					fmt.Println("Vous n'avez pas de Cuir de Sanglier à vendre.")
				}
			case 6:
				if p.Inventaire["Fourrure de loup"] > 0 {
					// Vendre une Fourrure de loup
					p.Inventaire["Fourrure de loup"]--
					fmt.Println("Vous avez vendu une Fourrure de loup.")
					// Ajouter de l'argent au personnage
					p.Argent += 30
					fourruredeloup++
				} else {
					fmt.Println("Vous n'avez pas de Fourrure de loup à vendre.")
				}
			case 7:
				if p.Inventaire["Peau de troll"] > 0 {
					// Vendre une Peau de troll
					p.Inventaire["Peau de troll"]--
					fmt.Println("Vous avez vendu une Peau de troll.")
					// Ajouter de l'argent au personnage
					p.Argent += 42
					peaudetroll++
				} else {
					fmt.Println("Vous n'avez pas de Peau de troll à vendre.")
				}
			case 8:
				if p.Inventaire["Potion de mana"] > 0 {
					// Vendre une Potion de mana
					p.Inventaire["Potion de mana"]--
					fmt.Println("Vous avez vendu une Potion de mana.")
					// Ajouter de l'argent au personnage
					p.Argent += 60
					potionsDeMana++
				} else {
					fmt.Println("Vous n'avez pas de Potion de mana à vendre.")
				}
			case 9:
				if p.Argent >= 100 {
					UpgradeInventorySlot(p)
					fmt.Println("\nVous venez d'acheter un sac à dos.")
					fmt.Println("\nVous avez obtenu 10 places supplémentaires.")
					p.Argent -= 100
					Plusdinv++
				} else {
					fmt.Println("Vous n'avez pas assez d'argent pour acheter un sac à dos.")
				}
			case 10:
				if p.Inventaire["Gemme de Reparation"] > 0 {
					// Vendre une Gemme de Réparation
					p.Inventaire["Gemme de Reparation"]--
					fmt.Println("Vous avez vendu une Gemme de Réparation.")
					// Ajouter de l'argent au personnage
					p.Argent += 250
					GemmedeReparation++
				} else {
					fmt.Println("Vous n'avez pas de Gemme de Réparation à vendre.")
				}
			case 0:
				return
			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		case 3:
			fmt.Println("Vous quittez le menu du marchand.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
