package affichage

import (
	"log"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/devices/v3/tm1637"
	"periph.io/x/host/v3"
)

const INTENSITE = tm1637.Brightness4

var dev *tm1637.Dev

func initAfficheur() {
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
	dev2, err := tm1637.New(clk, data)
	if err != nil {
		log.Fatalf("failed to initialize tm1637: %v", err)
	}
	dev = dev2
	if err := dev.SetBrightness(INTENSITE); err != nil {
		log.Fatalf("failed to set brightness on tm1637: %v", err)
	}
	log.Print("initialisation de l'affichage ok")
	//return dev
}

func arret() error {
	return dev.Halt()
}

func afficheLed(heure []byte, intensite int) {
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
	if _, err := dev.Write(heure); err != nil {
		log.Fatalf("failed to write to tm1637: %v", err)
	}
}
