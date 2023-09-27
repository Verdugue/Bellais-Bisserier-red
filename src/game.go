package gamePlay

import (
	"fmt"
	"math/rand"
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
	Argent           int
	PointsVieActuels int
	Inventaire       map[string]int
	Skills           []string // Utilisez []string pour stocker plusieurs compétences
	Equipement       Equipement
	Monstre          Monstre
	Initiative       int
	Xp               int
	Mana             int
}

func LevelUp(p *Personnage) {
	if p.Xp == 10 {
		p.Niveau += 1
	}
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

func InfoDebut(p *Personnage) {
	fmt.Println("-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-")
	fmt.Println("Voici votre personnage :")
	fmt.Println("Nom:", p.Nom)
	fmt.Println("Classe:", p.Classe)
	fmt.Println("Niveau:", p.Niveau)
	fmt.Println("-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-")

}

func IsAlpha(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false
		}
	}
	return true
}

func MonstreInfo(m *Monstre) {
	fmt.Println("Vous Combattez contre ", m.Nom)
	fmt.Println("Nom du monstre : ", m.Nom)
	fmt.Println("Points de vie maximum : ", m.PvMaximum)
	fmt.Println("Points de vie actuel : ", m.PvActuel)
	fmt.Println("Point d'attaque : ", m.PointAttact)
}

