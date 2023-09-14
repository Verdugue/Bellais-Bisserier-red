package red

import "fmt"

// Détails personnages

type personnage interface { 
	name  string
	class  string
	level  int
	pv_max  int
	pv  int
	inventory  []string
}

// Créations personnages

func(p personnage) DetailsPerso{
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Class %s\n", p.class)
	fmt.Printf("Level %s\n", p.level)
	fmt.Printf("Pv Max %s\n", p.pv_max)
	fmt.Printf("Pv %s\n", p.pv)
	fmt.Printf("Inventory %s\n", p.inventory)
}