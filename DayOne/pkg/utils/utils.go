package utils

import (
	"encoding/json"
	"os"

	models "github.com/jpramirez/HundredDaysOfCode/DayOne/pkg/models"
)

//LoadConfiguration returns the read Configuration and error while reading.
func LoadConfiguration(file string) (models.Config, error) {
	var config models.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
