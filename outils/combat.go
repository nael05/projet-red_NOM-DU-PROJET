package outils

import (
	"fmt"
	"math/rand"
)

func Combat(j *Perso, ordi *Perso) string {
	manche := 1
	var potionUtilisee, potionOrdi string
	poisonToursJoueur, poisonToursOrdi := 0, 0
	poisonDegat := 5
	voyanceActive := false
	var prediction string
	attaques := []string{"🪨", "🍃", "✂️"}

	for j.Pv > 0 && ordi.Pv > 0 {
		ClearScreen()
		fmt.Println("========================================")
		fmt.Println(j.Nom, ":", j.Pv, "PV\tOrdinateur :", ordi.Pv, "PV")
		fmt.Println("---------- Manche", manche, "----------")
		manche++

		if voyanceActive {
			fmt.Println("Voyance activée ! Devine l'attaque de l'ennemi (1-🪨, 2-🍃, 3-✂️) :")
			for i, a := range attaques {
				fmt.Printf("%d : %s\t", i+1, a)
			}
			fmt.Print("\nTon choix : ")
			var c int
			fmt.Scanln(&c)
			if c >= 1 && c <= 3 {
				prediction = attaques[c-1]
			} else {
				prediction = attaques[rand.Intn(3)]
			}
			voyanceActive = false
		}

		fmt.Println("\nOptions :")
		fmt.Println("1 : 🪨\t2 : 🍃\t3 : ✂️\t4 : Inventaire")
		fmt.Print("Ton choix : ")
		var choix int
		fmt.Scanln(&choix)

		potionUtilisee = ""
		vengeanceActive := false

		if choix == 4 {
			if len(j.Inventaire) == 0 {
				fmt.Println("Ton inventaire est vide ! Choisis une attaque (1-3).")
				continue
			}
			fmt.Println("\n------ INVENTAIRE ------")
			for i, p := range j.Inventaire {
				fmt.Printf("%d : %s\n", i+1, p)
			}
			fmt.Print("Numéro de potion à utiliser : ")
			var pChoix int
			fmt.Scanln(&pChoix)
			if pChoix < 1 || pChoix > len(j.Inventaire) {
				fmt.Println("Choix invalide !")
				continue
			}

			potionUtilisee = j.Inventaire[pChoix-1]
			j.Inventaire = append(j.Inventaire[:pChoix-1], j.Inventaire[pChoix:]...)

			switch potionUtilisee {
			case "Potion rouge":
				j.Pv += 15
				if j.Pv > j.PvMax {
					j.Pv = j.PvMax
				}
				fmt.Println("Tu gagnes 15 PV !")
			case "Potion bleu":
				fmt.Println("Potion bleu activée : +15 dégâts au prochain tour si tu gagnes")
			case "Potion poison":
				poisonToursOrdi = 3
				fmt.Println("Potion poison activée : -5 PV à l'ennemi pendant 3 tours")
			case "Potion vengeance":
				vengeanceActive = true
				fmt.Println("Potion vengeance activée : les dégâts reçus seront renvoyés ce tour uniquement")
			case "Potion voyance":
				voyanceActive = true
				fmt.Println("Potion voyance activée ! Tu pourras deviner l'attaque de l'ennemi au prochain tour.")
			}

			fmt.Print("Choisis ton attaque (1-3) : ")
			fmt.Scanln(&choix)
			if choix < 1 || choix > 3 {
				choix = 1
			}
		}

		attaqueJoueur := attaques[choix-1]

		choixOrdi := rand.Intn(4)
		potionOrdi = ""
		attaqueOrdi := attaques[choixOrdi%3]
		vengeanceOrdi := false

		if choixOrdi == 3 && len(ordi.Inventaire) > 0 {
			pIndex := rand.Intn(len(ordi.Inventaire))
			potionOrdi = ordi.Inventaire[pIndex]
			ordi.Inventaire = append(ordi.Inventaire[:pIndex], ordi.Inventaire[pIndex+1:]...)

			switch potionOrdi {
			case "Potion rouge":
				ordi.Pv += 15
				if ordi.Pv > ordi.PvMax {
					ordi.Pv = ordi.PvMax
				}
			case "Potion bleu":
				fmt.Println("Ordinateur active Potion bleu")
			case "Potion poison":
				poisonToursJoueur = 3
				fmt.Println("Ordinateur active Potion poison : -5 PV à toi pendant 3 tours")
			case "Potion vengeance":
				vengeanceOrdi = true
				fmt.Println("Ordinateur active Potion vengeance")
			case "Potion voyance":
				fmt.Println("Ordinateur active Potion voyance")
			}

			choixOrdi = rand.Intn(3)
			attaqueOrdi = attaques[choixOrdi]
		}

		degatJoueur, degatOrdi := 0, 0

		if attaqueJoueur == "🪨" && attaqueOrdi == "✂️" {
			degatJoueur = 12
		} else if attaqueJoueur == "✂️" && attaqueOrdi == "🍃" {
			degatJoueur = 11
		} else if attaqueJoueur == "🍃" && attaqueOrdi == "🪨" {
			degatJoueur = 10
		}

		if attaqueOrdi == "🪨" && attaqueJoueur == "✂️" {
			degatOrdi = 12
		} else if attaqueOrdi == "✂️" && attaqueJoueur == "🍃" {
			degatOrdi = 11
		} else if attaqueOrdi == "🍃" && attaqueJoueur == "🪨" {
			degatOrdi = 10
		}

		if potionUtilisee == "Potion bleu" && degatJoueur > 0 {
			degatJoueur += 15
		}
		if potionOrdi == "Potion bleu" && degatOrdi > 0 {
			degatOrdi += 15
		}

		if poisonToursJoueur > 0 {
			j.Pv -= poisonDegat
			poisonToursJoueur--
		}
		if poisonToursOrdi > 0 {
			ordi.Pv -= poisonDegat
			poisonToursOrdi--
		}

		if prediction != "" && attaqueOrdi == prediction {
			fmt.Println("Voyance réussie ! Tu infliges 50 PV supplémentaires.")
			ordi.Pv -= 50
			prediction = ""
		}

		if degatOrdi > 0 {
			if vengeanceActive {
				fmt.Println("Potion vengeance : tu renvoies", degatOrdi, "PV à l'ordi")
				ordi.Pv -= degatOrdi
			} else {
				j.Pv -= degatOrdi
			}
		}

		if degatJoueur > 0 {
			if vengeanceOrdi {
				fmt.Println("Potion vengeance de l'ordi : il te renvoie", degatJoueur, "PV")
				j.Pv -= degatJoueur
			} else {
				ordi.Pv -= degatJoueur
			}
		}

		fmt.Println("\n------ RÉSULTAT DU TOUR ------")
		fmt.Println(j.Nom, "a utilisé :", attaqueJoueur)
		if potionUtilisee != "" {
			fmt.Println("+ potion :", potionUtilisee)
		}
		if degatJoueur > 0 && !vengeanceActive {
			fmt.Println("-> inflige", degatJoueur, "PV à l'ordinateur")
		}

		fmt.Println("\nOrdinateur a utilisé :", attaqueOrdi)
		if potionOrdi != "" {
			fmt.Println("+ potion :", potionOrdi)
		}
		if degatOrdi > 0 && !vengeanceOrdi {
			fmt.Println("-> inflige", degatOrdi, "PV à", j.Nom)
		}

		fmt.Println("\n========================================")
		fmt.Print("Appuie sur Entrée pour passer à la manche suivante...")
		fmt.Scanln()
	}

	if j.Pv <= 0 {
		return "Ordinateur"
	}
	return j.Nom
}
