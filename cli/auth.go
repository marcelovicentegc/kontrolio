package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
)

var authQs = []*survey.Question{
	{
		Name:      "email",
		Prompt:    &survey.Input{Message: "Please, type your email"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "password",
		Prompt: &survey.Password{
			Message: "Please, type your password",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
}

func auth() {
	isDevEnvironment := config.IsDevEnvironment()

	if !isDevEnvironment {
		fmt.Println("\n⚠️  sorry, this functionality is a work in progress...")
		return
	}

	if config.Network.Status == config.Offline {
		// TODO: This flow should be on the online condition

		fmt.Println(messages.IsOnline)

		answers := struct {
			Email    string
			Password string
		}{}

		err := survey.Ask(authQs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// TODO: Authenticate user

		return
	}

	if config.Network.Status == config.Online {
		// TODO: Move everything from Offline scope here
	}
}
