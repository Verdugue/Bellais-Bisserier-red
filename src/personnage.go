package gamePlay

import (
	"fmt"
	"math/rand"
	"strings"
)

type Personnage struct {
	Nom              string
	Classe           string
	Niveau           int
	PointsVieMax     int
	Argent           int
	PointsVieActuels int
	Inventaire       map[string]int
	Skills           []string // Utilisez []string pour stocker plusieurs compétences
	Equipement       Equipement
	Monstre          Monstre
	Initiative       int
	Xp               int
	Mana             int
	Graine           int
}

type Equipement struct {
	Head      string
	Body      string
	Shoe      string
	HeadBonus int
	BodyBonus int
	ShoeBonus int
	HeadDura  int
	BodyDura  int
	ShoeDura  int
	HeadMax   int
	BodyMax   int
	ShoeMax   int
}

func Init(nom string, classe string, niveau int, pvMax int, pvActuels int, stuff Equipement) *Personnage {
	// Initialiser le personnage avec les valeurs spécifiées
	Personnage := &Personnage{
		Nom:              nom,
		Classe:           classe,
		Niveau:           niveau,
		PointsVieMax:     pvMax,
		PointsVieActuels: pvActuels,
		Argent:           100,
		Inventaire: map[string]int{
			"Potion de soin":   3,
			"Potion de Poison": 3,
		},
		Initiative: rand.Intn(10) + 1,
		Xp:         0,
		Skills:     []string{"Coup de poing"}, // Ajoutez le sort de base "Coup de poing"
		Mana:       100,
		Graine:     0,

		Equipement: Equipement{
			Head:      stuff.Head,
			Body:      stuff.Body,
			Shoe:      stuff.Shoe,
			HeadBonus: stuff.HeadBonus, // Utilisez les noms corrects de la structure Equipement
			BodyBonus: stuff.BodyBonus,
			ShoeBonus: stuff.ShoeBonus,
			HeadDura:  stuff.HeadDura,
			BodyDura:  stuff.BodyDura,
			ShoeDura:  stuff.ShoeDura,
		},
	}
	return Personnage
}

func CharCreation() *Personnage {
	equipement := Equipement{
		Head:      "",
		Body:      "",
		Shoe:      "",
		HeadBonus: 0,
		BodyBonus: 0,
		ShoeBonus: 0,
		HeadDura:  0,
		BodyDura:  0,
		ShoeDura:  0,
		HeadMax:   100,
		BodyMax:   100,
		ShoeMax:   100,
	}

	var nom string
	var classe string

	// Demander le nom de l'utilisateur et le valider
	for {
		fmt.Print("Choisissez un nom : ")
		fmt.Scanln(&nom)
		if IsAlpha(nom) {
			break
		}
		fmt.Println("Le nom doit contenir uniquement des lettres.")
	}

	// Mettre le nom en majuscule pour la première lettre et en minuscules pour le reste
	nom = strings.Title(strings.ToLower(nom))
	ClearConsole()
	// Demander la classe de l'utilisateur et initialiser les points de vie en conséquence
	for {
		fmt.Print("Choisissez une classe : ")
		fmt.Print("\n Humain : La classe Humain vous octroie 100 Points de vie maximum en partant de 40")
		fmt.Print("\n Elfe :  La classe Elfe vous octroie 80 Points de vie maximum en partant de 40")
		fmt.Print("\n Nain :  La classe Nain vous octroie 120 Points de vie maximum en partant de 60")
		fmt.Print("\nVotre choix : ")
		fmt.Scanln(&classe)
		switch classe {
		case "Humain":
			return Init(nom, classe, 1, 100, 40, equipement)
		case "Elfe":
			return Init(nom, classe, 1, 80, 40, equipement)
		case "Nain":
			return Init(nom, classe, 1, 120, 60, equipement)
		default:
			fmt.Println("Classe invalide, veuillez choisir parmi Humain, Elfe ou Nain.")
		}
	}
}

func Dead(p *Personnage) {
	if p.PointsVieActuels <= 0 {
		fmt.Println("Vous êtes mort !")
		// Ressuscitez avec 50% de vos points de vie maximum
		p.PointsVieActuels = p.PointsVieMax / 2
		fmt.Printf("Vous avez été ressuscité avec %d points de vie.\n", p.PointsVieActuels)
	}
}

func ManaUPgrade(p *Personnage) {
	if p.Niveau >= 5 {
		p.Mana += 20
	}
	if p.Niveau >= 10 {
		p.Mana += 50
	}
	if p.Niveau >= 15 {
		p.Mana += 100
	}
	if p.Niveau >= 20 {
		p.Mana += 200
	}
	if p.Niveau >= 25 {
		p.Mana += 500
	}
}

func LevelUp(p *Personnage) {
	if p.Xp == 10 {
		p.Niveau += 1
	}
}
