package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
	"gopkg.in/yaml.v2"
)

type FirstRoundAnswers struct {
	Name           string
	HasWorkingTime string `survey:"hasWorkingTime"`
	WorkingDays    []string
}

type SecondRoundAnswers struct {
	WorkingTime string
}

var qsFirstRound = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "What is your name?", Help: "I'll use this information to salute you"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "workingDays",
		Prompt: &survey.MultiSelect{
			Message: "On what days do you usually work?",
			Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			Help:    "I'll use this information to generate meaningful insights to you",
		},
	},
	{
		Name: "hasWorkingTime",
		Prompt: &survey.Select{
			Message: "Do you have a fixed/regular working time?",
			Options: []string{"yes", "no"},
			Default: "yes",
		},
	},
}

var qsSecondRound = []*survey.Question{
	{
		Name:   "workingTime",
		Prompt: &survey.Input{Message: "How many hours a week do you work?", Help: "I'll let you know when you need to compensate"},
		Validate: func(val interface{}) error {

			if hours, err := strconv.Atoi(val.(string)); err == nil && hours > 0 {
				return nil
			}

			return errors.New("only positive real numbers are accepted.")
		},
	},
}

func configure() {
	if config.Network.Status == config.Offline {
		fmt.Println(messages.IsOffline)

		firstRoundAnswers := FirstRoundAnswers{}
		secondRoundAnswers := SecondRoundAnswers{}

		err := survey.Ask(qsFirstRound, &firstRoundAnswers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if firstRoundAnswers.HasWorkingTime == "yes" {

			err := survey.Ask(qsSecondRound, &secondRoundAnswers)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		saveConfiguration(firstRoundAnswers, secondRoundAnswers)

		fmt.Println(messages.ColorWhiteBold(messages.DoneConfiguring) + messages.ColorGreenBold(messages.KontrolioConfigCommand))

		return
	}

	if config.Network.Status == config.Online {
		// TODO: Sync data online
	}
}

func saveConfiguration(firstRoundAnswers FirstRoundAnswers, secondRoundAnswers SecondRoundAnswers) {
	filePath := config.GetConfigFilePath()

	data := fmt.Sprintf(`
name: %s
working_days: %s
working_time: %s	
`,
		firstRoundAnswers.Name,
		firstRoundAnswers.WorkingDays,
		strings.Replace(secondRoundAnswers.WorkingTime, " ", ", ", -1),
	)

	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			file := utils.CreateFile(filePath)

			defer utils.CloseFile(file)

			utils.WriteFile(file, data)

			return
		}

		log.Fatal(err)
	}

	currentConfig := map[string]interface{}{}
	updatedConfig := map[string]interface{}{}

	if err = yaml.Unmarshal(buffer, currentConfig); err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(data), &updatedConfig); err != nil {
		log.Fatal(err)
	}

	for key, value := range updatedConfig {
		currentConfig[key] = value
	}

	updatedData, err := yaml.Marshal(currentConfig)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer utils.CloseFile(file)

	utils.WriteFile(file, string(updatedData))
}
