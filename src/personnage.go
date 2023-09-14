package red

import "fmt"

// Détails personnages

type Personnage struct {
    name            string
    class          	string
    level         	int
    pv_max       	int
    pv       		int
    inventory      	[]string
}

// Créations personnages

func (p Personnage) DetailsPerso() {
    fmt.Printf("Name: %s\n", p.name)
    fmt.Printf("Class: %s\n", p.class)
    fmt.Printf("Level: %d\n", p.level)
    fmt.Printf("Pv Max: %d\n", p.pv_max)
    fmt.Printf("Pv: %d\n", p.pv)
    fmt.Printf("Inventory: %v\n", p.inventory)
}