func DisplayInfo(p *Personnage) {
	fmt.Println("\nInformations du personnage :")
	fmt.Println("Nom:", p.Nom)
	fmt.Println("Classe:", p.Classe)
	fmt.Println("Niveau:", p.Niveau)
	fmt.Println("Points de vie maximum:", p.PointsVieMax)
	fmt.Println("Points de vie actuels:", p.PointsVieActuels)
	fmt.Println("Inventaire: Vous avez maintenant\n", limit, "place dans votre inventaire.")
	fmt.Println("Equipement:", p.Equipement.Head, p.Equipement.Body, p.Equipement.Shoe)
	fmt.Printf("\nVotre Xp%d", p.Xp)
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

func ManaPot(p *Personnage) {
	if !(p.Mana >= 100 && p.Inventaire["Potion de Mana"] > 0) {
		p.Mana += 50
		fmt.Printf("\nVous avez gagné 50 mana")
	} else {
		fmt.Println("")
	}
}

func TakePot(p *Personnage) {
	for key, value := range p.Inventaire {
		//regarde s'il l'objet selectionner est bien une potion de soin et s'il en reste et si mes hps ne sont pas au max
		if key == "Potion de soin" && value > 0 {
			if p.PointsVieActuels < p.PointsVieMax {
				//si oui, alors ca l'utilise
				p.PointsVieActuels += 50
				fmt.Printf("Vous avez utilisé une Potion de soin et récupéré 50 points de vie.\n")
				p.Inventaire["Potion de soin"]--
			} else if p.PointsVieActuels > p.PointsVieMax {
				fmt.Println("Vos hp sont au max.")

			}
			//sinon ca print que mes hps sont au max
			fmt.Println("Vous n'avez pas de potion de soin.")
			return

		}
	}
	fmt.Println("Vous n'avez plus de Potion de soin dans votre inventaire.") // Trouver une potion de soin dans l'inventaire
}

func PoisonPot(p *Personnage, m *Monstre) {
	duration := 3 // Durée en secondes de l'empoisonnement
	damagePerSecond := 10

	fmt.Printf("Vous avez empoisonné le monstre pendant %d secondes!\n", duration)
	for i := 0; i < duration; i++ {
		m.PvActuel -= damagePerSecond
		if m.PvActuel <= 0 {
			break
		}
		// Afficher les points de vie actuels sur les points de vie maximum
		fmt.Printf("Points de vie actuels : %d/%d\n", m.PvActuel, m.PvMaximum)
		// Attendre 1 seconde avant d'infliger les dégâts suivants
		time.Sleep(1 * time.Second)
	}
	fmt.Println("L'effet empoisonnement s'est dissipé.")
}

func (p *Personnage) Dead() {
	if p.PointsVieActuels <= 0 {
		fmt.Println("Vous êtes mort !")
		// Ressuscitez avec 50% de vos points de vie maximum
		p.PointsVieActuels = p.PointsVieMax / 2
		fmt.Printf("Vous avez été ressuscité avec %d points de vie.\n", p.PointsVieActuels)
	}
}

func Qui() {
	var quiqui int
	fmt.Println("\n1. Personnalité 1.")
	fmt.Println("\n2. Personnalité 2.")
	fmt.Println("\n3. Personnalité 3.")
	fmt.Println("\n0. Quitter.")
	fmt.Scanln(&quiqui)

	switch quiqui {
	case 1:
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@   A   B   B   A  @@@@@@@@@@@@@@@@@@@@@@@@@@#==**=@@@@")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#***=*++++:*@@@@")
		fmt.Println("@@@#*:--:-::=@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#=++**=##=*=*")
		fmt.Println("@@@@@*:::::+***+:+@@@@@@@@@*+++=#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#=++**=##=*=*+:*@@")
		fmt.Println("@@@#+*==##=*****++@@@@@@#*+*#===*=@@@@@@@@@@*--::::=@@@@@@@@@==+++++*******+:#@")
		fmt.Println("@@@++==*=====*+++++#@@@#=@@@@##====@@@@@@@@-.:*+*::::*@@@@@@@##+*=+:+*+::=**++@")
		fmt.Println("@@#+*===@=***+++*#@@#@@*:::::+::#@=#@@#-*+:++:::::::-#@@@@@#=+++++++:::*=*=+#")
		fmt.Println("@@*:*#*+*+:++:++=**#@@##@#+*+:+**#@==@@@@*-::+::::+----@@@@@@@#+++**++::**##=*+")
		fmt.Println("@@*++=*::+::++:+***#@@#@@*:::::+::#@=#@@#-*+:++:::::::-+#@@@@@#=**++::+***#@===")
		fmt.Println("@@#++*=:::::+*++#==#@#@@@#::::::::@@#=#@=-=*::+:::::*+-:#@@@@==##=+++=#*-==@=*+")
		fmt.Println("@@=+=*=#::::+++#@@@@####@@+::::::*@@#=@@@-+*::*++:::*:--=@@@@@@#@*:=*-:=##@@@#-")
		fmt.Println("@@=*###=::+++++##@@@###@@@=:::::=@@###@@#-:*+::::::*::::+@@@@@#:-:*#@==###=####")
		fmt.Println(":-:#@@@*:::::::*:+++#@@@###*=+++*###@@@W*-:+*++:+=#=+++:-=@@@@#=#*--.#==##@@##@")
		fmt.Println("-:::@@##==**+***==*=+==***=#=*:++*+*@@@@:-:+=*===@#*+++::+@@@@=@+--.+====##=#=#")
		fmt.Println("---:+=====***#@@#=#++++++++#=*::::::::*=-:+*#:+*#*++++++::#@@=@+-+-+***===##=**")
		fmt.Println(":::::+**====####=+:::::::::*==*:::+::+:--:+=+++**+::-:+:+++###+::-+=*=====#=+=*")
		fmt.Println("::::-:=***==#@#+:::::::::::+**::::++:*+::+=+++***:--::-::++==*:+++======***++=*")
		fmt.Println("@@+---+=*====:::::::++=*::++++::+:++++*+*=+:+**=*+::::::-:+*=*+++#***==+::*++=*")
		fmt.Println("#*----+#==+:::::::++*+*=++**=*:**++*==*=#:++**::***+:::::::=#***#=****=++++*+=*")
		fmt.Println("=:----*#*::::::::*@@@==++++++::+::+#@@@#=*++*+******+::::::=#=*=#=******+++++==")
		fmt.Println("+:---:*::::::++**#@@@#=+::::***==+=#@@#**+:+:--:::+=++:::::*@####@#=*+*====*+=#")
		fmt.Println("=:----:::-::++++*#@@@@=**:+*===**++=@@****===**++++=#++::::+=@#**==**+*==#===#@")
	case 2:
		fmt.Println("::::::::::::::::::::::::::::::::::::::::")
		fmt.Println("::::::::::  SPILBERG  ::::::::::::::::::")
		fmt.Println(":::::::::::::::::::+++::::::::::::::::::")
		fmt.Println(":::::::::::::+++++++*****+::::::::::::::")
		fmt.Println(":::::::::::******+:::::::+*+++::::::::::")
		fmt.Println("::::::::::****+::::::::::++*++::::::::::")
		fmt.Println(":::::::::++++:::::--::::+++**++:::::::::")
		fmt.Println(":::::::::+++++::::::::::++***+::::::::::")
		fmt.Println(":::::::::+**++:+++:::+++*+**++::::::::::")
		fmt.Println(":::::::::+++*#=*==#=##=#==@=*=*:::::::::")
		fmt.Println(":::::::::::++++:::*:+=++++*=*=::::::::::")
		fmt.Println("::::::::::++*++:+*+::***++****::::::::::")
		fmt.Println(":::::::::::+*+++::*==#*+*****:::::::::::")
		fmt.Println(":::::::::::::++++++***==*++*##+:::::::::")
		fmt.Println(":::::::::::::+++:::::+++++=#@#=+::::::::")
		fmt.Println("::::::::::::+*#*+:::+++**=@@###==++++:::")
		fmt.Println(":::::::::::*#@@*=****==#W@@##==###=*++++")
		fmt.Println(":::::::::+#@@@@=**==*#@@@@##=###@##====*")
		fmt.Println("::::::::*@@@@##:+**=@@@#@##=###@#######=")
		fmt.Println("::::::+#@@#@W##-.-@W@##@##=###@@######@@")
		fmt.Println(":::::*#@#==#@##@*@@######=###@@#@####@@#")
		fmt.Println("::::=##@===@####@@===###==###@@#@#######")
		fmt.Println(":::+##@====@######=#####=###########=##@")
	case 3:
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@#####=#################################################=**")
		fmt.Println("@@@@@@@@@@@@@@@@@@###@@#@@@@@###############==============#####################")
		fmt.Println("@@@        QUEEN         @@@@@@@######================================#########")
		fmt.Println("@@@@@@@@@@@@@@@@@#@@@@@WW@@@@@W@@@#======================================###==#")
		fmt.Println("@@@@@@@@@@@@@@@@W@@@WW@WW@@@@WWWWW@@==========*===####@@######==================")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@#@=++#@@WWWW@@=========*===#@@@@#@@@@##@##================")
		fmt.Println("@@@@@@@@@@@@@@@WWW@++++:+*=#@WWWW@#=========######@@@@@@@@##@@@@==*============")
		fmt.Println("@@@@@#@@@@@@@@@WW@W=+++::=+*#@WWWW@==========####===####@@@@@@@@=*********======")
		fmt.Println("@@@@@#@@@@@@@#@WWW@=:::::=*+=@WWWW#========#####===***=##@@#@@@@=************==")
		fmt.Println("@@@@@#@@@@@@##@WWWW=++::*##=#@WW@@@@@@W@@=*=###==**+++=##@#@@WW#**************=")
		fmt.Println("@@@@@@@@@@@@@##@@WWW@*+++**=##WW@@@@@@@WWWW@**@=*+**+:#**####@W#****************")
		fmt.Println("@@@@@@@@@@@#####@@W@#:**++*###WW@@@#@@@WWWWW@*==::::::=*++=##@@*****************")
		fmt.Println("@@@@@@@@@@#*:---*@=:-::+=#####=W#***=****#@WW=***+::+*#=+*#@@=******************")
		fmt.Println("#@@@@@@+------------+++*=###=*@=++:::++*@@WW#****++**====#@W@#*+**+++*+++++++++")
		fmt.Println("##@@@@@+------------=#++**=#+:*=+++:+*#@@@@@=@@@@=*+:+*###@W#@@*+******+***++++")
		fmt.Println("##@@@@W#-----------:@WW@*:+*=-:*===**#====@@#@@#@@@*=====#@#=W@==#@@@@@###***++")
		fmt.Println("#@@@@W@@+--------*@WWWWW=+*=W#-+++++==++*#@@@@@@WW@#++#WW@+#W@**::=@W@@#====#*+")
		fmt.Println("@@@@WWW@=--------*@WWWW@#===@WW@**+*#@**=##@WWW@W@#@=+*##*@WW#***=@WW#==**###@=")
		fmt.Println("@@@WWWW@@--------*@@WWWWW**#WWW#-*@###@@###@@WWWW@#@=*@@@@@@W#@=#@@@=*=##=###@@")
		fmt.Println("@WWWWWW@@--------*@@WWWW@=*=@W@@+***++*=@##@@@W@=@@@*@@@@@@#@##==**+*==##==##@@")
		fmt.Println("@@WWWWWW#--::----=W@@WWWW#**+*W@+++=###==##@@W@=:::+*#@@@#@#@#=*+::::+**==##@@#")
		fmt.Println("@@@WWWWW+------:-#WWWWW#*++++=@@*++++####@W@#W@++::::+*#=@@@@W#@=#@@@=*=##=###@@")
		fmt.Println("@@@#**=@@:-------@WWW#+++++++=W@=+++++*##=++@W@+:::::+*=@@@@W#@@*+******+***+++")
		fmt.Println("=@#====@@*-------#WW#++::++++@WWWW#=====#@WWWW@=++:++++=#WWWW@@@#*+++*=#=*==#@#")
		fmt.Println("#@#==#@@@*-------=WW#++++++==@WWWWWWWWWWWWWWWWW@=+++:++=#WWWW@*:*=++***=====#@#")
		fmt.Println("@#=*=@@@@*-------#WW#++++++==@WWWWWWWWWWWWWWWWWWW+:::++*=@WW*::::::**+++===#@=#")
		fmt.Println("=@#=***=@#.-----#WWW@+++++*#@WWWWWWWWWWWWWWWWWWWWW+:::++*=@WW*::::::=####==#@===")
		fmt.Println("@@#=***=@@=------=WWW=+++++*=WWWWWWWWWWWWWWWWWWWWW+:::++*=@WW*::::::+@WW@W@#=#**")
		fmt.Println("=@#=***=@@=------=WWW=+++++*=WWWWWWWWWWWWWWWWWWWWW+:::++*=@WW*::::::+@WW@W@#=#**")
		fmt.Println("=@#=***=@@=------=WWW=+++++*=WWWWWWWWWWWWWWWWWWWWW+:::++*=@WW*::::::+@WW@W@#=#**")

	case 0:
		return
	}
}
