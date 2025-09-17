package outils

import (
	"fmt"
	"math/rand"
	"time"
)

func Boutique(j *Perso) {
	solde := j.pieces
	inventaire := []string{}

	for solde >= 150 {
		ClearScreen()

		fmt.Println("\nAchète tes potions ici :")
		fmt.Println("\n1 : Potion rouge, augmente de 15 PV 150 Pièces")
		fmt.Println("2 : Potion bleu, inflige 15 PV supplémentaires en cas de victoire de la manche suivante 200 Pièces")
		fmt.Println("3 : Potion poison, empoisonne l'ennemie de 5 PV pendant 3 tours 250 Pièces")
		fmt.Println("4 : Potion vengeance, renvoie les dégâts reçus en cas de défaite de la manche 400 Pièces")
		fmt.Println("5 : Potion voyance, si le joueur devine le signe du joueur adverse, il lui infligera 50 PV 500 Pièces")
		fmt.Println("\nSolde :", solde)
		fmt.Println("Inventaire :", inventaire)

		var choix int
		fmt.Print("\nChoisissez une potion (1-5, 0 pour arrêter) : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			if solde >= 150 {
				inventaire = append(inventaire, "Potion rouge")
				solde -= 150
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 2:
			if solde >= 200 {
				inventaire = append(inventaire, "Potion bleu")
				solde -= 200
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 3:
			if solde >= 250 {
				inventaire = append(inventaire, "Potion poison")
				solde -= 250
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 4:
			if solde >= 400 {
				inventaire = append(inventaire, "Potion vengeance")
				solde -= 400
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 5:
			if solde >= 500 {
				inventaire = append(inventaire, "Potion voyance")
				solde -= 500
			} else {
				fmt.Println("Solde insuffisant pour acheter cette potion !")
			}
		case 0:
			fmt.Println("\nAchat terminé.")
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}

	j.inventaire = inventaire
	j.pieces = solde
	fmt.Println("\nInventaire final :", j.inventaire)
	fmt.Println("Solde final :", j.pieces)
}

func GenererInventaireOrdi() Perso {
	rand.Seed(time.Now().UnixNano())
	ordi := Perso{
		nom:        "Ordinateur",
		niveau:     1,
		pvmax:      100,
		pv:         100,
		pieces:     1000,
		inventaire: []string{},
	}

	for ordi.pieces >= 150 {
		choix := rand.Intn(5) + 1

		switch choix {
		case 1:
			if ordi.pieces >= 150 {
				ordi.inventaire = append(ordi.inventaire, "Potion rouge")
				ordi.pieces -= 150
			}
		case 2:
			if ordi.pieces >= 200 {
				ordi.inventaire = append(ordi.inventaire, "Potion bleu")
				ordi.pieces -= 200
			}
		case 3:
			if ordi.pieces >= 250 {
				ordi.inventaire = append(ordi.inventaire, "Potion poison")
				ordi.pieces -= 250
			}
		case 4:
			if ordi.pieces >= 400 {
				ordi.inventaire = append(ordi.inventaire, "Potion vengeance")
				ordi.pieces -= 400
			}
		case 5:
			if ordi.pieces >= 500 {
				ordi.inventaire = append(ordi.inventaire, "Potion voyance")
				ordi.pieces -= 500
			}
		}
	}

	return ordi
}
