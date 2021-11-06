package clients

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func getDriveClient(ctx context.Context, oauthConfig *oauth2.Config) *http.Client {
	authURL := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	utils.OpenBrowser(authURL)

	authCode := runLocalOAuthServer(config.GoogleDriveOAuthCallback)

	token, err := oauthConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("unable to retrieve token from web: %v", err)
	}

	return oauthConfig.Client(context.Background(), token)
}

func UploadFileToDrive() {
	ctx := context.Background()
	credentialsPath := config.GetGoogleCredentialsPath()

	if _, err := os.Stat(credentialsPath); os.IsNotExist(err) {
		file, err := os.Create(credentialsPath)
		defer file.Close()

		if err != nil {
			log.Fatalln("failed to create credentials.json", err)
		}
	}

	credentials, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		log.Fatalf("unable to read client secret file: %v", err)
	}

	configFromJson, err := google.ConfigFromJSON(credentials, config.KontrolioGoogleScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getDriveClient(ctx, configFromJson)

	service, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("unable to retrieve Drive client: %v", err)
	}

	dataStorePath := config.GetLocalDataStorePath()

	dataStoreFile, err := os.OpenFile(dataStorePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("failed to open or create "+dataStorePath, err)
	}
	defer dataStoreFile.Close()

	media := drive.File{
		Name:     config.KontrolioFilename,
		MimeType: config.KontrolioFileMimeType,
	}

	file, err := service.Files.Create(&media).Media(dataStoreFile).Do()

	if err != nil {
		log.Fatalf("could not create "+media.Name+" : %v", err)
	}

	log.Println(messages.FormatSuccessfullUploadMessage(file.Name))
	os.Exit(0)
}
