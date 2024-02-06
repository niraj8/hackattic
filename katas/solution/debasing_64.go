package solution

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func Debasing64() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		decoded, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Println("Error decoding base64:", err)
			return
		}
		fmt.Println(string(decoded))
	}
}
