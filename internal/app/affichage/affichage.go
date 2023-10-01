package affichage

import (
	"gohorloge/internal/pkg/bus_message"
	"log"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/devices/v3/tm1637"
	"periph.io/x/host/v3"
)

const INTENSITE = tm1637.Brightness4

func initAfficheur() *tm1637.Dev {
	log.Print("initialisation de l'affichage ...")
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	clk := gpioreg.ByName("GPIO5")
	data := gpioreg.ByName("GPIO4")
	if clk == nil || data == nil {
		log.Fatal("Failed to find pins")
	}
	dev, err := tm1637.New(clk, data)
	if err != nil {
		log.Fatalf("failed to initialize tm1637: %v", err)
	}
	if err := dev.SetBrightness(INTENSITE); err != nil {
		log.Fatalf("failed to set brightness on tm1637: %v", err)
	}
	log.Print("initialisation de l'affichage ok")
	return dev
}

func affiche() {

	dev := initAfficheur()

	log.Print("démarrage de la boucle d'affichage ...")

	for {
		msg := <-bus_message.Messages
		if msg.Fin {
			break
		} else if !msg.AfficheHeure {
			if err := dev.Halt(); err != nil {
				log.Fatalf("failed to halt to tm1637: %v", err)
			}
		} else if msg.AfficheHeure && msg.Heure >= 0 && msg.Heure <= 99 && msg.Minute >= 0 && msg.Minute <= 99 {
			hours := msg.Heure
			minutes := msg.Minute
			intensite := msg.Intensite
			intensite2 := INTENSITE
			if intensite > 0 {
				if intensite == 1 {
					intensite2 = tm1637.Brightness1
				} else if intensite == 2 {
					intensite2 = tm1637.Brightness2
				} else if intensite == 3 {
					intensite2 = tm1637.Brightness4
				} else if intensite == 4 {
					intensite2 = tm1637.Brightness10
				} else if intensite == 5 {
					intensite2 = tm1637.Brightness11
				} else if intensite == 6 {
					intensite2 = tm1637.Brightness12
				} else if intensite == 7 {
					intensite2 = tm1637.Brightness13
				} else if intensite == 8 {
					intensite2 = tm1637.Brightness14
				}
			}
			log.Print("intensite:", intensite2, intensite)
			if err := dev.SetBrightness(intensite2); err != nil {
				log.Fatalf("failed to write to tm1637: %v", err)
			}
			//var heure=tm1637.Clock(hours, minutes, true)
			var heure = clock(hours, minutes, true)
			if _, err := dev.Write(heure); err != nil {
				log.Fatalf("failed to write to tm1637: %v", err)
			}
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

func Init() {

	go func() { affiche() }()

}
