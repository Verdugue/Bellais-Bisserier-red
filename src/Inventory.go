package gamePlay

import "fmt"

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
	fmt.Println("Vous avez maintenant \n", limit, "place dans votre inventaire.")
}

func AccessInventory(p *Personnage) {
	for {
		var index int
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
			PoisonPot(p, &p.Monstre)
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

func Tuto() {
	fmt.Println("\nVous avez rejoint le tutoriel.")
	fmt.Println("\nDans ce jeux vous Pourrez :")
	fmt.Println("\nDans ce jeux vous pourrez effectuer des choix avec 1, 2, 3...")
	fmt.Println("\nAjouter et Retirer à votre inventaire des item, via Le marchand ou le Forgeron")
	fmt.Println("\nVous battre contre des Monstres")
	fmt.Println("\nGagnez des niveau pour debloquer de nouvelles Compétences en battant des monstres!")
	fmt.Println("\nEn battant des monstres vous pourrez avoir 10 Pourcent de chance d'obtenir un item rare !")
	fmt.Println("\n 1. avancer ...")
	var quitter int
	fmt.Scanln(&quitter)
	switch quitter {
	case 1:
		return
	default:
		break
	}
}
