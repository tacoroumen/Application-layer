package main

import (
	"fmt"
	"os"
	"time"
)

func check_license(text string) func() bool {
	result := false
	allowed_plates := []string{"L653FD", "TL78JJ", "P565TT", "YT31VD", "XL655V"}

	return func() bool {
		for i := 0; i < len(allowed_plates); i++ {
			if allowed_plates[i] == text {
				result = true

			}
		}
		return result
	}
}

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
	plate := os.Args[1]
	plate_allowed := check_license(plate)
	dt := time.Now()
	if plate_allowed() {
		switch {
		case dt.Hour() >= 23 && dt.Hour() < 7:
			fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")

		default:
			message := get_message()
			fmt.Println(message(), "Welkom bij Fonteyn Vakantieparken")
		}
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}
