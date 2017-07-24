package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func parseArgs() (string, int) {
	app := kingpin.New("redirector", "Minimal HTTP redirector")
	to := app.Arg("to", "URL to redirect to").Required().String()
	port := app.Flag("port", "Listen port").Short('p').Default("80").Int()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	return *to, *port
}

func main() {
	redirectTo, port := parseArgs()

	handle := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
	}
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
