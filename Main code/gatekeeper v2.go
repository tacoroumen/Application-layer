package main

import (
	"fmt"
	"time"
)

func get_message() func() string {
	hello_message := ""
	dt := time.Now()

	return func() string {
		switch {
		case dt.Hour() >= 7 && dt.Hour() < 12:
			hello_message = "Goedemorgen"
		case dt.Hour() >= 12 && dt.Hour() < 18:
			hello_message = "Goedemiddag"
		case dt.Hour() >= 18 && dt.Hour() < 23:
			hello_message = "Goedenavond"
		default:
			fmt.Println("")
		}

		return hello_message
	}
}

func main() {
	dt := time.Now()
	switch {
	case dt.Hour() >= 23 && dt.Hour() < 7:
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")

	default:
		message := get_message()
		fmt.Println(message(), "Welkom bij Fonteyn Vakantieparken")
	}
}
