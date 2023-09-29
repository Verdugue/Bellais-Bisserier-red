package gamePlay

import "fmt"

func Forgeron(p *Personnage) {
	var demo int
	fmt.Println(" _______   ______   .______        _______  _______ .______        ______   .__   __.")
	fmt.Println("|   ____| /  __  \\  |   _  \\      /  _____||   ____||   _  \\      /  __  \\  |  \\ |  |")
	fmt.Println("|  |__   |  |  |  | |  |_)  |    |  |  __  |  |__   |  |_)  |    |  |  |  | |   \\|  |")
	fmt.Println("|   __|  |  |  |  | |      /     |  | |_ | |   __|  |      /     |  |  |  | |  . `  |")
	fmt.Println("|  |     |  `--'  | |  |\\  \\----.|  |__| | |  |____ |  |\\  \\----.|  `--'  | |  |\\   |")
	fmt.Println("|__|      \\______/  | _| `._____| \\______| |_______|| _| `._____| \\______/  |__| \\__|")

	fmt.Println("\nMenu du Forgeron:")
	fmt.Println("1. Souhaitez-vous forger un équipement ? ")
	fmt.Println("2. Recycler un équipement? ")
	fmt.Println("0. Quitter. ")
	fmt.Scanln(&demo)
	switch demo {
	case 1:
		var choix int
		fmt.Println("Que souhaitez-vous faire?: ")
		fmt.Println("1. Fabriquer le Chapeau de l'aventurier. ")
		fmt.Println("2. Fabriquer la Tunique de l'aventurier. ")
		fmt.Println("3. Fabriquer les Bottes de l'aventurier. ")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("Vous avez choisi de fabriquer LE CHAPEAU DE L'AVENTURIER !")
			fmt.Println("Vous aurez besoin pour cela de : 1 Plume de Corbeau et 1 Cuir de sanglier.")
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
			}
		case 2:
			fmt.Println("Vous avez choisi de fabriquer TUNIQUE DE L'AVENTURIER !")
			fmt.Println("Vous aurez besoin pour cela de : 1 Fourrure de loup et 1 Peau de troll.")
			var bebe int
			fmt.Print("Voulez-vous continuer : OUI (1) / NON (2) ")
			fmt.Scanln(&bebe)

			switch bebe {
			case 1:
				if p.Inventaire["Fourrure de loup"] >= 1 && p.Inventaire["Peau de troll"] >= 1 {
					// Vous avez les ingrédients nécessaires, fabriquez une tunique
					fmt.Println("Tunique de l'aventurier fabriquée !")
					p.RemoveInventory("Fourrure de loup")
					p.RemoveInventory("Peau de troll")
					p.Equipement.Body = "Tunique de l'aventurier"
					Equipe(p)
					p.Argent -= 5
				} else {
					fmt.Println("Vous n'avez pas les ingrédients nécessaires pour fabriquer la tunique.")
				}
			}
		case 3:
			fmt.Println("Vous avez choisi de fabriquer BOTTES DE L'AVENTURIER !")
			fmt.Println("Vous aurez besoin pour cela de : 1 Fourrure de loup et 1 Cuir de Sanglier.")
			var dede int
			fmt.Print("Voulez-vous continuer : OUI (1) / NON (2) ")
			fmt.Scanln(&dede)

			switch dede {
			case 1:
				if p.Inventaire["Fourrure de loup"] >= 1 && p.Inventaire["Cuir de Sanglier"] >= 1 {
					// Vous avez les ingrédients nécessaires, fabriquez des bottes
					fmt.Println("Bottes de l'aventurier fabriquées !")
					p.RemoveInventory("Fourrure de loup")
					p.RemoveInventory("Cuir de Sanglier")
					p.Equipement.Shoe = "Bottes de l'aventurier"
					p.Argent -= 5
					Equipe(p)
				} else {
					fmt.Println("Vous n'avez pas les ingrédients nécessaires pour fabriquer les bottes.")
				}
			}
		}
	case 2:
		var lil int
		fmt.Println("Que voulez-vous recycler?")
		fmt.Println("1. Chapeau de l'aventurier")
		fmt.Println("2. Tunique de l'aventurier")
		fmt.Println("3. Bottes de l'aventurier")
		fmt.Println("0. Quitter")
		fmt.Scanln(&lil)
		switch lil {
		case 1:
			if p.Equipement.Head == "Chapeau de l'aventurier" {
				fmt.Println("Vous avez recyclé le chapeau de l'aventurier.")
				p.Equipement.Head = "Chapeau de l'aventurier"
				p.Inventaire["Plume de Corbeau"]++
				p.Inventaire["Cuir de sanglier"]++
			} else {
				fmt.Println("Vous n'avez pas de chapeau de l'aventurier à recycler.")
			}
		case 2:
			if p.Equipement.Body == "Tunique de l'aventurier" {
				fmt.Println("Vous avez recyclé la tunique de l'aventurier.")
				p.Equipement.Body = "Tunique de l'aventurier"
				p.Inventaire["Fourrure de loup"]++
				p.Inventaire["Peau de troll"]++
			} else {
				fmt.Println("Vous n'avez pas de tunique de l'aventurier à recycler.")
			}
		case 3:
			if p.Equipement.Shoe == "Bottes de l'aventurier" {
				fmt.Println("Vous avez recyclé les bottes de l'aventurier.")
				p.Equipement.Shoe = "Bottes de l'aventurier"
				p.Inventaire["Fourrure de loup"]++
				p.Inventaire["Cuir de sanglier"]++
			} else {
				fmt.Println("Vous n'avez pas de bottes de l'aventurier à recycler.")
			}
		case 0:
			return
		}
	}
}
