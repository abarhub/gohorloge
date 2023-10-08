package bus_message

const AFFICHE_HEURE = "HEURE"
const AFFICHE_RIEN = "RIEN"
const AFFICHE_MINUTEUR = "MINUTEUR"
const AFFICHE_TEXTE = "AFFICHE_TEXTE"

type Action struct {
	Action    string
	Heure     int
	Minute    int
	Intensite int
	Texte     string
}

var ActionMessage = make(chan Action)
