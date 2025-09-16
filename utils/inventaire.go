package main

import "fmt"

type potion struct {
	numero int
	degats int
	prix   int
}

func (p *potion) achat() {
	fmt.Println("")
	fmt.Println("Achète tes potions ici !!")
	fmt.Println("")
	fmt.Println("1 : Potion rouge, augmente de 15 PV 150 Pièces")
	fmt.Println("2 : Potion bleu, inflige 15 PV supplémentaires en cas de victoire de la manche suivante 200 Pièces")
	fmt.Println("3 : Potion poison, empoisonne l'ennemie de 5 PV pendant 3 tours 250 Pièces")
	fmt.Println("4 : Potion vengeance, renvoie les dégâts reçus en cas de défaite de la manche 400 Pièces")
	fmt.Println("5 : Potion voyance, si le joueur devine le signe du joueur adverse, il lui infligera 50 PV 500 Pièces")

	Potion_rouge := potion{1, 15, 150}
	Potion_bleu := potion{2, 15, 200}
	Potion_poison := potion{3, 5, 250}
	Potion_vengeance := potion{4, 0, 400}
	Potion_voyance := potion{5, 50, 500}

	inventaire := []string{}
	solde := 1000
	fmt.Println("")
	fmt.Println("Solde :", solde)

	for solde >= 150 {
		var choix int
		fmt.Print("Choisissez une potion (1-5, 0 pour arrêter) : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("")
			fmt.Println("Vous avez choisi :", "Potion_rouge")
			inventaire = append(inventaire, "Potion rouge")
			solde -= Potion_rouge.prix
			fmt.Println("")
			fmt.Println("Solde restant :", solde)
		case 2:
			fmt.Println("")
			fmt.Println("Vous avez choisi :", "Potion_bleu")
			inventaire = append(inventaire, "Potion bleu")
			solde -= Potion_bleu.prix
			fmt.Println("")
			fmt.Println("Solde restant :", solde)
		case 3:
			fmt.Println("")
			fmt.Println("Vous avez choisi :", "Potion_poison")
			inventaire = append(inventaire, "Potion poison")
			solde -= Potion_poison.prix
			fmt.Println("")
			fmt.Println("Solde restant :", solde)
		case 4:
			fmt.Println("")
			fmt.Println("Vous avez choisi :", "Potion_vengeance")
			inventaire = append(inventaire, "Potion vengeance")
			solde -= Potion_vengeance.prix
			fmt.Println("")
			fmt.Println("Solde restant :", solde)
		case 5:
			fmt.Println("")
			fmt.Println("Vous avez choisi :", "Potion_voyance")
			inventaire = append(inventaire, "Potion voyance")
			solde -= Potion_voyance.prix
			fmt.Println("")
			fmt.Println("Solde restant :", solde)
		case 0:
			fmt.Println("")
			fmt.Println("Achat terminé.")
			solde = 0
			fmt.Println("")
		default:
			fmt.Println("")
			fmt.Println("Choix invalide !")
			fmt.Println("")
		}
	}

	fmt.Println("Inventaire final :", inventaire)
}

func main() {
	var p potion
	p.achat()
}
