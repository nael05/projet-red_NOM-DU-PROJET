package outils

import (
	"fmt"
	"os"
)
// affiche le menu de depart au lancement du jeu
func Menu() {
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1 : Jouer")
		fmt.Println("2 : Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("Le jeu commence !")
			return 
		case 2:
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
