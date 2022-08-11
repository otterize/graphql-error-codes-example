package resolvers

import (
	"fmt"
	"os"
	"server/pkg/model"
	"time"
)

var dogsDB = make(map[string]model.DogInput)
var readPassword = "ilovedogs"

func calculateAge(birthday string) (int, error) {
	birthdayTime, err := time.Parse(time.RFC3339, birthday)
	if err != nil {
		return 0, err
	}
	if birthdayTime.After(time.Now()) {
		timezoneName, _ := time.Now().Zone()
		ntpConf, err := os.ReadFile("/etc/ntp.conf")
		if err != nil {
			return 0, err
		}
		// this is a sensitive error message, it should not be exposed to clients
		return 0, fmt.Errorf("error calculating age. server time: %s server timezone: %s server ntp conf: %s", time.Now(), timezoneName, string(ntpConf))
	}
	age := int(time.Now().Sub(birthdayTime).Hours() / 24 / 365)
	return age, nil
}
