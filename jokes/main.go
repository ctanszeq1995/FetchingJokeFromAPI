package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is a simple program that fetches joke from the API provided by 'THE INTERNET CHUCK NORRIS DATABASE'")
	fmt.Println("Website to 'THE INTERNET CHUCK NORRIS DATABASE': http://www.icndb.com/")
	input := ""
	for {
		fmt.Println()
		fmt.Println("Enter \"exit\" to end the program. Press any key to get jokes.")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			if strings.ToLower(input) == "exit" {
				os.Exit(0)
			}
			response, err := http.Get("http://api.icndb.com/jokes/random")
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				if strings.Contains(string(data), "success") {
					joke := string(data)[strings.Index(string(data), "\"joke\"") + 9 : strings.Index(string(data), "\"categories\"") - 3]
					fmt.Println(joke)
				} else {
					fmt.Println("Unexpected error while retrieving joke from API!")
					os.Exit(1)
				}
			}
		}
	}
}
