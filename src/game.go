package game

import (
	"fmt"
    "time"
    "unicode"
)

type Character struct {
	name      string
	class     string
	level     int
	pv_max    int
	pv        int
    spells    []string
	inventory []struct {
		name   string
		amount int
	}
}

func charCreation() *Character {
    var name string
    var class string

    // Demandez le nom de l'utilisateur
    for {
        fmt.Print("Choisissez votre nom (uniquement des lettres): ")
        fmt.Scanln(&name)

        // Vérifiez si le nom ne contient que des lettres
        if isAlpha(name) {
            break
        } else {
            fmt.Println("Le nom doit contenir uniquement des lettres.")
        }
    }

    // Demandez la classe de l'utilisateur
    for {
        fmt.Print("Choisissez votre classe (Humain, Elfe ou Demon): ")
        fmt.Scanln(&class)

        // Vérifiez si la classe est valide
        if class == "Humain" || class == "Elfe" || class == "Demon" {
            break
        } else {
            fmt.Println("Classe invalide. Choisissez parmi Humain, Elfe ou Demon.")
        }
    }

    // Initialisez le personnage avec les paramètres choisis par l'utilisateur
    character := Init(name, class)

    return character
}

func isAlpha(str string) bool {
    for _, char := range str {
        if !unicode.IsLetter(char) {
            return false
        }
    }
    return true
}

func Init(name string, class string) *Character {
    character := &Character{
        name:   name,
        class:  class,
        level:  1,
        pv_max: 100,
        pv:     40,
        inventory: []struct {
            name   string
            amount int
        }{
            {"Potions de soin", 3},
        },
        spells: []string{"Coup de poing"}, // Initialisez la liste des sorts avec "Coup de poing"
    }
    return character
}

func DisplayInfo(character *Character) {
	fmt.Printf("Nom: %s\n", character.name)
	fmt.Printf("Classe: %s\n", character.class)
	fmt.Printf("Niveau: %d\n", character.level)
	fmt.Printf("PV maximum: %d\n", character.pv_max)
	fmt.Printf("PV actuels: %d\n", character.pv)
	fmt.Println("Inventaire:")
	for _, item := range character.inventory {
		fmt.Printf("Nom: %s, Quantité: %d\n", item.name, item.amount)
	}
}

func AccessInventory(character *Character) {
    for {
        fmt.Println("Inventaire du personnage:")
        for i, item := range character.inventory {
            fmt.Printf("%d. %s (Quantité: %d)\n", i+1, item.name, item.amount)
        }
        fmt.Println("\nMenu de l'inventaire:")
        fmt.Println("1. Utiliser un objet")
        fmt.Println("2. Visiter le marchand")
        fmt.Println("3. Retour au menu principal")

        var choice int
        fmt.Print("Choisissez une option: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            useItem(character)
        case 2:
            visitMerchant(character) // Visiter le marchand
        case 3:
            return // Retour au menu principal
        default:
            fmt.Println("Option invalide")
        }
    }
}

func accessSubInventory(character *Character) {
    for {
        fmt.Println("\nMenu de l'inventaire:")
        fmt.Println("1. Utiliser une Potion de soin (Revenir en arrière)") // Ajoutez "Revenir en arrière" dans l'option
        fmt.Println("2. Retour au menu précédent")

        var choice int
        fmt.Print("Choisissez une option: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            useHealthPotion(character) // Appel d'une fonction spécifique pour utiliser une Potion de soin
        case 2:
            return // Retour au menu précédent
        default:
            fmt.Println("Option invalide")
        }
    }
}

func useHealthPotion(character *Character) {
    // Ajoutez ici la logique pour utiliser une Potion de soin
    fmt.Println("Vous avez utilisé une Potion de soin.")
    // N'oubliez pas de mettre à jour la quantité de Potions de soin dans l'inventaire
}

func onlyContainsHealthPotions(character *Character) bool {
    for _, item := range character.inventory {
        if item.name != "Potions de soin" {
            return false
        }
    }
    return true
}

