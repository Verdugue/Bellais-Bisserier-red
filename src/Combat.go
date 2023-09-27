package gamePlay

import (
	"fmt"
	"math/rand"
)

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
		fmt.Printf("Vous avez encore %d mana", p.Mana)
		fmt.Printf("%s - PV: %d / %d\n", m.Nom, m.PvActuel, m.PvMaximum)
		fmt.Println("\nVous vous battez contre :", m.Nom)
		fmt.Println("\nMenu de combat:")
		fmt.Println("\n1. Attaquer")
		fmt.Println("\n2. Acceder a l'inventaire")
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
					p.Xp = p.Xp + m.Xp
					return
				}
				AttaquePersonnage(p, m)

				if m.PvActuel <= 0 {
					fmt.Printf("%s a été vaincu !\n", m.Nom)
					p.Niveau = p.Niveau + m.Xp
					LevelUp(p)
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

/*HeadDura:  0,
BodyDura:  0,
ShoeDura:  0,*/

func AttaqueMonstre(p *Personnage, m *Monstre) {
	r := 0
	// Logique d'attaque du monstre
	if r <= 3 {
		attaque := rand.Intn(3) + 1
		// Vérifie où le monstre a touché
		switch attaque {
		case 1:
			if !(p.Equipement.Head == "") {
				fmt.Println("Le monstre vous a touché à la tête !")
				p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
				fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
				r++
				p.Equipement.HeadDura -= m.PointAttact
				CheckDura(p)
				BonusHp(p)
			}
		case 2:
			if !(p.Equipement.Body == "") {
				fmt.Println("Le monstre vous a touché à la tête !")
				p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
				fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
				r++
				p.Equipement.BodyDura -= m.PointAttact
				CheckDura(p)
				BonusHp(p)
			}
		case 3:
			if !(p.Equipement.Shoe == "") {
				fmt.Println("Le monstre vous a touché à la tête !")
				p.PointsVieActuels = p.PointsVieActuels - m.PointAttact
				fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m.Nom, p.Nom, m.PointAttact)
				r++
				p.Equipement.ShoeDura -= m.PointAttact
				CheckDura(p)
				BonusHp(p)
			}
		}
	}

	if r >= 3 {
		ClearConsole()
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

func CheckDura(p *Personnage) {
	if p.Equipement.HeadDura <= 0 {
		BonusHp(p)
	} else if p.Equipement.HeadDura > 0 {
		fmt.Printf("\nIl reste %d de durabilité a votre casque", p.Equipement.HeadDura)
	}
	if p.Equipement.BodyDura <= 0 {
		BonusHp(p)
	} else if p.Equipement.BodyDura > 0 {
		fmt.Printf("\nIl reste %d de durabilité a votre casque", p.Equipement.BodyDura)
	}
	if p.Equipement.ShoeDura <= 0 {
		BonusHp(p)
	} else if p.Equipement.ShoeDura > 0 {
		fmt.Printf("\nIl reste %d de durabilité a votre casque", p.Equipement.ShoeDura)
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

func AttaquePersonnage(p *Personnage, m *Monstre) {
	var Atak int
	fmt.Println("\nMenu d'attaque:")
	fmt.Println("1. Coup de poing")
	fmt.Println("2. Boule de Feu")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&Atak)

	switch Atak {
	case 1:
		CoupDePoing(p, m)
	case 2:
		BouleDeFeu(p, m)
	case 3:
		PoisonPot(p, m)
	default:
		fmt.Println("Attaque invalide.")
	}
}
