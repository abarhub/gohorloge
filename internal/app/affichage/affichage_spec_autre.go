//go:build !arm

package affichage

import "log"

func initAfficheur() {
	log.Print("Initialisation affichage")
}

func arret() error {
	log.Print("arret")
	return nil
}

func afficheLed(heure []byte, intensite int) {
	log.Print("affichage led", heure, intensite)
}
