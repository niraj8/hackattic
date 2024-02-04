package main

import (
	"fmt"
	"os"
)

func main() {
	// var problem struct {
	// 	ZipURL string `json:"zip_url"`
	// }

	// err := helper.GetChallenge("brute_force_zip", &problem)
	// if err != nil {
	// 	log.Fatalf("error getting challenge: %v", err)
	// }

	// resp, err := http.Get(problem.ZipURL)
	// if err != nil {
	// 	log.Fatalf("error fetching the zipfile")
	// }
	// defer resp.Body.Close()

	// file, err := os.Create("brute_force_zip.zip")
	// if err != nil {
	// 	log.Fatalf("error creating file: %v", err)
	// }
	// defer file.Close()

	// _, err = io.Copy(file, resp.Body)
	// if err != nil {
	// 	log.Fatalf("error writing to file: %v", err)
	// }

	dump()
}

func dump() {
	// can use https://www.kali.org/tools/crunch/ for this
	f, err := os.Create("passwords.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var chars []rune
	for i := 48; i <= 57; i++ {
		chars = append(chars, rune(i))
	}
	for i := 97; i <= 122; i++ {
		chars = append(chars, rune(i))
	}

	password := make([]rune, 4)
	for i := range chars {
		for j := range chars {
			for k := range chars {
				for l := range chars {
					password[0] = chars[i]
					password[1] = chars[j]
					password[2] = chars[k]
					password[3] = chars[l]
					f.Write([]byte(string(password) + "\n"))
				}
			}
		}
	}

	password = make([]rune, 5)
	for i := range chars {
		for j := range chars {
			for k := range chars {
				for l := range chars {
					for m := range chars {
						password[0] = chars[i]
						password[1] = chars[j]
						password[2] = chars[k]
						password[3] = chars[l]
						password[4] = chars[m]
						f.Write([]byte(string(password) + "\n"))
					}
				}
			}
		}
	}

	password = make([]rune, 6)
	for i := range chars {
		for j := range chars {
			for k := range chars {
				for l := range chars {
					for m := range chars {
						for n := range chars {
							password[0] = chars[i]
							password[1] = chars[j]
							password[2] = chars[k]
							password[3] = chars[l]
							password[4] = chars[m]
							password[5] = chars[n]
							f.Write([]byte(string(password) + "\n"))
						}
					}
				}
			}
		}
	}
}
