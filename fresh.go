package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var converters []converter

func main() {
	fmt.Println("Start the Engines")

	converters = loadConverters("converters.json")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	w.Write([]byte(fmt.Sprintf("Converters:%v", converters)))
}

func loadConverters(filename string) []converter {
	//TODO validate filename
	var converters []converter

	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &converters)

	return converters
}
