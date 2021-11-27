package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DeleteSong(songId int) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl --location --request DELETE 'http://127.0.0.1:8000/delete' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{"delete":11}'
	//

	type Payload struct {
		Delete int `json:"delete"`
	}

	data := Payload{
		// fill struct
		Delete: songId,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
		log.Println(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8000/delete", body)
	if err != nil {
		// handle err
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Println(err)
	}
	defer resp.Body.Close()

	responseB, _ := ioutil.ReadAll(resp.Body)
	// print reponseb in green
	fmt.Printf("\x1b[32m%s\x1b[0m\n", string(responseB))

}
