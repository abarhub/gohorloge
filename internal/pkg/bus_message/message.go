package bus_message

type Heure struct {
	AfficheHeure bool
	Fin          bool
	Heure        int
	Minute       int
	Intensite    int
}

var Messages = make(chan Heure)
