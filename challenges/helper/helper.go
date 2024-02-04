package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type SolutionResponse map[string]interface{}

func SubmitChallengeSolution(problemSlug string, solution interface{}) (*SolutionResponse, error) {
	hackatticAccessToken := os.Getenv("HACKATTIC_ACCESS_TOKEN")

	responseJSON, err := json.Marshal(&solution)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal solution to json: %v", err)
	}
	log.Println("response json" + string(responseJSON))
	responseBody := bytes.NewBuffer(responseJSON)

	resp, err := http.Post(
		"https://hackattic.com/challenges/"+problemSlug+"/solve?access_token="+hackatticAccessToken,
		"application/json", responseBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var solutionResponse SolutionResponse

	err = json.NewDecoder(resp.Body).Decode(&solutionResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal solution response: %v", err)
	}

	return &solutionResponse, nil
}
