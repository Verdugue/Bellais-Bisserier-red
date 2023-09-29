package gamePlay

import (
	"fmt"
	"math/rand"
	"time"
)

func Jardin(p *Personnage, m *Monstre) {
	rand.Seed(time.Now().UnixNano())
	var planter int
	fmt.Printf("\nBravo, vous avez battu  %s, vous avez donc 10 Pourcent de chance d'obtenir une graine Spéciale", m.Nom)
	fmt.Println("\nVous pourrez obtenir des compétences spéciales.")
	fmt.Println("\nVoulez-vous la planter ?")
	fmt.Println("\n1. Oui.")
	fmt.Println("\n2. Non.")
	fmt.Scanln(&planter)
	switch planter {
	case 1:
		fmt.Println("\nVous avez planter votre graine Spéciale")
		randomNumber := rand.Intn(10) + 1
		if randomNumber == 1 {
			fmt.Println("\nVous avez reussir a planter la graine, Vous obtenez :")
			p.PointsVieMax += rand.Intn(100) + 1 // Points de vie max entre 1 et 100
			p.Argent += rand.Intn(100) + 1       // Argent entre 1 et 1000
			p.Xp += rand.Intn(50) + 1            // Expérience entre 1 et 10000
		}
	}
}

func Competences(p *Personnage) {
	if p.Niveau >= 5 {
		p.Inventaire["Poing de Fer"]++
		fmt.Println("\nVous avez obtenu une competences spéciale: Poing de Fer")
	}
	if p.Niveau >= 10 {
		p.Inventaire["Lancer de hache"]++
		fmt.Println("\nVous avez obtenu une compétences spéciale: Lancer de hache.")
	}
	if p.Niveau >= 15 {
		p.Inventaire["Tir perforant"]++
		fmt.Println("\nVous avez obtenu une compétences spéciale: Tir perforant.")
	}
	if p.Niveau >= 20 {
		p.Inventaire["Marteau de justice"]++
		fmt.Println("\nVous avez obtenu une compétences spéciale: Marteau de justice.")
	}
}
