package web

import (
	"embed"
	"fmt"
	"gohorloge/internal/pkg/action_simple"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func actionHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "horloge") {
		log.Print("horloge")
		intensite := 0
		if r.URL.Query().Has("intensite") {
			intense, err := strconv.Atoi(r.URL.Query().Get("intensite"))
			if err != nil {
				log.Print("erreur", err)
			} else {
				intensite = intense
			}
		}
		action_simple.Horloge(intensite)
	} else if strings.HasSuffix(r.URL.Path, "minuteur") {
		log.Print("minuteur")
		if r.URL.Query().Has("time") {
			time := r.URL.Query().Get("time")
			log.Print("time", time)
			tab := strings.Split(time, ":")
			log.Print("tab:", tab)
			if len(tab) == 3 {
				minutes, err := strconv.Atoi(tab[1])
				if err == nil {
					secondes, err := strconv.Atoi(tab[2])
					if err == nil && minutes >= 0 && secondes >= 0 && !(minutes == 0 && secondes == 0) {
						action_simple.Minuteur(minutes, secondes)
					} else {
						log.Print("erreur", err)
					}
				} else {
					log.Print("erreur", err)
				}
			}
		}
	} else if strings.HasSuffix(r.URL.Path, "arret") {
		log.Print("arret")
		action_simple.Arret()
	}
	fmt.Fprint(w, "Action")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("02 Jan 2006 15:04:05 MST"))
}

//go:embed static
var staticFiles embed.FS

func InitWeb() error {

	var staticFS = fs.FS(staticFiles)
	htmlContent, err := fs.Sub(staticFS, "static")

	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/static/", fs)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/api/action/horloge", actionHandler)
	http.HandleFunc("/api/action/minuteur", actionHandler)
	http.HandleFunc("/api/action/arret", actionHandler)
	//http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/", http.FileServer(http.FS(htmlContent)))

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}
	log.Print("Fin")
	return nil
}
