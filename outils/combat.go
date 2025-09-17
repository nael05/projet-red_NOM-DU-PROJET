package outils

import (
	"fmt"
	"math/rand"
)

func Combat(j *Perso, ordi *Perso) string {
	manche := 1
	for j.pv > 0 && ordi.pv > 0 {
		ClearScreen()
		fmt.Println("\npV", j.nom, "restant :", j.pv, "          pV ordi restant :", ordi.pv)
		fmt.Println("")
		fmt.Println("Manche", manche, ":")
		manche++
		fmt.Println("\n1 : 🪨")
		fmt.Println("2 : 🍃")
		fmt.Println("3 : ✂️")
		fmt.Println("4 : inventaire")
		var choix int
		fmt.Print("\nQuelle attaque veux-tu lancer ? : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("\nTu as lancé 🪨")
		case 2:
			fmt.Println("\nTu as lancé 🍃")
		case 3:
			fmt.Println("\nTu as lancé ✂️")
		case 4:
			fmt.Println("\nInventaire :", j.inventaire)
		default:
			fmt.Println("\nChoix invalide !")
			continue
		}
		choixrobot := rand.Intn(4) + 1
		if choixrobot == 1 {
			fmt.Println("\nTon ennemi a lancé 🪨")
		} else if choixrobot == 2 {
			fmt.Println("\nTon ennemi a lancé 🍃")
		} else if choixrobot == 3 {
			fmt.Println("\nTon ennemi a lancé ✂️")
		} else if choixrobot == 4 && len(ordi.inventaire) > 0 {
			choixrobotpotion := rand.Intn(len(ordi.inventaire))
			fmt.Println("\nTon ennemi a utilisé la potion", ordi.inventaire[choixrobotpotion])
			// Retirer la potion de l’inventaire
			ordi.inventaire = append(ordi.inventaire[:choixrobotpotion], ordi.inventaire[choixrobotpotion+1:]...)
		} else {
			choixrobot2 := rand.Intn(3) + 1
			if choixrobot2 == 1 {
				fmt.Println("\nTon ennemi a lancé 🪨")
			} else if choixrobot2 == 2 {
				fmt.Println("\nTon ennemi a lancé 🍃")
			} else {
				fmt.Println("\nTon ennemi a lancé ✂️")
			}
		}
	}

	if j.pv <= 0 {
		return "ordi"
	}
	return "joueur"
}
