package outils

import (
	"fmt"
	"math/rand"
)

func Combat(j *Perso, ordi *Perso) string {
	manche := 1
	attaques := []string{"🪨", "🍃", "✂️"}
	poisonJ, poisonO := 0, 0
	poisonDegat := 5
	voyanceActive := false
	prediction := ""

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
			var c int
			fmt.Scanln(&c)
			if c >= 1 && c <= 3 {
				prediction = attaques[c-1]
			} else {
				prediction = attaques[rand.Intn(3)]
			}
			voyanceActive = false
		}

		fmt.Println("\n1: 🪨  2: 🍃  3: ✂️  4: Inventaire")
		fmt.Print("Ton choix : ")
		var choix int
		fmt.Scanln(&choix)

		potionJ := ""
		vengeanceJ := false

		if choix == 4 {
			if len(j.Inventaire) == 0 {
				fmt.Println("Inventaire vide ! Choisis une attaque.")
				fmt.Print("Ton choix : ")
				fmt.Scanln(&choix)
			} else {
				fmt.Println("\n------ INVENTAIRE ------")
				for i, p := range j.Inventaire {
					fmt.Printf("%d : %s\n", i+1, p)
				}
				fmt.Print("Numéro de potion à utiliser : ")
				var pChoix int
				fmt.Scanln(&pChoix)
				if pChoix >= 1 && pChoix <= len(j.Inventaire) {
					potionJ = j.Inventaire[pChoix-1]
					j.Inventaire = append(j.Inventaire[:pChoix-1], j.Inventaire[pChoix:]...)
					switch potionJ {
					case "Potion rouge":
						j.Pv += 15
						if j.Pv > j.PvMax {
							j.Pv = j.PvMax
						}
						fmt.Println(j.Nom, "a utilisé Potion rouge (+15 PV)")
					case "Potion bleu":
						fmt.Println(j.Nom, "a utilisé Potion bleu (+15 PV dégâts si attaque réussie)")
					case "Potion poison":
						poisonO = 3
						fmt.Println(j.Nom, "a utilisé Potion poison (-5 PV à l'ennemi pendant 3 tours)")
					case "Potion vengeance":
						vengeanceJ = true
						fmt.Println(j.Nom, "a utilisé Potion vengeance (renvoie les dégâts reçus ce tour)")
					case "Potion voyance":
						voyanceActive = true
						fmt.Println(j.Nom, "a utilisé Potion voyance (devine attaque de l'ordi prochain tour)")
					}
				}
				fmt.Print("Choisis ton attaque (1-3) : ")
				fmt.Scanln(&choix)
			}
		}

		if choix < 1 || choix > 3 {
			choix = 1
		}

		attaqueJ := attaques[choix-1]

		choixO := rand.Intn(4)
		potionO := ""
		attaqueO := attaques[choixO%3]
		vengeanceO := false
		if choixO == 3 && len(ordi.Inventaire) > 0 {
			pIndex := rand.Intn(len(ordi.Inventaire))
			potionO = ordi.Inventaire[pIndex]
			ordi.Inventaire = append(ordi.Inventaire[:pIndex], ordi.Inventaire[pIndex+1:]...)
			switch potionO {
			case "Potion rouge":
				ordi.Pv += 15
				if ordi.Pv > ordi.PvMax {
					ordi.Pv = ordi.PvMax
				}
				fmt.Println("Ordinateur a utilisé Potion rouge (+15 PV)")
			case "Potion bleu":
				fmt.Println("Ordinateur a utilisé Potion bleu (+15 PV dégâts si attaque réussie)")
			case "Potion poison":
				poisonJ = 3
				fmt.Println("Ordinateur a utilisé Potion poison (-5 PV à toi pendant 3 tours)")
			case "Potion vengeance":
				vengeanceO = true
				fmt.Println("Ordinateur a utilisé Potion vengeance (renvoie les dégâts reçus ce tour)")
			case "Potion voyance":
				fmt.Println("Ordinateur a utilisé Potion voyance (devinera ton attaque prochain tour)")
			}
			choixO = rand.Intn(3)
			attaqueO = attaques[choixO]
		}

		degJ, degO := 0, 0
		if attaqueJ == "🪨" && attaqueO == "✂️" {
			degJ = 12
		} else if attaqueJ == "✂️" && attaqueO == "🍃" {
			degJ = 11
		} else if attaqueJ == "🍃" && attaqueO == "🪨" {
			degJ = 10
		}
		if attaqueO == "🪨" && attaqueJ == "✂️" {
			degO = 12
		} else if attaqueO == "✂️" && attaqueJ == "🍃" {
			degO = 11
		} else if attaqueO == "🍃" && attaqueJ == "🪨" {
			degO = 10
		}

		if potionJ == "Potion bleu" && degJ > 0 {
			degJ += 15
		}
		if potionO == "Potion bleu" && degO > 0 {
			degO += 15
		}

		if poisonJ > 0 {
			j.Pv -= poisonDegat
			poisonJ--
			fmt.Println(j.Nom, "subit", poisonDegat, "PV de poison")
		}
		if poisonO > 0 {
			ordi.Pv -= poisonDegat
			poisonO--
			fmt.Println("Ordinateur subit", poisonDegat, "PV de poison")
		}

		if prediction != "" && attaqueO == prediction {
			fmt.Println("Voyance réussie ! 50 PV supplémentaires à l'ordi")
			ordi.Pv -= 50
			prediction = ""
		}

		if degO > 0 {
			if vengeanceJ {
				ordi.Pv -= degO
				fmt.Println(j.Nom, "renvoie", degO, "PV à l'ordi avec vengeance")
			} else {
				j.Pv -= degO
			}
		}

		if degJ > 0 {
			if vengeanceO {
				j.Pv -= degJ
				fmt.Println("Ordinateur renvoie", degJ, "PV à", j.Nom, "avec vengeance")
			} else {
				ordi.Pv -= degJ
			}
		}

		fmt.Println("\n------ RÉSULTAT DU TOUR ------")
		fmt.Println(j.Nom, "attaque :", attaqueJ)
		if potionJ != "" {
			fmt.Println("Potion utilisée :", potionJ)
		}
		fmt.Println("Ordinateur attaque :", attaqueO)
		if potionO != "" {
			fmt.Println("Potion utilisée :", potionO)
		}

		if degJ > degO {
			fmt.Println("Gagnant de la manche :", j.Nom)
		} else if degO > degJ {
			fmt.Println("Gagnant de la manche : Ordinateur")
		} else {
			fmt.Println("La manche est nulle")
		}

		fmt.Println("========================================")
		fmt.Print("Appuie sur Entrée pour passer à la manche suivante...")
		fmt.Scanln()
	}

	fmt.Println("\n=== Partie terminée ===")
	if j.Pv <= 0 {
		fmt.Println("Vainqueur :", "Ordinateur")
		return "Ordinateur"
	}
	fmt.Println("Vainqueur :", j.Nom)
	return j.Nom
}
