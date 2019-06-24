package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("https://avena.io/panel/#/login");

	if(err != nil){
		panic(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(n)

}
