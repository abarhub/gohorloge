package action_simple

import "gohorloge/internal/pkg/bus_message"

func Horloge(intensite int) {
	bus_message.ActionMessage <- bus_message.Action{Action: bus_message.AFFICHE_HEURE, Intensite: intensite}
}

func Minuteur(minute, secondes int) {
	bus_message.ActionMessage <- bus_message.Action{Action: bus_message.AFFICHE_MINUTEUR, Heure: minute, Minute: secondes}
}

func AffichageTexte(texte string) {
	bus_message.ActionMessage <- bus_message.Action{Action: bus_message.AFFICHE_TEXTE, Texte: texte}
}

func Arret() {
	bus_message.ActionMessage <- bus_message.Action{Action: bus_message.AFFICHE_RIEN}
}
