package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GetChallenge(problemSlug string, problem interface{}) error {
	hackatticAccessToken := os.Getenv("HACKATTIC_ACCESS_TOKEN")
	resp, err := http.Get("https://hackattic.com/challenges/" + problemSlug + "/problem?access_token=" + hackatticAccessToken)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&problem)
	if err != nil {
		return err
	}
	return nil
}

// func SubmitChallengeSolution() {}
