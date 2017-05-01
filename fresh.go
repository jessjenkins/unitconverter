package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	converters := make([]converter, 3)

	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		var cv converter
		if err := dec.Decode(&cv); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %v\n", cv.Name, cv.Unit)
		converters = append(converters, cv)
	}
	return converters
}
