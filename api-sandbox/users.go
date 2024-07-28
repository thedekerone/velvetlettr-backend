package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {

	createUser()
}

func getUsers() {
	resp, err := http.Get("http://localhost:8080/users")

	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Printf("%s", body)
}

func createUser() {
	body := []byte(`{
		"firstName": "wedding test",
		"lastName": "idk some",
		"email": "test_wed@gmail.com",
		"password": "internetes"
	}`)
	res, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Print(err)
	}

	defer res.Body.Close()
	respBody, _ := io.ReadAll(res.Body)

	log.Printf("%s", respBody)

}
