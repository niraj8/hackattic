package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"niraj8/hackattic/challenges/helper"
)

type qrServerReadAPISuccessResponse struct {
	Type   string `json:"type"`
	Symbol []struct {
		Seq  int    `json:"seq"`
		Data string `json:"data"`
	} `json:"symbol"`
}

func main() {
	var problem struct {
		ImageURL string `json:"image_url"`
	}
	var solution struct {
		Code string `json:"code"`
	}
	helper.GetChallenge("reading_qr", &problem)

	// code, err := apiCall(problem.ImageURL)
	code, err := zbar(problem.ImageURL)
	if err != nil {
		log.Fatal(err)
	}
	solution.Code = *code

	solutionResponse, err := helper.SubmitChallengeSolution("reading_qr", &solution)
	helper.HandleError(err)
	log.Println(solutionResponse)
}

func zbar(imageURL string) (*string, error) {
	// write the file to disk
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the qr image: %v", err)
	}
	defer resp.Body.Close()
	f, err := os.Create("reading_qr.png")
	if err != nil {
		return nil, fmt.Errorf("failed to create file on disk: %v", err)
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to write qr image to disk: %v", err)
	}

	// exec zbarimg and get decoded value and return
	decodeQrCmd := exec.Command("zbarimg", "-q", "--raw", "reading_qr.png")
	output, err := decodeQrCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get output from exec cmd: %v", err)
	}
	code := strings.Trim(string(output), "\n")
	return &code, nil
}

// this doesn't work, TLE
func apiCall(imageURL string) (*string, error) {
	resp, err := http.Get("https://api.qrserver.com/v1/read-qr-code/?fileurl=" + url.QueryEscape(imageURL))
	if err != nil {
		log.Fatalf("request to decode qr code failed: %v", err)
	}
	var responseBody []qrServerReadAPISuccessResponse

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	return &responseBody[0].Symbol[0].Data, nil
}
