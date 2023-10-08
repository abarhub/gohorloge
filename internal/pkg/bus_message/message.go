package bus_message

type Heure struct {
	AfficheHeure bool
	AfficheTexte bool
	Fin          bool
	Heure        int
	Minute       int
	Intensite    int
	Texte        string
}

var Messages = make(chan Heure)
