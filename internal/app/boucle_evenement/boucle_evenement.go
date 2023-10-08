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
	var texteCompletAAffiche string
	var boucleTexte int
	var affichageTextePrecedant time.Time

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
			afficheOk := false
			if nouveau {
				texteCompletAAffiche = actionSelectionnee.Texte
				boucleTexte = 0
				affichageTextePrecedant = time.Now()
			}
			texte := texteCompletAAffiche
			var texteAAfficher = ""
			if len(texte) == 0 {
				texteAAfficher = " "
				afficheOk = nouveau
			} else if len(texte) <= 4 {
				texteAAfficher = texte
				for len(texteAAfficher) < 4 {
					texteAAfficher += " "
				}
				afficheOk = nouveau
			} else {
				now := time.Now()
				diff := now.Sub(affichageTextePrecedant)

				if nouveau || diff.Milliseconds() > 1000 {
					afficheOk = true
					affichageTextePrecedant = now
					boucleTexte++
					buf := ""
					s := texte + " "
					for i := 0; i < 4; i++ {
						buf += string(s[(i+boucleTexte)%len(s)])
					}
					texteAAfficher = buf
				}
			}
			if len(texte) > 0 && afficheOk {
				bus_message.Messages <- bus_message.Heure{AfficheTexte: true, Texte: texteAAfficher}
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
