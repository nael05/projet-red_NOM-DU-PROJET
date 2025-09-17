package outils

import (
	"fmt"
)

type Perso struct {
	Nom        string
	Niveau     int
	PvMax      int
	Pv         int
	Inventaire []string
	Pieces     int
}

func (p *Perso) InitCharacter() {
	fmt.Print("Entre ton nom : ")
	fmt.Scanln(&p.Nom)
	p.Niveau = 1
	p.PvMax = 100
	p.Pv = 100
	p.Pieces = 1000
}

func Printcharactere(joueur Perso) {
	fmt.Println("\nBienvenue,", joueur.Nom, "!")
	fmt.Println("\nPV restant :", joueur.Pv, "        Solde :", joueur.Pieces)
}
