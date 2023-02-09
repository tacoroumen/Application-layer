package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check_license(text string) func() bool {
	result := false
	plates := []string{"L653FD\r\n", "TL78JJ\r\n", "P565TT\r\n", "YT31VD\r\n", "XL655V\r\n"}

	return func() bool {
		for i := 0; i < len(plates); i++ {
			if plates[i] == text {
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the licenseplate: ")
	text, _ := reader.ReadString('\n')
	plate_allowed := check_license(text)
	dt := time.Now()
	if plate_allowed() {
		switch {
		case dt.Hour() >= 23 && dt.Hour() < 7:
			fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")

		default:
			message := get_message()
			fmt.Println(message(), "Welkom bij Fonteyn Vakantieparken")
		}
	}
}
