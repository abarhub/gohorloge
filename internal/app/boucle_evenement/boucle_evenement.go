package boucle_evenement

import (
	"gohorloge/internal/pkg/bus_message"
	"log"
	"time"
)

func boucleEvenement() {

	var actionPrecedante bus_message.Action
	var minutes, secondes int
	var begin, end time.Time
	var difference time.Duration

	log.Print("démarrage de la boucle d'evenement ...")

	for {

		nouveau := false
		var actionSelectionnee bus_message.Action

		select {
		case actionCourante := <-bus_message.ActionMessage:
			actionSelectionnee = actionCourante
			nouveau = true
			log.Printf("nouvelle action_simple: %v", actionSelectionnee)
		case <-time.After(1000 * time.Millisecond):
			actionSelectionnee = actionPrecedante
		}

		if actionSelectionnee.Action == bus_message.AFFICHE_HEURE {
			now := time.Now()
			bus_message.Messages <- bus_message.Heure{AfficheHeure: true, Heure: now.Hour(), Minute: now.Minute(), Intensite: actionSelectionnee.Intensite}
		} else if actionSelectionnee.Action == bus_message.AFFICHE_MINUTEUR {
			if nouveau {
				minutes = actionSelectionnee.Heure
				secondes = actionSelectionnee.Minute
				begin = time.Now()
				end = begin.Add(time.Minute*time.Duration(minutes) + time.Second*time.Duration(secondes))
			} else {

			}
			now := time.Now()
			if now.After(end) || now.Equal(end) {
				bus_message.Messages <- bus_message.Heure{AfficheHeure: true, Heure: 0, Minute: 0}
			} else {
				diff := end.Sub(now)
				if nouveau || int(diff.Seconds()) != int(difference.Seconds()) {
					minutes := int(diff.Minutes())
					secondes := int(diff.Seconds()) - int(diff.Minutes())*60
					bus_message.Messages <- bus_message.Heure{AfficheHeure: true, Heure: minutes, Minute: secondes}
					difference = diff
				}
			}
		} else if actionSelectionnee.Action == bus_message.AFFICHE_TEXTE {
			texte := actionSelectionnee.Texte
			if len(texte) > 0 {
				bus_message.Messages <- bus_message.Heure{AfficheTexte: true, Texte: texte}
			}
		} else if actionSelectionnee.Action == bus_message.AFFICHE_RIEN {
			if nouveau {
				bus_message.Messages <- bus_message.Heure{AfficheHeure: false}
			}
		}

		//time.Sleep(500 * time.Millisecond)

		actionPrecedante = actionSelectionnee
	}
}

func Init() {

	go func() { boucleEvenement() }()

}
