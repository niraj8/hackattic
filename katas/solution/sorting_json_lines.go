package solution

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

func SortingJsonLines(input []string) []string {
	type account struct {
		Balance       *int `json:"balance"`
		AccountNumber *int `json:"account_number"`
	}

	type userBalance struct {
		name    string
		balance int
	}

	var users []userBalance
	for _, line := range input {
		var account map[string]account
		buf := bytes.NewBuffer([]byte(line))
		err := json.NewDecoder(buf).Decode(&account)
		if err != nil {
			log.Println(err)
		}
		var userBalance userBalance
		for name, accountDetails := range account {
			if name == "extra" {
				userBalance.balance = *accountDetails.Balance
			} else if len(account) == 1 {
				userBalance.name = name
				userBalance.balance = *accountDetails.Balance
			} else {
				userBalance.name = name
			}

		}
		users = append(users, userBalance)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].balance < users[j].balance
	})

	result := make([]string, 0)
	for _, user := range users {
		result = append(result, fmt.Sprintf("%s: %s", user.name, thousandSeparatedInt(user.balance)))
	}

	return result
}

func thousandSeparatedInt(i int) string {
	s := reverse(fmt.Sprint(i))
	parts := make([]string, 0)
	for i := 0; i < len(s); i += 3 {
		parts = append(parts, s[i:min(i+3, len(s))])
	}
	return reverse(strings.Join(parts, ","))
}

func reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i -= 1 {
		result += string(s[i])
	}
	return result
}

// the std lib min() is available starting go 1.21 https://tip.golang.org/ref/spec#Min_and_max
func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
