package gamePlay

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func IsAlpha(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false
		}
	}
	return true
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

func Tuto() {
	ClearConsole()
	fmt.Println(".___________. __    __  .___________.  ______")
	fmt.Println("|           ||  |  |  | |           | /  __  \\")
	fmt.Println("`---|  |----`|  |  |  | `---|  |----`|  |  |  |")
	fmt.Println("    |  |     |  |  |  |     |  |     |  |  |  |")
	fmt.Println("    |  |     |  `--'  |     |  |     |  `--'  |")
	fmt.Println("    |__|      \\______/      |__|      \\______/")

	fmt.Println("\nVous avez rejoint le tutoriel.")
	fmt.Println("\nDans ce jeux vous Pourrez :")
	fmt.Println("\nEffectuer des choix avec 1, 2, 3...")
	fmt.Println("\nAjouter et Retirer à votre inventaire des item, via Le marchand ou le Forgeron")
	fmt.Println("\nVous battre contre des Monstres")
	fmt.Println("\nGagnez des niveau pour debloquer de nouvelles Compétences en battant des monstres!")
	fmt.Println("\nEn battant des monstres vous pourrez avoir 10 Pourcent de chance d'obtenir un item rare !")
	fmt.Println("\n 1. avancer ...")
	var quitter int
	fmt.Scanln(&quitter)
	switch quitter {
	case 1:
		return
	default:
		break
	}
}
