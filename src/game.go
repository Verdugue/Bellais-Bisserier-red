package gamePlay

import (
    "fmt"
    "strings"
    "os"
    "os/exec"
    "runtime"
)

type Personnage struct {
    Nom            string
    Classe         string
    Niveau         int
    PointsVieMax   int
    PointsVieActuels int
    Inventaire     []struct {
        Nom     string
        Quantite int
    }
    Skills []string // Champ pour les compétences du personnage
}


func Init(nom, classe string, niveau, pvMax, pvActuels int) *Personnage {
    // Initialiser le personnage avec les valeurs spécifiées
    personnage := &Personnage{
        Nom:          nom,
        Classe:       classe,
        Niveau:       niveau,
        PointsVieMax: pvMax,
        PointsVieActuels: pvActuels,
        Inventaire: []struct {
            Nom     string
            Quantite int
        }{
            {"Potion de soin", 3},
        },
    }
    return personnage
}

func CharCreation() *Personnage {
    var nom string
    var classe string

    // Demander le nom de l'utilisateur et le valider
    for {
        fmt.Print("Choisissez un nom (uniquement des lettres) : ")
        fmt.Scanln(&nom)
        if isAlpha(nom) {
            break
        }
        fmt.Println("Le nom doit contenir uniquement des lettres.")
    }

    // Mettre le nom en majuscule pour la première lettre et en minuscules pour le reste
    nom = strings.Title(strings.ToLower(nom))

    // Demander la classe de l'utilisateur et initialiser les points de vie en conséquence
    for {
        fmt.Print("Choisissez une classe (Humain/Elfe/Nain) : ")
        fmt.Scanln(&classe)
        switch classe {
        case "Humain":
            return Init(nom, classe, 1, 100, 40)
        case "Elfe":
            return Init(nom, classe, 1, 80, 40)
        case "Nain":
            return Init(nom, classe, 1, 120, 60)
        default:
            fmt.Println("Classe invalide, veuillez choisir parmi Humain, Elfe ou Nain.")
        }
    }
}

func isAlpha(s string) bool {
    for _, char := range s {
        if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
            return false
        }
    }
    return true
}

func DisplayInfo(p *Personnage) {
    fmt.Println("\nInformations du personnage :")
    fmt.Println("Nom:", p.Nom)
    fmt.Println("Classe:", p.Classe)
    fmt.Println("Niveau:", p.Niveau)
    fmt.Println("Points de vie maximum:", p.PointsVieMax)
    fmt.Println("Points de vie actuels:", p.PointsVieActuels)
    fmt.Println("Inventaire:")
    for _, item := range p.Inventaire {
        fmt.Printf("%s x%d\n", item.Nom, item.Quantite)
    }
    fmt.Println("Compétences:")
    for _, skill := range p.Skills {
        fmt.Println(skill)
    }
}

func ClearConsole() {
    if runtime.GOOS == "windows" {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    } else {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func AccessInventory(p *Personnage) {
    for {
        fmt.Println("\nInventaire du personnage :")
        for i, item := range p.Inventaire {
            fmt.Printf("%d. %s x%d\n", i+1, item.Nom, item.Quantite)
        }
        fmt.Println("0. Quitter")

        var choix int
        fmt.Print("Votre choix : ")
        fmt.Scanln(&choix)

        if choix == 0 {
            // Quitter le sous-menu
            break
        } else if choix > 0 && choix <= len(p.Inventaire) {
            item := &p.Inventaire[choix-1]
            if item.Quantite > 0 {
                fmt.Printf("Vous avez choisi d'utiliser %s.\n", item.Nom)
                // Ajouter ici la logique pour utiliser l'objet (par exemple, utiliser une potion de santé)
                UseItem(p, item.Nom)
                p.PointsVieActuels += 50
                if p.PointsVieActuels > p.PointsVieMax {
                    p.PointsVieActuels = p.PointsVieMax
                }
                fmt.Printf("Points de vie actuels : %d/%d\n", p.PointsVieActuels, p.PointsVieMax)
                if item.Quantite == 0 {
                    fmt.Printf("Vous n'avez plus de %s dans votre inventaire.\n", item.Nom)
                }
            } else {
                fmt.Printf("Vous n'avez plus de %s dans votre inventaire.\n", item.Nom)
            }
        } else {
            fmt.Println("Choix invalide, veuillez réessayer.")
        }
    }
}



func UseItem(p *Personnage, itemName string) {
    for _, item := range p.Inventaire {
        if item.Nom == itemName {
            // Vous pouvez ajouter ici la logique pour utiliser l'objet
            // Par exemple, si l'objet est une potion de santé, vous pouvez augmenter les points de vie du personnage.
            // Ici, nous affichons simplement un message pour illustrer l'idée.
            fmt.Printf("Utilisé %s\n", itemName)
            item.Quantite--
            return
        }
    }
    fmt.Printf("%s n'est pas dans votre inventaire.\n", itemName)
}

func TakePot(p *Personnage) {
    // Parcourir l'inventaire pour trouver une potion de santé
    for _, item := range p.Inventaire {
        if item.Nom == "Potion de soin" {
            if p.PointsVieActuels == p.PointsVieMax {
                fmt.Println("Vos points de vie sont déjà au maximum, vous ne pouvez pas utiliser la potion.")
            } else {
                p.PointsVieActuels += 50
                if p.PointsVieActuels > p.PointsVieMax {
                    p.PointsVieActuels = p.PointsVieMax
                }
                fmt.Println("Vous avez utilisé une potion de soin et récupéré 50 points de vie.")
            }
            return
        }
    }
    fmt.Println("Vous n'avez pas de potion de soin dans votre inventaire.")
}

func MarchandMenu(p *Personnage) {
    potionsSoin := 1 // Nombre initial de potions de soin disponibles chez le marchand

    for {
        fmt.Println("Menu du marchand :")
        fmt.Println("1. Potion de soin (Gratuite)")
        fmt.Println("0. Quitter")

        var choix int
        fmt.Print("Votre choix : ")
        fmt.Scanln(&choix)

        switch choix {
        case 1:
            if potionsSoin > 0 {
                // Ajouter une potion de soin à l'inventaire du personnage
                p.Inventaire = append(p.Inventaire, struct{ Nom string; Quantite int }{"Potion de soin", 1})
                fmt.Println("Vous avez acheté une Potion de soin.")
                // Réduire le nombre de potions de soin disponibles chez le marchand
                potionsSoin--
            } else {
                fmt.Println("Le marchand n'a plus de potions de soin en stock.")
            }
        case 0:
            fmt.Println("Vous quittez le menu du marchand.")
            return
        default:
            fmt.Println("Choix invalide, veuillez réessayer.")
        }
    }
}