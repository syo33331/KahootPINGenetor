package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func generatePin() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%07d", rand.Intn(10000000))
}

func checkPinValidity(pin string) (bool, error) {
	url := fmt.Sprintf("https://kahoot.it/reserve/session/%s/", pin)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

func main() {
	var numberOfPins int
	fmt.Print("Enter the number of PINs to generate: ")
	fmt.Scanln(&numberOfPins)

	validFile, err := os.Create("Valid.txt")
	if err != nil {
		fmt.Println("Error creating Valid.txt:", err)
		return
	}
	defer validFile.Close()

	invalidFile, err := os.Create("Invalid.txt")
	if err != nil {
		fmt.Println("Error creating Invalid.txt:", err)
		return
	}
	defer invalidFile.Close()

	for i := 0; i < numberOfPins; i++ {
		pin := generatePin()
		valid, err := checkPinValidity(pin)
		if err != nil {
			fmt.Println("Error checking PIN validity:", err)
			continue
		}
		if valid {
			fmt.Println("[✔︎]Valid PIN:", pin)
			validFile.WriteString(pin + "\n")
		} else {
			fmt.Println("[✖︎]Invalid PIN:", pin)
			invalidFile.WriteString(pin + "\n")
		}
	}
}
