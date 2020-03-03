package getline

// note src for packages needs to be in ~/go/src/packagename see: https://golangbot.com/packages/

import (
	"bufio"
	"fmt"
	"os"
)

func Fetch(prompt string) (lineOfCharacters string) {

	fmt.Println(prompt)
	/* Note to include whitespace when fetching input, you need to use bufio.Scanner, see
			   https://stackoverflow.com/questions/43843477/scanln-in-golang-doesnt-accept-whitespace
	   	   	   	fmt.Scanln(&aString)
	*/
	keyScanner := bufio.NewScanner(os.Stdin)
	if keyScanner.Scan() {
		lineOfCharacters = keyScanner.Text()
	}

	return lineOfCharacters
}
