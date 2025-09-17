package outils

import (
	"fmt"
)

type Perso struct {
	nom        string
	niveau     int
	pvmax      int
	pv         int
	inventaire []string
	pieces     int
}

func (p *Perso) InitCharacter() {
	fmt.Print("Entre ton nom : ")
	fmt.Scanln(&p.nom)
	p.niveau = 1
	p.pvmax = 100
	p.pv = 100
	p.pieces = 1000
}

func Printcharactere(joueur Perso) {
	fmt.Println("\nBienvenue,", joueur.nom, "!")
	fmt.Println("\nPV restant :", joueur.pv, "        Solde :", joueur.pieces)
}