func useItem(character *Character) {
    for {
        fmt.Println("Menu de l'inventaire:")
        fmt.Println("1. Utiliser un objet")
        fmt.Println("2. Revenir en arrière")

        var choice int
        fmt.Print("Choisissez une option: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            fmt.Println("Choisissez un objet à utiliser (1, 2, etc.):")
            for i, item := range character.inventory {
                fmt.Printf("%d. %s (Quantité: %d)\n", i+1, item.name, item.amount)
            }

            var itemIndex int
            fmt.Scanln(&itemIndex)

            if itemIndex >= 1 && itemIndex <= len(character.inventory) {
                item := &character.inventory[itemIndex-1]
                switch item.name {
                case "Potions de soin":
                    if item.amount > 0 {
                        takePot(character)
                        item.amount-- // Retire une potion de l'inventaire
                        fmt.Println("Vous avez utilisé une Potion de soin.")
                        fmt.Printf("Vous avez reçu 50 points de vie. PV actuels : %d\n", character.pv)
                    } else {
                        fmt.Println("Vous n'avez pas de Potion de soin.")
                    }
                case "Potion de poison":
                    if item.amount > 0 {
                        poisonPot(character) // Applique les dégâts de poison
                        removeInventory(character, "Potion de poison", 1) // Retire la Potion de poison de l'inventaire
                        fmt.Println("Vous avez utilisé une Potion de poison.")
                        fmt.Printf("Vous avez subi des dégâts de poison. PV actuels : %d\n", character.pv)
                    } else {
                        fmt.Println("Vous n'avez pas de Potion de poison.")
                    }
                default:
                    fmt.Println("Impossible d'utiliser cet objet.")
                }
            } else {
                fmt.Println("Option invalide")
            }
        case 2:
            return // Revenir en arrière
        default:
            fmt.Println("Option invalide")
        }
    }
}



func takePot(character *Character) {
    if character.pv < character.pv_max {
        beforeHeal := character.pv
        if character.pv+50 <= character.pv_max {
            character.pv += 50
        } else {
            character.pv = character.pv_max
        }
        afterHeal := character.pv
        healed := afterHeal - beforeHeal
        fmt.Printf("Vous avez utilisé une Potion de soin et récupéré %d HP.\n", healed)
        fmt.Printf("PV actuels: %d / PV max: %d\n", character.pv, character.pv_max)
    } else {
        fmt.Println("Impossible d'utiliser la potion car les PV sont au maximum.")
    }
}

func visitMerchant(character *Character) {
    fmt.Println("Bienvenue chez le marchand!")
    fmt.Println("Que souhaitez-vous acheter?")
    
    // Ajouter l'option de Potion de vie uniquement si l'inventaire contient autre chose que des Potions de soin
    if !onlyContainsHealthPotions(character) {
        fmt.Println("1. Potion de vie (gratuite)")
    }

    // Ajouter l'option de Potion de poison
    fmt.Println("2. Potion de poison")

    // Ajouter l'option de "Livre de Sort : Boule de Feu"
    fmt.Println("3. Livre de Sort : Boule de Feu")

    fmt.Println("4. Retour")

    var choice int
    fmt.Print("Choisissez une option: ")
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        if !onlyContainsHealthPotions(character) {
            buyItem(character, "Potion de vie", 1)
        }
    case 2:
        buyItem(character, "Potion de poison", 1) // Ajout de la Potion de poison à l'inventaire
    case 3:
        spellBook(character) // Appel de la fonction spellBook lorsque le joueur achète le livre de sort
    case 4:
        return // Retour au menu de l'inventaire
    default:
        fmt.Println("Option invalide")
    }
}



