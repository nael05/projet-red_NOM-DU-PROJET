package outils

import (
	"fmt"
	"math/rand"
)

func Boutique(j *Perso) {
	solde := j.Pieces

	for solde >= 150 {
		ClearScreen()

		fmt.Println("\nAchète tes potions ici :")
		fmt.Println("\n1 : Potion rouge, augmente de 15 PV (150 Pièces)")
		fmt.Println("2 : Potion bleu, inflige +15 PV dégâts à la manche suivante (200 Pièces)")
		fmt.Println("3 : Potion poison, empoisonne l'ennemi de 5 PV pendant 3 tours (250 Pièces)")
		fmt.Println("4 : Potion vengeance, renvoie les dégâts reçus (400 Pièces)")
		fmt.Println("5 : Potion voyance, inflige 50 PV si le joueur devine le signe adverse (500 Pièces)")
		fmt.Println("\nSolde :", solde)
		fmt.Println("Inventaire :", j.Inventaire)

		var choix int
		fmt.Print("\nChoisissez une potion (1-5, 0 pour arrêter) : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			if solde >= 150 {
				j.Inventaire = append(j.Inventaire, "Potion rouge")
				solde -= 150
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 2:
			if solde >= 200 {
				j.Inventaire = append(j.Inventaire, "Potion bleu")
				solde -= 200
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 3:
			if solde >= 250 {
				j.Inventaire = append(j.Inventaire, "Potion poison")
				solde -= 250
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 4:
			if solde >= 400 {
				j.Inventaire = append(j.Inventaire, "Potion vengeance")
				solde -= 400
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 5:
			if solde >= 500 {
				j.Inventaire = append(j.Inventaire, "Potion voyance")
				solde -= 500
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 0:
			fmt.Println("\nAchat terminé.")
			j.Pieces = solde
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}

	j.Pieces = solde
	fmt.Println("\nInventaire final :", j.Inventaire)
	fmt.Println("Solde final :", j.Pieces)
}

func GenererInventaireOrdi() Perso {
	ordi := Perso{
		Nom:        "Ordinateur",
		Niveau:     1,
		Pv:         100,
		Pieces:     1000,
		Inventaire: []string{},
	}

	for ordi.Pieces >= 150 {
		choix := rand.Intn(5) + 1

		switch choix {
		case 1:
			if ordi.Pieces >= 150 {
				ordi.Inventaire = append(ordi.Inventaire, "Potion rouge")
				ordi.Pieces -= 150
			}
		case 2:
			if ordi.Pieces >= 200 {
				ordi.Inventaire = append(ordi.Inventaire, "Potion bleu")
				ordi.Pieces -= 200
			}
		case 3:
			if ordi.Pieces >= 250 {
				ordi.Inventaire = append(ordi.Inventaire, "Potion poison")
				ordi.Pieces -= 250
			}
		case 4:
			if ordi.Pieces >= 400 {
				ordi.Inventaire = append(ordi.Inventaire, "Potion vengeance")
				ordi.Pieces -= 400
			}
		case 5:
			if ordi.Pieces >= 500 {
				ordi.Inventaire = append(ordi.Inventaire, "Potion voyance")
				ordi.Pieces -= 500
			}
		}
	}

	return ordi
}
