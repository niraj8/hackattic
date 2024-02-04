package main

import (
	"log"
	"net/http"

	"niraj8/hackattic/challenges/helper"

	"github.com/tuotoo/qrcode"
)

func main() {
	var problem struct {
		ImageURL string `json:"image_url"`
	}
	helper.GetChallenge("reading_qr", &problem)
	log.Println(problem.ImageURL)
	resp, err := http.Get(problem.ImageURL)
	helper.HandleError(err)
	qrMatrix, err := qrcode.Decode(resp.Body)
	helper.HandleError(err)

	log.Println(qrMatrix.Content)
}
