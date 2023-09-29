package gamePlay

import (
	"fmt"
	"math/rand"
)

func ChoisirAdversaire(p *Personnage) {
	fmt.Println("  ______   ______   .___  ___. .______        ___      .___________.")
	fmt.Println(" /      | /  __  \\  |   \\/   | |   _  \\      /   \\     |           |")
	fmt.Println("|  ,----'|  |  |  | |  \\  /  | |  |_)  |    /  ^  \\    `---|  |----`")
	fmt.Println("|  |     |  |  |  | |  |\\/|  | |   _  <    /  /_\\  \\       |  |")
	fmt.Println("|  `----.|  `--'  | |  |  |  | |  |_)  |  /  _____  \\      |  |")
	fmt.Println(" \\______| \\______/  |__|  |__| |______/  /__/     \\__\\     |__|")
	var choix int
	fmt.Println("\nChoisissez un adversaire :")
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
	nbtours := 0
	for p.PointsVieActuels > 0 && m.PvActuel > 0 {
		fmt.Printf("Tour %d du combat :\n", nbtours+1)
		fmt.Printf("%s (%s, Niveau %d) - PV: %d / %d\n", p.Nom, p.Classe, p.Niveau, p.PointsVieActuels, p.PointsVieMax)
		fmt.Printf("\nVous avez encore %d mana", p.Mana)
		fmt.Printf("\n%s - PV: %d / %d\n", m.Nom, m.PvActuel, m.PvMaximum)
		fmt.Println("\nVous vous battez contre :", m.Nom)
		fmt.Println("\nMenu de combat:")
		fmt.Println("\n1. Attaquer")
		fmt.Println("\n2. Acceder a l'inventaire")
		fmt.Println("\n0. Quitter")
		var combat int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&combat)

		switch combat {
		case 1:
			AttaquePersonnage(p, m)

			if m.PvActuel <= 0 {
				fmt.Print("__ __ __      # ________     # ___   __      #\n")
				fmt.Print("/_//_//_/\\    #/_______/\\   #/__/\\ /__/\\    #\n")
				fmt.Print("\\:\\:\\:\\ \\    #\\__.::._\\/    #\\::\\_\\  \\ \\   #\n")
				fmt.Print(" \\:\\:\\:\\ \\   #   \\::\\ \\     # \\:. `-\\  \\ \\  #\n")
				fmt.Print("  \\:\\:\\:\\ \\  #   _\\::\\ \\__  #  \\:. _    \\ \\ #\n")
				fmt.Print("   \\:\\:\\:\\ \\ #  /__\\::\\__/\\ #   \\. \\`-\\  \\ \\#\n")
				fmt.Print("    \\_______\\/ # \\________\\/ #   \\__\\/ \\__\\/#\n")
				fmt.Printf("\nBRAVO ! Vous avez vaincu %s !", m.Nom)
				break
			}

			AttaqueMonstre(p, m)
			CheckDura(p)
			if p.PointsVieActuels <= 0 {
				return
			}

			if m.PvActuel <= 0 {
				ClearConsole()
				fmt.Print("__ __ __      # ________     # ___   __      #\n")
				fmt.Print("/_//_//_/\\    #/_______/\\   #/__/\\ /__/\\    #\n")
				fmt.Print("\\:\\:\\:\\ \\    #\\__.::._\\/    #\\::\\_\\  \\ \\   #\n")
				fmt.Print(" \\:\\:\\:\\ \\   #   \\::\\ \\     # \\:. `-\\  \\ \\  #\n")
				fmt.Print("  \\:\\:\\:\\ \\  #   _\\::\\ \\__  #  \\:. _    \\ \\ #\n")
				fmt.Print("   \\:\\:\\:\\ \\ #  /__\\::\\__/\\ #   \\. \\`-\\  \\ \\#\n")
				fmt.Print("    \\_______\\/ #  \\________\\/ #   \\__\\/ \\__\\/#\n")
				fmt.Printf("%s a été vaincu !\n", m.Nom)
				p.Xp = p.Xp + 5
				LevelUp(p)
				Jardin(p, m)
				return
			}
		case 2:
			FightInventory(p)

		case 0:
			return

		}
	}
}

/*HeadDura:  0,
BodyDura:  0,
ShoeDura:  0,*/

func AttaqueMonstre(p *Personnage, m *Monstre) {
	r := 0
	// Logique d'attaque du monstre
	if r <= 3 {
		fmt.Printf("\n%s vous attaque et inflige %d points de dégâts\n", m.Nom, m.PointAttact)
		p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
		r++
	}
	if r >= 3 {
		ClearConsole()
		m.PointAttact = m.PointAttact * 2
		attaque := rand.Intn(3) + 1
		r++
		// Vérifie où le monstre a touché
		switch attaque {
		case 1:
			fmt.Println("Le monstre vous a touché à la tête !")
			p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
			r++
		case 2:
			fmt.Println("Le monstre vous a touché au corps !")
			p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
			r++
		case 3:
			fmt.Println("Le monstre vous a touché aux pieds !")
			p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
			r++

		}
	}
	if r >= 6 {
		m.PointAttact = m.PointAttact * 2
		attaque := rand.Intn(3) + 1
		// Vérifie où le monstre a touché
		switch attaque {
		case 1:
			fmt.Println("Le monstre vous a touché à la tête !")
			p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
			r++
		case 2:
			fmt.Println("Le monstre vous a touché au corps !")
			p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
			r++
		case 3:
			fmt.Println("Le monstre vous a touché aux pieds !")
			p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
			r++
		}
	}
}

