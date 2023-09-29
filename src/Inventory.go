package gamePlay

import (
	"fmt"
	"time"
)

var limit int

func (p *Personnage) RemoveInventory(item string) {
	if p.Inventaire[item] > 0 {
		p.Inventaire[item]--
		fmt.Println("Vous avez retiré 1", item, "de votre inventaire.")
	} else {
		fmt.Println("Vous n'avez pas de", item, "de votre inventaire.")
	}
}

func (p *Personnage) CheckInventoryLimit() bool {
	limit = 10
	nbItem := 0
	for _, item := range p.Inventaire {
		fmt.Println(item)
		nbItem += item
	}
	return nbItem >= limit
}

func UpgradeInventorySlot(p *Personnage) {
	limit = limit + 10
	fmt.Println("Vous avez maintenant :\n", 10+limit, "place dans votre inventaire.")
}

func AccessInventory(p *Personnage) {
	for {
		ClearConsole()
		var index int
		fmt.Println(" __  .__   __. ____    ____  _______ .__   __. .___________.  ______   .______      ____    ____")
		fmt.Println("|  | |  \\ |  | \\   \\  /   / |   ____||  \\ |  | |           | /  __  \\  |   _  \\     \\   \\  /   /")
		fmt.Println("|  | |   \\|  |  \\   \\/   /  |  |__   |   \\|  | `---|  |----`|  |  |  | |  |_)  |     \\   \\/   /")
		fmt.Println("|  | |  . `  |   \\      /   |   __|  |  . `  |     |  |     |  |  |  | |      /       \\_    _/")
		fmt.Println("|  | |  |\\   |    \\    /    |  |____ |  |\\   |     |  |     |  `--'  | |  |\\  \\----.    |  |")
		fmt.Println("|__| |__| \\__|     \\__/     |_______||__| \\__|     |__|      \\______/  | _| `._____|    |__|")

		fmt.Println("\nInventaire du personnage :")
		fmt.Println("Votre choix : ")
		for key, value := range p.Inventaire {
			index++
			fmt.Printf("%d. %s: %d\n", index, key, value)
		}
		fmt.Println("0. Quitter")

		var pepe int
		fmt.Scanln(&pepe)
		switch pepe {
		case 1:
			TakePot(p)
		case 2:
			PoisonPot(p)
		case 3:
			ManaPot(p)
		case 4:
			Equipe(p)
		case 5:
			Gemme(p)
		case 0:
			fmt.Println("Vous quittez l'inventaire.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func ManaPot(p *Personnage) {
	if !(p.Mana >= 100 && p.Inventaire["Potion de Mana"] > 0) {
		p.Mana += 50
		fmt.Printf("\nVous avez gagné 50 mana")
	} else {
		fmt.Println("")
	}
}

func TakePot(p *Personnage) {
	for key, value := range p.Inventaire {
		//regarde s'il l'objet selectionner est bien une potion de soin et s'il en reste et si mes hps ne sont pas au max
		if key == "Potion de soin" && value > 0 {
			if p.PointsVieActuels < p.PointsVieMax {
				//si oui, alors ca l'utilise
				p.PointsVieActuels += 50
				fmt.Printf("Vous avez utilisé une Potion de soin et récupéré 50 points de vie.\n")
				p.Inventaire["Potion de soin"]--
			} else if p.PointsVieActuels > p.PointsVieMax {
				fmt.Println("Vos hp sont au max.")

			}
			//sinon ca print que mes hps sont au max
			fmt.Println("Vous n'avez pas de potion de soin.")
			return

		}
	}
	fmt.Println("Vous n'avez plus de Potion de soin dans votre inventaire.") // Trouver une potion de soin dans l'inventaire
}

func PoisonPot(p *Personnage) {
	if p.Inventaire["Potion de Poison"] > 0 {
		fmt.Println("Ceci est une potion de poison, vous pourrez l'utiliser pour attaquer vos adversaires.")
	} else {
		fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire.")
	}
}

func FightPoisonPot(p *Personnage, m *Monstre) {
	duration := 3 // Durée en secondes de l'empoisonnement
	damagePerSecond := 10

	fmt.Printf("Vous avez empoisonné le monstre pendant %d secondes!\n", duration)
	for i := 0; i < duration; i++ {
		m.PvActuel -= damagePerSecond
		if m.PvActuel <= 0 {
			break
		}
		// Afficher les points de vie actuels sur les points de vie maximum
		fmt.Printf("Points de vie actuels : %d/%d\n", m.PvActuel, m.PvMaximum)
		// Attendre 1 seconde avant d'infliger les dégâts suivants
		time.Sleep(1 * time.Second)
	}
	fmt.Println("L'effet empoisonnement s'est dissipé.")
}

func FightInventory(p *Personnage) {
	for {
		var index int
		fmt.Println("\nInventaire de combat:")
		fmt.Println("Votre choix : ")
		for key, value := range p.Inventaire {
			index++
			fmt.Printf("%d. %s: %d\n", index, key, value)
		}
		fmt.Println("0. Quitter")

		var pepe int
		fmt.Scanln(&pepe)
		switch pepe {
		case 1:
			TakePot(p)
		case 2:
			FightPoisonPot(p, &p.Monstre)
		case 3:
			ManaPot(p)
		case 4:
			Gemme(p)
		case 0:
			fmt.Println("Vous quittez l'inventaire.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
