package bus_message

const AFFICHE_HEURE = "HEURE"
const AFFICHE_RIEN = "RIEN"
const AFFICHE_MINUTEUR = "MINUTEUR"

type Action struct {
	Action    string
	Heure     int
	Minute    int
	Intensite int
}

var ActionMessage = make(chan Action)
