package gamePlay

import "fmt"

func Forgeron(p *Personnage) {
	fmt.Println("\nMenu du Forgeron:")
	fmt.Println("1. Chapeau de l'aventurier")
	fmt.Println("2. Tunique de l'aventurier")
	fmt.Println("3. bottes de l'aventurier")
	fmt.Println("0. Quitter")
	var choix int
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi de fabriquer LE CHEAPEAU DE L'AVENTUERIER !")
		fmt.Println(" Vous aurez besoin pour cela de : 1 Plume de Corbeau et 1 Cuir de sanglier.")
		var choise int
		fmt.Print("Voulez-vous continuer : OUI (1) / NON (2) ")
		fmt.Scanln(&choise)

		switch choise {
		case 1:
			if p.Inventaire["Plume de corbeau"] >= 1 && p.Inventaire["Cuir de sanglier"] >= 1 {
				// Vous avez les ingrédients nécessaires, fabriquez un chapeau
				fmt.Println("Chapeau de l'aventurier fabriqué !")
				p.RemoveInventory("Plume de corbeau")
				p.RemoveInventory("Cuir de sanglier")
				p.Equipement.Head = "Chapeau de l'aventurier"
				p.Argent -= 5
			} else {
				fmt.Println("Vous n'avez pas les ingrédients nécessaires pour fabriquer un chapeau.")
			}
		case 2:
			return
		}
	case 2:
		fmt.Println("Vous avez choisi de fabriquer TUNIQUE DE L'AVENTURIER !")
		fmt.Println("Vous aurez besoin pour cela de : 1 Fourrure de loup et 1 Peau de troll.")
		var bebe int
		fmt.Print("Voulez-vous continuer : OUI 1 / NON 2")
		fmt.Scanln(&bebe)

		switch bebe {
		case 1:
			if p.Inventaire["Fourrure de loup"] >= 1 && p.Inventaire["Peau de troll"] >= 1 {
				// Vous avez les ingrédients nécessaires, fabriquez un chapeau
				fmt.Println("Tunique de l'aventurier fabriqué !")
				p.RemoveInventory("Fourrure de loup")
				p.RemoveInventory("Peau de troll")
				p.Equipement.Body = "Tunique de l'aventurier"
				Equipe(p)
				p.Argent -= 5
			} else {
				fmt.Println("Vous n'avez pas les ingrédients nécessaires pour fabriquer la tunique.")
			}
		case 2:
			break
		case 3:
			fmt.Println("Vous avez choisi de fabriquer BOTTES DE L'AVENTURIER !")
			fmt.Println("Vous aurez besoin pour cela de :1 Fourrure de loup et 1 Cuir de Sanglier.")
			var dede int
			fmt.Print("Voulez-vous continuer : OUI 1 / NON 2")
			fmt.Scanln(&dede)

			switch dede {
			case 1:
				if p.Inventaire["Fourrure de loup"] >= 1 && p.Inventaire["Cuir de Sanglier"] >= 1 {
					// Vous avez les ingrédients nécessaires, fabriquez un chapeau
					fmt.Println("Bottes de l'aventurier fabriqué !")
					p.RemoveInventory("Fourrure de loup")
					p.RemoveInventory("Cuir de Sanglier")
					p.Equipement.Shoe = "Bottes de l'aventurier"
					p.Argent -= 5
					Equipe(p)
				} else {
					fmt.Println("Vous n'avez pas les ingrédients nécessaires pour fabriquer la tunique.")
				}
			}
		}
	case 0:
		return
	}
}
