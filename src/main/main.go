package main

func main() {
	// Création d'un personnage
	joueur := Personnage{
		Nom:        "Joueur1",
		Classe:     "Guerrier",
		Niveau:     1,
		PVMaximum:  100,
		PVActuels:  100,
		Inventaire: []string{"Épée", "Armure légère"},
	}

	// Affichage des détails du personnage
	joueur.AfficherDetails()
}
