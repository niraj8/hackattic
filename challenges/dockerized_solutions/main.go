package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"niraj8/hackattic/challenges/helper"
	"os"
)

const DOCKER_REGISTRY_URL = "a328-2409-40c0-228-ebd1-35aa-fcee-d263-e9b5.ngrok-free.app"

func main() {
	var problem struct {
		Credentials struct {
			User     string `json:"user"`
			Password string `json:"password"`
		} `json:"credentials"`
		IgnitionKey  string `json:"ignition_key"`
		TriggerToken string `json:"trigger_token"`
	}

	helper.GetChallenge("dockerized_solutions", &problem)

	os.WriteFile("IGNITION_KEY.txt", []byte(problem.IgnitionKey), os.ModePerm)

	triggerBody, _ := json.Marshal(map[string]string{
		"registry_host": DOCKER_REGISTRY_URL,
	})

	triggerRes, err := http.Post(helper.HACKATTIC_URL+"/_/push/"+problem.TriggerToken, "application/json", bytes.NewBuffer(triggerBody))

	if err != nil {
		log.Fatalf("failed to trigger push %v", err)
	}
	defer triggerRes.Body.Close()
	body, err := io.ReadAll(triggerRes.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var pushTriggerLogs map[string]interface{}
	json.Unmarshal(body, &pushTriggerLogs)
	fmt.Println(pushTriggerLogs["logs"])
}