func buyItem(character *Character, itemToBuy string, amountToBuy int) {
    switch itemToBuy {
    case "Potion de vie":
        if amountToBuy > 0 {
            addInventory(character, "Potions de soin", amountToBuy)
            fmt.Printf("Vous avez acheté %d %s.\n", amountToBuy, itemToBuy)
            fmt.Printf("Vous avez reçu %d Potions de soin.\n", amountToBuy)
        } else {
            fmt.Println("Le marchand ne vend pas cet objet.")
        }
    case "Potion de poison":
        if amountToBuy > 0 {
            addInventory(character, "Potion de poison", amountToBuy)
            fmt.Printf("Vous avez acheté %d %s.\n", amountToBuy, itemToBuy)
            fmt.Printf("Vous avez reçu %d Potions de poison.\n", amountToBuy)
        } else {
            fmt.Println("Le marchand ne vend pas cet objet.")
        }
    case "Livre de Sort : Boule de Feu":
        if amountToBuy > 0 {
            spellBook(character, "Boule de Feu")
            fmt.Printf("Vous avez acheté le %s.\n", itemToBuy)
            fmt.Printf("Vous avez appris le sort : %s.\n", "Boule de Feu")
        } else {
            fmt.Println("Le marchand ne vend pas cet objet.")
        }
    default:
        fmt.Println("Le marchand ne vend pas cet objet.")
    }
}



func addInventory(character *Character, itemToAdd string, amountToAdd int) {
    for i, item := range character.inventory {
        if item.name == itemToAdd {
            character.inventory[i].amount += amountToAdd
            return
        }
    }
    // Si l'objet n'est pas déjà dans l'inventaire, ajoutez-le
    character.inventory = append(character.inventory, struct{ name string; amount int }{itemToAdd, amountToAdd})
}

// Fonction pour retirer un objet de l'inventaire
func removeInventory(character *Character, itemToRemove string, amountToRemove int) {
    for i, item := range character.inventory {
        if item.name == itemToRemove {
            if item.amount >= amountToRemove {
                character.inventory[i].amount -= amountToRemove
                if character.inventory[i].amount == 0 {
                    // Supprime l'objet de l'inventaire si la quantité atteint 0
                    character.inventory = append(character.inventory[:i], character.inventory[i+1:]...)
                }
                return
            } else {
                fmt.Println("Vous n'avez pas assez de cet objet dans l'inventaire.")
                return
            }
        }
    }
    fmt.Println("Cet objet n'est pas dans votre inventaire.")
}

func dead(character *Character) {
    if character.pv <= 0 {
        fmt.Println("Vous êtes mort !")
        // Réinitialise les points de vie à 50% du maximum
        character.pv = character.pv_max / 2
    }
}

func poisonPot(character *Character) {
    duration := 3 // Durée en secondes
    damagePerSecond := 10

    fmt.Printf("Vous avez été empoisonné pendant %d secondes!\n", duration)
    for i := 0; i < duration; i++ {
        character.pv -= damagePerSecond
        fmt.Printf("PV actuels: %d / PV max: %d\n", character.pv, character.pv_max)
        time.Sleep(1 * time.Second) // Pause d'une seconde entre chaque tick de poison
        dead(character) // Vérifiez si le personnage est mort à chaque tick
    }
}

func spellBook(character *Character, spellName string) {
    // Vérifiez si le sort est déjà dans la liste des sorts
    for _, spell := range character.spells {
        if spell == spellName {
            fmt.Println("Vous avez déjà appris ce sort.")
            return
        }
    }

    // Ajoutez le sort à la liste des sorts
    character.spells = append(character.spells, spellName)
    fmt.Printf("Vous avez appris le sort : %s\n", spellName)
}

func useSpell(character *Character, spellName string) {
    if spellName == "Coup de poing" {
        fmt.Println("Vous avez utilisé Coup de poing.")
    } else if spellName == "Boule de feu" {
        fmt.Println("Vous avez utilisé la compétence : Boule de feu.")
    } else {
        fmt.Println("Sort inconnu.")
    }
}

func useFireballSpell(character *Character) {
    // Ajoutez ici la logique pour utiliser la compétence "Boule de feu"
    fmt.Println("Vous avez utilisé la compétence : Boule de feu.")
    // Vous pouvez ajouter ici le code pour infliger des dégâts à une cible si nécessaire
}