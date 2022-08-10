package resolvers

import (
	"server/pkg/model"
	"time"
)

var dogsDB = make(map[string]model.DogInput)

func calculateAge(birthday string) (int, error) {
	birthdayTime, err := time.Parse(time.RFC3339, birthday)
	if err != nil {
		return 0, err
	}
	return int(time.Now().Sub(birthdayTime).Hours() / 24 / 365), nil
}
