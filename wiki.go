package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, loadIndex())
}

func loadFile(path string) string {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return err.Error()
	}

	return string(body)
}

func loadIndex() string {
	return loadFile("index.html")
}

func main() {
	var port = os.Getenv("port")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
