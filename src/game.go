package gamePlay

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Personnage struct {
	Nom              string
	Classe           string
	Niveau           int
	PointsVieMax     int
	PointsVieActuels int
	Inventaire       map[string]int 
	Skills []string // Champ pour les compétences du personnage
}

func Init(nom, classe string, niveau, pvMax, pvActuels int) *Personnage {
	// Initialiser le personnage avec les valeurs spécifiées
	personnage := &Personnage{
		Nom:              nom,
		Classe:           classe,
		Niveau:           niveau,
		PointsVieMax:     pvMax,
		PointsVieActuels: pvActuels,
		Inventaire: map[string]int{
            "Potion de soin": 3,
		},
		Skills: []string{"Coup de poing"}, // Ajoutez le sort de base "Coup de poing"
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
		for key, value := range p.Inventaire {
			fmt.Printf(" %s: %d\n", key, value)
		}
		fmt.Println("0. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)
        switch choix {
        case 1:
            TakePot(p)
        case 2:
            PoisonPot(p)
        case 0:
			fmt.Println("Vous quittez l'inventaire.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
        }
	}
}

func TakePot(p *Personnage) {
	for key, value := range p.Inventaire {
		if key == "Potion de soin" && value > 0 {
            fmt.Printf("Vous avez utilisé une Potion de soin et récupéré 50 points de vie.\n")
			p.PointsVieActuels += 50

				// Décrémenter la quantité de potions de soin
			p.Inventaire["Potion de soin"]--

				// Vérifier si les points de vie actuels ne dépassent pas les points de vie maximum
			if p.PointsVieActuels > p.PointsVieMax {
				p.PointsVieActuels = p.PointsVieMax
				}

			fmt.Printf("Points de vie actuels : %d/%d\n", p.PointsVieActuels, p.PointsVieMax)
			return
		}
	}
	fmt.Println("Vous n'avez plus de Potion de soin dans votre inventaire.") // Trouver une potion de soin dans l'inventair
}
func Forgeron(p *Personnage) {
    
	fmt.Println("\nMenu du Forgeron:")
	fmt.Println("Chapeau de l'aventurier")
	fmt.Println("Tunique de l'aventurier")
    fmt.Println("bottes de l'aventurier")
    fmt.Println("0. Quitter")
    var choix int
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choix)

		switch choix {
		case 1:
            fmt.Println("Vous avez choisi de fabriquer LE CHEAPEAU DE L'AVENTUERIER ! Vous aurez besoin pour cela de : 1 Plume de Corbeau et 1 Cuir de sanglier.")
}
}

func (p *Personnage)Dead() {
	if p.PointsVieActuels <= 0 {
		fmt.Println("Vous êtes mort !")
		// Ressuscitez avec 50% de vos points de vie maximum
		p.PointsVieActuels = p.PointsVieMax / 2
		fmt.Printf("Vous avez été ressuscité avec %d points de vie.\n", p.PointsVieActuels)
	}
}

func PoisonPot(p *Personnage) {
	duration := 3 // Durée en secondes de l'empoisonnement
	damagePerSecond := 10

	fmt.Printf("Vous êtes empoisonné pendant %d secondes!\n", duration)
	for i := 0; i < duration; i++ {
		p.PointsVieActuels -= damagePerSecond
		if p.PointsVieActuels <= 0 {
			p.Dead()
            break
        }
		// Afficher les points de vie actuels sur les points de vie maximum
		fmt.Printf("Points de vie actuels : %d/%d\n", p.PointsVieActuels, p.PointsVieMax)

		// Attendre 1 seconde avant d'infliger les dégâts suivants
		time.Sleep(1 * time.Second)
	}
	fmt.Println("L'effet empoisonnement s'est dissipé.")
}

func spellBook(p *Personnage) {
	// Vérifier si le personnage n'a pas déjà appris "Boule de feu"
	for _, skill := range p.Skills {
		if skill == "Boule de feu" {
			fmt.Println("Vous avez déjà appris le sort 'Boule de feu'.")
			return
		}
	}

	// Si le personnage n'a pas déjà appris le sort, ajoutez-le à ses compétences
	p.Skills = append(p.Skills, "Boule de feu")
	fmt.Println("Vous avez appris le sort 'Boule de feu' !")
}
