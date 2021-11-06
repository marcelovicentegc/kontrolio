package clients

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/marcelovicentegc/kontrolio-cli/messages"
)

func runLocalOAuthServer(oauthEndpoint string) string {
	mux := http.NewServeMux()
	server := http.Server{Addr: ":8080", Handler: mux}
	authCodeChannel := make(chan string)

	mux.HandleFunc(oauthEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			writer.WriteHeader(http.StatusBadRequest)
		}

		writer.Write([]byte(messages.OAuthAuthenticationFlow))

		authCodeChannel <- request.URL.Query().Get("code")
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	select {
	case code := <-authCodeChannel:
		server.Shutdown(context.Background())
		close(authCodeChannel)

		return code
	}
}
