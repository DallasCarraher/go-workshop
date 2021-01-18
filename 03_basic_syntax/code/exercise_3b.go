package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var name, hometown string
	var years int
	var weather string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your hometown: ")
	fmt.Scan(&hometown)
	fmt.Print("Enter how many years you've lived there: ")
	fmt.Scan(&years)
	fmt.Print("Is the weather there nice?: ")
	fmt.Scan(&weather)

	weather = strings.ToUpper(weather)
	if weather == "YES" {
		weather = "true"
	} else {
		weather = "false"
	}
	weatherBool, _ := strconv.ParseBool(weather)

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t.\n", name, hometown, int(years), weatherBool)
}
