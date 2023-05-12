package main

import (
	"bytes"
	//"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	
	//_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Morning_start_time          int
	Noon_start_time             int
	Evening_start_time          int
	No_parking_acces_start_time int
	API_ip_or_domain            string
	API_port					string
	Morning_message             string
	Noon_message                string
	Evening_message             string
	No_parking_acces_message    string
	Technical_dificulties       string
	Welcome_message             string
	Not_allowed                 string
}

type Payload struct {
	Data string `json:"data"`
}

func main() {
	url := "http://127.0.0.1:8080/nummerplaat"
	//enable logger to errors.log
	logger, eror := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if eror != nil {
		log.Fatal(eror.Error())
	}
	defer logger.Close()
	access, eror := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if eror != nil {
		log.Fatal(eror.Error())
	}
	defer logger.Close()

	log.SetOutput(logger)

	//check if lincesplate provided
	plate := flag.String("plate", "", "Must provide a plate to check acces")
	flag.Parse()
	if !flag.Parsed() || *plate == "" {
		flag.Usage()
		log.Println("No licenseplate provided, program terminated.")
		os.Exit(1)
	}

	file, err := os.Open("config.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	var conf []Config

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Println(err)
		return
	}
	dt := time.Now()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("licenseplate", *plate)
	req.URL.RawQuery = q.Encode()

	// Make the request and get the response
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse the JSON response and get the name
	var response struct {
		Name string `json:"naam"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	userName := response.Name

	for _, Config := range conf {
		if userName == "" {
			fmt.Printf("%s\n", Config.Not_allowed)
			os.Exit(403)
		}
		log.SetOutput(access)
		log.Printf("%s gained access to parking.\n", *plate)
		log.SetOutput(logger)
		payload := Payload{Data: ""}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			log.Println("Failed to marshal payload:", err)
			return
		}
		URL := Config.API_ip_or_domain + ":" + Config.API_port + "/switch/gate/turn_on"
		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonPayload))
		req.SetBasicAuth("Proftaak", "AfslagA1!")
		if err != nil {
			log.Println("Failed to create http request:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		httpclient := &http.Client{}
		resp, err := httpclient.Do(req)
		if err != nil {
			log.Println("Failed to send http request:", err)
			return
		}
		defer resp.Body.Close()

		log.Println("Response status:", resp.Status)
        
		switch {
		case dt.Hour() >= Config.No_parking_acces_start_time:
			fmt.Printf("%s", Config.No_parking_acces_message)

		case dt.Hour() >= Config.Evening_start_time:
			fmt.Printf("%s %s! %s", Config.Evening_message, userName, Config.Welcome_message)

		case dt.Hour() >= Config.Noon_start_time:
			fmt.Printf("%s %s! %s", Config.Noon_message, userName, Config.Welcome_message)

		case dt.Hour() >= Config.Morning_start_time:
			fmt.Printf("%s %s! %s", Config.Morning_message, userName, Config.Welcome_message)

		default:
			fmt.Printf("%s", Config.Technical_dificulties)
		}
	}
}