func Poingdefer(p *Personnage, m *Monstre) {
	if p.Mana >= 0 {
		m.PvActuel -= 6
		p.Mana -= 5
		fmt.Println("\nVous avez utiliser 5 mana")
	} else {
		fmt.Println("\nVous n'avez pas assez de mana")
	}
}

func CheckDura(p *Personnage) {
	if p.Equipement.HeadDura <= 0 {
		BonusHp(p)
	} else if p.Equipement.HeadDura > 0 {
		fmt.Printf("\nIl reste %d de durabilité a votre casque", p.Equipement.HeadDura)
	}
	if p.Equipement.BodyDura <= 0 {
		BonusHp(p)
	} else if p.Equipement.BodyDura > 0 {
		fmt.Printf("\nIl reste %d de durabilité a votre Tunique", p.Equipement.BodyDura)
	}
	if p.Equipement.ShoeDura <= 0 {
		BonusHp(p)
	} else if p.Equipement.ShoeDura > 0 {
		fmt.Printf("\nIl reste %d de durabilité a vos Bottes", p.Equipement.ShoeDura)
	}
}

func CoupDePoing(p *Personnage, m *Monstre) {
	if p.Mana >= 0 {
		m.PvActuel -= 8
		p.Mana -= 5
		fmt.Println("\nVous avez utiliser 5 mana")
	} else {
		fmt.Println("\nVous n'avez pas assez de mana")
	}
}

func BouleDeFeu(p *Personnage, m *Monstre) {
	if p.Mana >= 0 {
		m.PvActuel -= 18
		p.Mana -= 12
		fmt.Println("\nVous avez utiliser 12 mana")
	} else {
		fmt.Println("\nVous n'avez pas assez de mana")
	}
}

func Tirperforant(p *Personnage, m *Monstre) {
	if p.Mana >= 0 {
		m.PvActuel -= 20
		p.Mana -= 10
		fmt.Println("\nVous avez utiliser 10 mana")
	} else {
		fmt.Println("\nVous n'avez pas assez de mana")
	}
}

func Lancerdehache(p *Personnage, m *Monstre) {
	if p.Mana >= 0 {
		m.PvActuel -= 20
		p.Mana -= 10
		fmt.Println("\nVous avez utiliser 10 mana")
	} else {
		fmt.Println("\nVous n'avez pas assez de mana")
	}
}

func Marteaudejustice(p *Personnage, m *Monstre) {
	if p.Mana >= 0 {
		m.PvActuel -= 20
		p.Mana -= 10
		fmt.Println("\nVous avez utiliser 10 mana")
	} else {
		fmt.Println("\nVous n'avez pas assez de mana")
	}
}

func AttaquePersonnage(p *Personnage, m *Monstre) {
	var Atak int
	fmt.Println("\nMenu d'attaque:")
	fmt.Println("\n1. Coup de poing")
	fmt.Println("\n2. Boule de Feu")
	if p.Inventaire["Poing de fer"] > 0 {
		fmt.Println("\n3. Poing de fer")
	} else {
		fmt.Println("\n3. Poing de fer (Indisponible -> lvl 5) ")
	}
	if p.Inventaire["Tir perforant"] > 0 {
		fmt.Println("\n4. Tir perforant")
	} else {
		fmt.Println("\n4. Tir perforant (Indisponible -> lvl 10) ")
	}
	if p.Inventaire["Lancer de hache"] > 0 {
		fmt.Println("\n5. Lancer de hache")
	} else {
		fmt.Println("\n5. Lancer de hache (Indisponible -> lvl 15) ")
	}
	if p.Inventaire["Marteau de justice"] > 0 {
		fmt.Println("\n6. Marteau de justice")
	} else {
		fmt.Println("\n6. Marteaudejustice (Indisponible -> lvl 20) ")
	}
	fmt.Print("Votre choix : ")
	fmt.Scanln(&Atak)

	switch Atak {
	case 1:
		CoupDePoing(p, m)
	case 2:
		BouleDeFeu(p, m)
	case 3:
		if p.Inventaire["Poing de fer"] > 0 {
			Poingdefer(p, m)
		} else {
			return
		}
	case 4:
		if p.Inventaire["Tir perforant"] > 0 {
			Tirperforant(p, m)
		} else {
			return
		}
	case 5:
		if p.Inventaire["Lanceur de hache"] > 0 {
			Lancerdehache(p, m)
		} else {
			return
		}
	case 6:
		if p.Inventaire["Marteau de justice"] > 0 {
			Marteaudejustice(p, m)
		} else {
			return
		}
	default:
		fmt.Println("Attaque invalide.")
	}
}
