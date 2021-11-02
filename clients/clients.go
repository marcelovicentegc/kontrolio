package clients

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	c "github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func getClient(config *oauth2.Config) *http.Client {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	utils.OpenBrowser(authURL)

	runLocalOAuthServer()

	var authCode string

	// Failing here because it expects the user to input the token from the redirected URL
	// Read https://www.digitalocean.com/community/tutorials/an-introduction-to-oauth-2
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return config.Client(context.Background(), token)
}

func UploadFileToDrive() {
	ctx := context.Background()
	credentialsPath := c.GetGoogleCredentialsPath()

	if _, err := os.Stat(credentialsPath); os.IsNotExist(err) {
		file, err := os.Create(credentialsPath)
		defer file.Close()

		if err != nil {
			log.Fatalln("Failed to create credentials.json", err)
		}
	}

	credentials, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	configFromJson, err := google.ConfigFromJSON(credentials, config.KontrolioGoogleScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(configFromJson)

	service, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	dataStorePath := config.GetLocalDataStorePath()

	dataStoreFile, err := os.OpenFile(dataStorePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("failed to open or create "+dataStorePath, err)
	}
	defer dataStoreFile.Close()

	media := drive.File{
		Name:     "Kontrolio punch record",
		MimeType: "text/csv",
	}

	file, err := service.Files.Create(&media).Media(dataStoreFile).Do()

	if err != nil {
		log.Fatalf("Could not create "+media.Name+" : %v", err)
	}

	log.Println(file.Name + " successfully uploaded at " + file.CreatedTime)
	return
}
