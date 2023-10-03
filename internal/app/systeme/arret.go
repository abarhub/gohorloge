package systeme

import (
	"gohorloge/internal/pkg/action_simple"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range c {
			log.Print("Fin signal")
			action_simple.Arret()
			time.Sleep(10 * time.Second)
			os.Exit(0)
		}
	}()

}
