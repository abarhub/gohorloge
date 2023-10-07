package affichage

import (
	"gohorloge/internal/pkg/bus_message"
	"log"
	"strconv"
)

func affiche() {

	initAfficheur()

	log.Print("démarrage de la boucle d'affichage ...")

	for {
		msg := <-bus_message.Messages
		if msg.Fin {
			break
		} else if !msg.AfficheHeure {
			if err := arret(); err != nil {
				log.Fatalf("failed to halt to tm1637: %v", err)
			}
		} else if msg.AfficheHeure && msg.Heure >= 0 && msg.Heure <= 99 && msg.Minute >= 0 && msg.Minute <= 99 {
			hours := msg.Heure
			minutes := msg.Minute
			intensite := msg.Intensite
			var heure = clock(hours, minutes, true)
			afficheLed(heure, intensite)
		}
	}

}

// Hex digits from 0 to F.
var digitToSegment = []byte{
	0x3f, 0x06, 0x5b, 0x4f, 0x66, 0x6d, 0x7d, 0x07, 0x7f, 0x6f, 0x77, 0x7c, 0x39, 0x5e, 0x79, 0x71,
}

func clock(hour, minute int, showDots bool) []byte {
	heure := hour / 10
	heure2 := hour % 10
	minute2 := minute / 10
	minute3 := minute % 10
	seg := make([]byte, 4)
	if heure > 0 {
		seg[0] = byte(digitToSegment[heure])
	}
	seg[1] = byte(digitToSegment[heure2])
	seg[2] = byte(digitToSegment[minute2])
	seg[3] = byte(digitToSegment[minute3])
	if showDots {
		seg[1] |= 0x80
	}
	return seg[:]
}

func convertie(buf []byte) string {
	res := ""
	for _, c := range buf {
		dot := false
		if c&0x80 != 0 {
			dot = true
			c &^= 0x80
		}
		trouve := false
		for i, c2 := range digitToSegment {
			if c2 == c {
				trouve = true
				if i >= 0 && i <= 9 {
					res += strconv.Itoa(i)
				} else {
					res += "?"
				}
				break
			}
		}
		if !trouve {
			res += "?"
		}
		if dot {
			res += "."
		}
	}
	return res
}

func Init() {

	go func() { affiche() }()

}
