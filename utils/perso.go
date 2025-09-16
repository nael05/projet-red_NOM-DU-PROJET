package main

import (
	"fmt"
)

type perso struct {
	nom        string
	niveau     int
	pvmax      int
	pv         int
	inventaire []int
}

func (p *perso) initCharacter() {
	fmt.Print("Entre ton nom : ")
	fmt.Scanln(&p.nom)
	p.niveau = 1
	p.pvmax = 100
	p.pv = 100
}

func main() {
	var joueur perso
	joueur.initCharacter()
	fmt.Println("Bienvenue,", joueur.nom, "!")
	fmt.Println("PV restant :", joueur.pv, "        Solde :", joueur.pieces)
	fmt.Println(" ")
	fmt.Println("MENU:")
	fmt.Println("1 : Jouer")
	fmt.Println("2 : Sortir")
}
