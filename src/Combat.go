package gamePlay

import "fmt"

var tours int

func ChoisirAdversaire(p *Personnage) {
	var choix int
	fmt.Println("Choisissez un adversaire :")
	fmt.Println("1. Gobelin")
	fmt.Println("2. Orc")
	fmt.Println("3. Dragon")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		CombatTourParTour(p, &Gobelin)
		return // Renvoie un pointeur vers le Gobelin
	case 2:
		CombatTourParTour(p, &Orc)
		return // Renvoie un pointeur vers l'Orc
	case 3:
		CombatTourParTour(p, &Dragon)
		fmt.Println("Choix invalide.")
		return // Renvoie un pointeur nul en cas de choix invalide
	case 0:
		return
	}
}

func CombatTourParTour(p *Personnage, m *Monstre) {
	nbtours := &tours
	for p.PointsVieActuels > 0 && m.PvActuel > 0 {
		fmt.Printf("Tour %d du combat :\n", nbtours)
		fmt.Printf("%s (%s, Niveau %d) - PV: %d / %d\n", p.Nom, p.Classe, p.Niveau, p.PointsVieActuels, p.PointsVieMax)
		fmt.Printf("%s - PV: %d / %d\n", m.Nom, m.PvActuel, m.PvMaximum)
		fmt.Println("\nVous vous battez contre :", m.Nom)
		fmt.Println("\nMenu de combat:")
		fmt.Println("\n1. Attaquer")
		fmt.Println("\n2. Acceder a l'inventeur")
		fmt.Println("\n3. Quitter")
		var cicilafamille int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&cicilafamille)

		switch cicilafamille {
		case 1:
			if p.Initiative > m.Initiative {
				AttaquePersonnage(p, m)

				if m.PvActuel <= 0 {
					fmt.Printf("%s a été vaincu !\n", m.Nom)
					break
				}

				AttaqueMonstre(p, m)
				if p.PointsVieActuels <= 0 {
					fmt.Printf("%s a été vaincu !\n", p.Nom)
					fmt.Println("\nBRAVO ! Vous avez vaincu l'ennemi !")
					return
				}
			} else if m.Initiative < p.Initiative {
				AttaqueMonstre(p, m)

				if p.PointsVieActuels <= 0 {
					fmt.Printf("%s a été vaincu !\n", p.Nom)
					fmt.Println("\nBRAVO ! Vous avez vaincu l'ennemi !")
					return
				}
				AttaquePersonnage(p, m)

				if m.PvActuel <= 0 {
					fmt.Printf("%s a été vaincu !\n", m.Nom)
					break
				}
			}
		case 2:
			AccessInventory(p)

		case 0:
			return

		}
	}
}
