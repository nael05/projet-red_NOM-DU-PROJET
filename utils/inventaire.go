package main

import "fmt"

type potion struct {
	numero int
	degats int
	prix   int
}

func (p *potion) achat() {
	fmt.Println("Achète tes potions ici !!")
	fmt.Println("")
	fmt.Println("1 : Potion rouge, augmente de  15 PV 150 Pièces")
	fmt.Println("2 : Potion bleu, inflige 15 PV supplémentaires en cas de victoire de la manche suivante200    200 Pièces")
	fmt.Println("3 : Potion poison, empoisonne l'ennemie de 5 PV pendant 3 tours    250 Pièces")
	fmt.Println("4 : Potion vengeance, renvoie les dégats reçus en cas de défaite de la manche    400 Pièces")
	fmt.Println("5 : Potion voyance, si le joueur devine le signe du joueur adverse, il lui infligera 50 PV    500 Pièces")
	
	Potion_rouge := potion{1, 150, 15}
	Potion_bleu := potion{2, 200, 15}
	Potion_poison := potion{3, 250, 5}
	Potion_vengeance := potion{4, 400, 0}
	Potion_voyance := potion{5, 500, 50}

	var choix int
	fmt.Print("Choisissez une potion (1-5) : ")
	fmt.Scanln(&choix)

	inventaire := []string()

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi :", Potion_rouge)
		inventaire = inventaire + "Potion rouge"
	case 2:
		fmt.Println("Vous avez choisi :", Potion_bleu)
		inventaire = inventaire + "Potion bleu"
	case 3:
		fmt.Println("Vous avez choisi :", Potion_poison)
		inventaire = inventaire + "Potion poison"
	case 4:
		fmt.Println("Vous avez choisi :", Potion_vengeance)
		inventaire = inventaire + "Potion vengeance"
	case 5:
		fmt.Println("Vous avez choisi :", Potion_voyance)
		inventaire = inventaire + "Potion voyance"
	default:
		fmt.Println("Choix invalide !")
	}
}

	

}
