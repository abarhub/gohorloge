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
		} else if !msg.AfficheHeure && !msg.AfficheTexte {
			if err := arret(); err != nil {
				log.Fatalf("failed to halt to tm1637: %v", err)
			}
		} else if msg.AfficheHeure {
			if msg.Heure >= 0 && msg.Heure <= 99 && msg.Minute >= 0 && msg.Minute <= 99 {
				hours := msg.Heure
				minutes := msg.Minute
				intensite := msg.Intensite
				var heure = clock(hours, minutes, true)
				afficheLed(heure, intensite)
			}
		} else if msg.AfficheTexte {
			if len(msg.Texte) > 0 {
				buf := text(msg.Texte)
				afficheLed(buf, 5)
			}
		}
	}

}

// Hex digits from 0 to F.
//var digitToSegment = []byte{
//	0x3f, 0x06, 0x5b, 0x4f, 0x66, 0x6d, 0x7d, 0x07, 0x7f, 0x6f, 0x77, 0x7c, 0x39, 0x5e, 0x79, 0x71,
//}

var segment = []byte{
	0x3F, 0x06, 0x5B, 0x4F, 0x66, 0x6D, 0x7D, 0x07, 0x7F, 0x6F, // digits
	0x77, 0x7C, 0x39, 0x5E, 0x79, 0x71, 0x3D, 0x76, 0x06, 0x1E, 0x76, 0x38, 0x55, 0x54, 0x3F, 0x73, 0x67, // letters
	0x50, 0x6D, 0x78, 0x3E, 0x1C, 0x2A,
	0x76, // space
	0x6E, 0x5B, 0x00,
	0x40, // dash
	0x63} // star/degrees

func clock(hour, minute int, showDots bool) []byte {
	heure := hour / 10
	heure2 := hour % 10
	minute2 := minute / 10
	minute3 := minute % 10
	seg := make([]byte, 4)
	if heure > 0 {
		seg[0] = segment[heure]
	}
	seg[1] = segment[heure2]
	seg[2] = segment[minute2]
	seg[3] = segment[minute3]
	if showDots {
		seg[1] |= 0x80
	}
	return seg[:]
}

func text(s string) []byte {
	space := segment[36]
	res := make([]byte, 4)
	for i := 0; i < 4; i++ {
		res[i] = space
	}
	for i, c := range s {
		if i >= 4 {
			break
		}
		if c == 32 { // space
			res[i] = space
		} else if c == 42 { // star/degrees
			res[i] = segment[38]
		} else if c == 45 { // dash
			res[i] = segment[37]
		} else if c >= 65 && c <= 90 { // uppercase A-Z
			res[i] = segment[c-55]
		} else if c >= 97 && c <= 122 { // lowercase a-z
			res[i] = segment[c-87]
		} else if c >= 48 && c <= 57 { // 0-9
			res[i] = segment[c-48]
		} else {
			// on affiche espace
			res[i] = space
		}
	}
	return res
}

func convertie(buf []byte) string {
	res := ""
	for _, c := range buf {
		dot := false
		if c&0x80 != 0 {
			dot = true
			c &^= 0x80
		}
		if c == 0 {
			res += " "
		} else {
			trouve := false
			for i, c2 := range segment {
				if c2 == c {
					trouve = true
					if i >= 0 && i <= 9 {
						res += strconv.Itoa(i)
					} else if i >= 10 && i <= 35 {
						res += string(rune(int('a') + (i - 10)))
					} else if i == 36 {
						res += " "
					} else if i == 37 {
						res += "#"
					} else if i == 38 {
						res += "°"
					} else {
						res += "?"
					}
					break
				}
			}
			if !trouve {
				res += "?"
			}
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
