package gamePlay

import "math/rand"

type Monstre struct {
	Nom         string
	PvMaximum   int
	PvActuel    int
	PointAttact int
	Initiative  int
	Xp          int
}

var (
	Gobelin = Monstre{
		Nom:         "Gobelin d'entrainement",
		PvMaximum:   40,
		PvActuel:    40,
		PointAttact: 5,
		Initiative:  rand.Intn(10) + 1, // Initiative aléatoire entre 1 et 10
		Xp:          rand.Intn(5) + 1,
	}
	Orc = Monstre{
		Nom:         "Orc",
		PvMaximum:   50,
		PvActuel:    50,
		PointAttact: 8,
		Initiative:  rand.Intn(10) + 1, // Initiative aléatoire entre 1 et 10Dragon = Monstre
		Xp:          rand.Intn(10) + 1,
	}
	Dragon = Monstre{
		Nom:         "Dragon",
		PvMaximum:   100,
		PvActuel:    100,
		PointAttact: 15,
		Initiative:  rand.Intn(10) + 1, // Initiative aléatoire entre 1 et 10
		Xp:          rand.Intn(20) + 1,
	}
)
