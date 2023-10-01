package main

import (
	"gohorloge/internal/app/affichage"
	"gohorloge/internal/app/boucle_evenement"
	"gohorloge/internal/app/systeme"
	"gohorloge/internal/app/web"
	"gohorloge/internal/pkg/action_simple"
	"log"
)

func main() {

	// initialise la boucle d'affichage
	affichage.Init()

	// initialise la boucle d'evenement
	boucle_evenement.Init()

	// affiche l'horloge
	action_simple.Horloge(0)

	// initialisation de l'arret du programme
	systeme.Init()

	// démarrage du serveur web
	err := web.InitWeb()
	if err != nil {
		log.Fatal(err)
	}
}
