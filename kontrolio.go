package main

import (
	"log"
	"net/http"

	"github.com/marcelovicentegc/kontrolio-cli/cli"
)

func main() {
	cli.Kontrolio()

	http.HandleFunc("/drive/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Callback was called")
		log.Println(r.Header)
	})

	http.ListenAndServe(":8080", nil)
}
